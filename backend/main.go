package main

import (
	"backend/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 8877...")
	log.Fatal(http.ListenAndServe(":8877", r))

}
