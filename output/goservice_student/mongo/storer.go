package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strings"

	"student/models"
)

//Repository ...
type Storer struct{
	Sesion Session
}

// DBNAME the name of the DB instance
const DBNAME = "studentdb"

// COLLECTION is the name of the collection in DB
const COLLECTION = "students"




// GetBusinessObjects returns the list of BusinessObjects
func (s Storer) GetBusinessObjects() models.BusinessObjects {
	session, err := NewSession()

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	//defer session.Close()

	c := session.Copy().session.DB(DBNAME).C(COLLECTION)
	results := models.BusinessObjects{}

	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	session.session.Close()
	return results
}

// GetBusinessObjecttById returns a unique BusinessObject
func (s Storer) GetBusinessObjectById(id string ) models.BusinessObject {
	session, err := NewSession()

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	//defer session.Close()

	c := session.Copy().session.DB(DBNAME).C(COLLECTION)
	var result models.BusinessObject

	fmt.Println("ID in GetBusinessObjectById", id)
	//bson.M{"_id": id}
	query:=bson.M{"_id": id}
	if err := c.Find(query).One(&result); err != nil {
		fmt.Println("Failed to write result:", err)
	}

	//result.ID=bson.ObjectId(result.ID).Hex()
	session.Close()
	return result
}

// GetBusinessObjectbyString takes a search string as input and returns BusinessObjects
func (s Storer) GetBusinessObjectsByString(query string) models.BusinessObjects {
	session, err := NewSession()

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	//defer session.Close()

	c := session.Copy().session.DB(DBNAME).C(COLLECTION)
	result := models.BusinessObjects{}

	// Logic to create filter
	qs := strings.Split(query, " ")
	and := make([]bson.M, len(qs))
	for i, q := range qs {
		and[i] = bson.M{"title": bson.M{
			"$regex": bson.RegEx{Pattern: ".*" + q + ".*", Options: "i"},
		}}
	}
	filter := bson.M{"$and": and}

	if err := c.Find(&filter).Limit(5).All(&result); err != nil {
		fmt.Println("Failed to write result:", err)
	}

	//for _,r := range result{
	//	r.ID=bson.ObjectId(r.ID).Hex()
	//}
	session.Close()
	return result
}

// AddBusinessObject adds a BusinessObject in the DB
func (s Storer) AddBusinessObject(businessObject models.BusinessObject) bool {
	session, err := NewSession()
	//defer session.Close()

	i := bson.NewObjectId()
	businessObject.ID=i.Hex()
	session.Copy().session.DB(DBNAME).C(COLLECTION).Insert(&businessObject)

	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Added New  BusinessObject Title- ", businessObject.ID)
	session.Close()
	return true
}

// UpdateBusinessObject updates a businessObject in the DB
func (s Storer) UpdateBusinessObject(id string, businessObject models.BusinessObject) bool {
	session, err := NewSession()
	//defer session.Close()


	//session.Copy().session.DB(DBNAME).C(COLLECTION).UpdateId(businessObject.ID, businessObject)
	err = session.Copy().session.DB(DBNAME).C(COLLECTION).UpdateId(id, businessObject)

	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Updated BusinessObject ID - ", id)
	session.Close()
	return true
}

// DeleteBusinessObject deletes an BusinessObject
func (s Storer) DeleteBusinessObject(id string) string {
	session, err := NewSession()
	//defer session.Close()

	// Remove businessObject
	query:=bson.M{"_id": id}
	if err = session.Copy().session.DB(DBNAME).C(COLLECTION).Remove(query); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}

	fmt.Println("Deleted BusinessObject ID - ", id)
	// Write status
	session.Close()
	return "OK"
}