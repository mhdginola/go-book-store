package routes

import (
	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func (router *mux.Router){
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
}