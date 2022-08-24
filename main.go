package main

import (
	"log"
	"mopi/mopi"
	"net/http"
)

func main() {
	s := mopi.NewServer()
	http.HandleFunc("/register", s.Register)
	http.HandleFunc("/", s.Endpoint)

	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))

}
