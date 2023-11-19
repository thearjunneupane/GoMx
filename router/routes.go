package router

import (
	"gomx/handlers"
	"log"
	"net/http"
	"os"

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

	log.Println("Server Live at: http://localhost:3000")

	// utils.OpenInBrowser("chrome", "http://localhost:3000", true)
	host := "0.0.0.0:"
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	addrStr := host + port

	log.Fatal(http.ListenAndServe(addrStr, r))

}
