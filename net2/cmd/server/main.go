package main

import (
	"log"
	"os"

	"net/internal/server"
)

func main() {
	var port string
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = "8989"
	}

	server := server.NewServer(":" + port)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
