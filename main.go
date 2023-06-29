package main

import (
	"log"
	"net/http"
	"todo-list/db"
	"todo-list/task"
	"todo-list/user"
)

func main() {
	db.InitDB(&user.User{}, &task.Task{})
	router := initRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
