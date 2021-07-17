package main

import (
	"log"
	"net/http"

	"grhamm.com/todo/data"
	"grhamm.com/todo/routes"
)

func main() {

	data.InitDatabase()

	log.Print("Server on")
	http.ListenAndServe(":3000", routes.RegisterRoute())
}
