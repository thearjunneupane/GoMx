package router

import (
	"gomx/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupAndRun() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.TodoShowHandler).Methods("GET")
	todo := r.PathPrefix("/todo").Subrouter()
	todo.HandleFunc("/create", handlers.CreateTodo).Methods("POST")
	todo.HandleFunc("/{id}", handlers.MarkTodo).Methods("PUT")
	todo.HandleFunc("/edit/{id}", handlers.EditTodo).Methods("PUT")
	todo.HandleFunc("/{id}", handlers.DeleteTodo).Methods("DELETE")

	log.Println("Server Live at: http://localhost:8080")

	// utils.OpenInBrowser("chrome", "http://localhost:8080", true)

	log.Fatal(http.ListenAndServe(":8080", r))

}
