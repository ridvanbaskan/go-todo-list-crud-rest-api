package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"todo-list/auth"
	"todo-list/middleware"
	"todo-list/task"
)

func initRouter() *mux.Router {
	router := mux.NewRouter()

	taskHandler := task.NewTaskHandler()
	authHandler := &auth.Handler{}

	router.
		Handle("/tasks", middleware.Authenticate(http.HandlerFunc(taskHandler.CreateTask))).Methods("POST")
	router.
		Handle("/tasks", middleware.Authenticate(http.HandlerFunc(taskHandler.GetTasks))).Methods("GET")
	router.Handle("/tasks/{id}", middleware.Authenticate(http.HandlerFunc(taskHandler.GetTask))).Methods("GET")
	router.Handle("/tasks/{id}", middleware.Authenticate(http.HandlerFunc(taskHandler.DeleteTask))).Methods("DELETE")
	router.Handle("/tasks/{id}", middleware.Authenticate(http.HandlerFunc(taskHandler.UpdateTask))).Methods("PUT")
	router.Handle("/tasks/{id}/completed", middleware.Authenticate(http.HandlerFunc(taskHandler.MarkTaskAs))).Methods("PATCH")

	router.HandleFunc("/users/register", authHandler.RegisterUser).Methods("POST")
	router.HandleFunc("/users/login", authHandler.Login).Methods("POST")

	return router
}
