package routes

import (
	"log"
	"net/http"

	"github.com/Dagurasu56/go-rest-api/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")

	log.Fatal(http.ListenAndServe(":8000", r))
}
