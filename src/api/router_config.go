package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

const addr = "localhost:8080"

func ConfigureRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/employees", HandlePostEmployeesData).Methods("POST")

	http.ListenAndServe(addr, router)
}
