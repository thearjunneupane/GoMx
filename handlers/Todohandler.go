package handlers

import (
	"gomx/model"
	"gomx/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func TodoShowHandler(w http.ResponseWriter, req *http.Request) {
	todos, err := model.GetAllTodos()
	if err != nil {
		log.Println("Failed !! Error Getting Todos From Database: ", err)
	}

	if len(todos) < 1 {
		log.Println("Todos Empty in the database: ")
	}

	log.Println("Success!! Got All Todos")

	data := struct {
		TodosList []model.Todo
		ErrorMsg  string
	}{
		TodosList: todos,
		ErrorMsg:  "",
	}

	utils.ExecTemplate(w, "todosList.html", data)
}

func sendTodos(w http.ResponseWriter, errorMsg string) {
	todos, err := model.GetAllTodos()
	if err != nil {
		log.Println("Failed !! Error Getting Todos From Database: ", err)
	}

	data := struct {
		TodosList []model.Todo
		ErrorMsg  string
	}{
		TodosList: todos,
		ErrorMsg:  errorMsg,
	}
	utils.ExecTemplate(w, "Todos", data)
}

func MarkTodo(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(req)["id"], 10, 64)
	if err != nil {
		log.Println("Coultn't Parse ID", err)
	}

	done, err := model.MarkTodo(id)
	if err != nil {
		log.Println("Failed !! Error While Marking Todo!", err)
	}

	if done {
		log.Println("Success!! Marked Todo done")
	} else {
		log.Println("Success!! Marked Todo undone")
	}

	sendTodos(w, "")

}

func CreateTodo(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println("Failed !! Error While Parsing Form", err)
	}

	todo := req.FormValue("todo")

	if len(todo) < 1 {
		sendTodos(w, "Can't Create an empty Todo!!")
		return
	}

	if len(todo) > 255 {
		sendTodos(w, "Very Long Length For Just a Todo!!")
		return
	}

	err = model.CreateTodo(todo)
	if err != nil {
		log.Println("Failed !! Error While Creating Todo", err)
	}
	log.Println("Success!! Created New Todo")
	sendTodos(w, "")

}

func DeleteTodo(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(req)["id"], 10, 64)
	if err != nil {
		log.Println("Coultn't Parse ID", err)
	}

	err = model.DeleteTodo(id)
	if err != nil {
		log.Println("Failed !! Error While Deleting Todo", err)
	}
	log.Println("Success!! Deleted Todo")

	sendTodos(w, "")
}

func EditTodo(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(req)["id"], 10, 64)
	if err != nil {
		log.Println("Couldn't Parse ID", err)
	}

	if err = req.ParseForm(); err != nil {
		log.Println("Failed !! Error While Parsing Form", err)
	}

	newtodo := req.FormValue("newtodo")

	err = model.EditTodo(id, newtodo)
	if err != nil {
		log.Println("Failed !! Error While Editing Todo", err)
	}
	log.Println("Success!! Edited Todo")

	sendTodos(w, "")
}
