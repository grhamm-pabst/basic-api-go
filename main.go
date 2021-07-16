package main

import (
	"log"
	"net/http"

	"grhamm.com/todo/routes"
)

func main() {

	log.Print("Server on")
	http.ListenAndServe(":3000", routes.RegisterRoute())
}
