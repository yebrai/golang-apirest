package main

import (
	"fmt"
	"golang-apirest/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/todo/create", handlers.CreateTodo)
	http.HandleFunc("/todo/edit", handlers.EditTodo)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, welcome to my simple API!")
}
