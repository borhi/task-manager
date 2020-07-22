package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"task-manager/handlers"
)

type router struct {
}

func (router *router) InitRouter() *mux.Router {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./swaggerui/"))
	r.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", fs))

	s := r.PathPrefix("/v1").Subrouter()

	projectHandler := handlers.NewProjectHandler()
	s.HandleFunc("/projects", projectHandler.GetList).Methods(http.MethodGet)
	s.HandleFunc("/project/{id}", projectHandler.Get).Methods(http.MethodGet)
	s.HandleFunc("/project", projectHandler.Create).Methods(http.MethodPost)
	s.HandleFunc("/project/{id}", projectHandler.Update).Methods(http.MethodPut)
	s.HandleFunc("/project/{id}", projectHandler.Delete).Methods(http.MethodDelete)

	columnHandler := handlers.NewColumnHandler()
	s.HandleFunc("/columns/{projectId}", columnHandler.GetByProjectId).Methods(http.MethodGet)
	s.HandleFunc("/column/{id}", columnHandler.Get).Methods(http.MethodGet)
	s.HandleFunc("/column", columnHandler.Create).Methods(http.MethodPost)
	s.HandleFunc("/column/{id}", columnHandler.Update).Methods(http.MethodPut)
	s.HandleFunc("/column/{id}", columnHandler.Delete).Methods(http.MethodDelete)

	taskHandler := handlers.NewTaskHandler()
	s.HandleFunc("/tasks/{columnId}", taskHandler.GetByColumnId).Methods(http.MethodGet)
	s.HandleFunc("/task/{id}", taskHandler.Get).Methods(http.MethodGet)
	s.HandleFunc("/task", taskHandler.Create).Methods(http.MethodPost)
	s.HandleFunc("/task/{id}", taskHandler.Update).Methods(http.MethodPut)
	s.HandleFunc("/task/{id}", taskHandler.Delete).Methods(http.MethodDelete)

	commentHandler := handlers.NewCommentHandler()
	s.HandleFunc("/comments/{taskId}", commentHandler.GetByTaskId).Methods(http.MethodGet)
	s.HandleFunc("/comment/{id}", commentHandler.Get).Methods(http.MethodGet)
	s.HandleFunc("/comment", commentHandler.Create).Methods(http.MethodPost)
	s.HandleFunc("/comment/{id}", commentHandler.Update).Methods(http.MethodPut)
	s.HandleFunc("/comment/{id}", commentHandler.Delete).Methods(http.MethodDelete)

	return r
}
