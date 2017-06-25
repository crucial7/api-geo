package main

import (
	"log"
	"net/http"

	"github.com/crucial7/api-geo/handlers"
)

func main() {

	http.HandleFunc("/latlongByAddress/", handlers.LatLongByAddress)
	log.Println("Server listening on port 3001")
	log.Println("\tRoutes:")
	log.Println("\t\tGET /latlongByAddress")
	log.Fatal(http.ListenAndServe(":3001", nil))
}
