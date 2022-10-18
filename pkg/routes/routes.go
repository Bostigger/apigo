package routes

import (
	"github.com/bostigger/bukstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/create-book", controllers.CreateNewStoreBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/update-book/{id}", controllers.UpdateStoreBook).Methods("PUT")
	router.HandleFunc("/get-book/{id}", controllers.GetStoreBookById).Methods("GET")
	router.HandleFunc("/delete-book/{id}", controllers.DeleteStoreBook).Methods("DELETE")
}
