package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const addrTmplt = "localhost:%s"

func ConfigureRoutesAndListen() error {
	router := mux.NewRouter()

	router.HandleFunc("/employees", HandlePostEmployeesData).Methods("POST")

	addr := fmt.Sprintf(addrTmplt, os.Getenv(httpPortEnv))
	err := http.ListenAndServe(addr, router)
	if err != nil {
		return err
	}

	return nil
}
