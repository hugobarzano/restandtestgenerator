package controller

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"cats/mongo"
)


var controller = Controller{Storer: mongo.Storer{}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route


//Example var API_GENERATED = "/cats"

var API_GENERATED = "/cats"

var routes = Routes {
	Route {
		"Index",
		"GET",
		API_GENERATED,
		controller.Index,
	},
	Route {
		"AddBusinessObject",
		"POST",
		API_GENERATED,
		controller.AddBusinessObject,
	},
	Route {
		"UpdateBusinessObject",
		"PUT",
		API_GENERATED+"/{id}",
		controller.UpdateBusinessObject,
	},
	// Get BusinessObject by {id}
	Route {
		"GetBusinessObject",
		"GET",
		API_GENERATED+"/{id}",
		controller.GetBusinessObject,
	},
	// Delete BusinessObject by {id}
	Route {
		"DeleteBusinessObject",
		"DELETE",
		API_GENERATED+"/{id}",
		controller.DeleteBusinessObject,
	},
	// Search BusinessObject with string
	Route {
		"SearchBusinessObject",
		"GET",
		API_GENERATED+"/search/{query}",
		controller.SearchBusinessObject,
	},
	}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

