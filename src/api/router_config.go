package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

const addr = "localhost:8080"

func ConfigureRoutes() error {
	router := mux.NewRouter()

	router.HandleFunc("/employees", HandlePostEmployeesData).Methods("POST")

	err := http.ListenAndServe(addr, router)
	if err != nil {
		return err
	}

	return nil
}
