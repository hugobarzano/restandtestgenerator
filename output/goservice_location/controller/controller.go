package controller

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"fmt"
	"github.com/gorilla/mux"
	"location/models"
	"location/mongo"
	"strings"
)
//Controller ...
type Controller struct {
	Storer mongo.Storer
}

// Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	businessObjects := c.Storer.GetBusinessObjects() // list of all businessObjects
	data, _ := json.Marshal(businessObjects)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// AddBusinessObject POST /
func (c *Controller) AddBusinessObject(w http.ResponseWriter, r *http.Request) {
	var businessObject models.BusinessObject
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request

	log.Println(body)

	if err != nil {
		log.Fatalln("Error AddBusinessObject", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddBusinessObject", err)
	}

	if err := json.Unmarshal(body, &businessObject); err != nil { // unmarshall body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddBusinessObject unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	log.Println(businessObject)
	success := c.Storer.AddBusinessObject(businessObject) // adds the businessObject to the DB
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}

// UpdateBusinessObject PUT /
func (c *Controller) UpdateBusinessObject(w http.ResponseWriter, r *http.Request) {
	var businessObject models.BusinessObject
	vars := mux.Vars(r)
	log.Println(vars)
	id := vars["id"] // param id
	log.Println(id);

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Fatalln("Error UpdateBusinessObject", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}



	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error UpdateBusinessObject", err)
	}

	if err := json.Unmarshal(body, &businessObject); err != nil { // unmarshall body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error UpdateBusinessObject unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	fmt.Println("TO BE UPDATED BusinessObject ID - ", id)
	businessObject.ID=id
	success := c.Storer.UpdateBusinessObject(id,businessObject) // updates the businessObject in the DB

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

// GetBusinessObject GET - Gets a single businessObject by ID /
func (c *Controller) GetBusinessObject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)

	id := vars["id"] // param id
	log.Println(id);


	businessObject := c.Storer.GetBusinessObjectById(id)
	data, _ := json.Marshal(businessObject)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// DeleteBusinessObject DELETE /
func (c *Controller) DeleteBusinessObject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)
	id := vars["id"] // param id
	log.Println(id);


	if err := c.Storer.DeleteBusinessObject(id); err != "" { // delete a businessObject by id
		log.Println(err);
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

// SearchBusinessObject GET /
func (c *Controller) SearchBusinessObject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)

	query := vars["query"] // param query
	log.Println("Search Query - " + query);

	businessObjects := c.Storer.GetBusinessObjectsByString(query)
	data, _ := json.Marshal(businessObjects)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}