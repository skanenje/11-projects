package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileserve := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserve)
	http.HandleFunc("/get1", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name1")
		fmt.Fprintf(w, "Hello, %s", name)
	})
	http.HandleFunc("/get2", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name2")
		fmt.Fprintf(w, "Welcome to, %s", name)
	})
	fmt.Println("server started at http://localhost:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
