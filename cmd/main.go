package main

import (
	"github.com/gorilla/mux"
	"github.com/megamsquare/dsanda/routes"
	"log"
	"net/http"
	_"gorm.io/driver/postgres"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":9010", router))
}