package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/megamsquare/dsanda/routes"
	_ "gorm.io/driver/postgres"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":9010", router))
}
