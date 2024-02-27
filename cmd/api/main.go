package main

import (
	"faas/internal/server"
	"log"
)

func main() {

	server := server.NewServer()

	log.Fatalln(server.ListenAndServe())

}
