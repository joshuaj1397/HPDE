package main

import (
	"github.com/gorilla/mux"
	HPDE "github.com/joshuaj1397/HPDE/api"
)

const port = 3005

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", HPDE.Index).Methods("GET")
}
