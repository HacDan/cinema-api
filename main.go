package main

import (
	"fmt"
	"log"

	"github.com/hacdan/cinema-api/api"
)

func main() {
	listenAddr := ":8080"
	server := api.NewServer(listenAddr)
	fmt.Printf("Starting server on localhost%s\n", listenAddr)
	log.Fatal(server.Start())
}
