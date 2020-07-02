package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"task-manager/handlers"
)

type router struct {
	projectHandler *handlers.ProjectHandler
	columnHandler  *handlers.ColumnHandler
	taskHandler    *handlers.TaskHandler
	commentHandler *handlers.CommentHandler
}

func (router router) InitRouter() *mux.Router {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./swaggerui/"))
	r.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", fs))

	s := r.PathPrefix("/v1").Subrouter()

	s.HandleFunc("/projects", router.projectHandler.GetList).Methods(http.MethodGet)
	s.HandleFunc("/project/{id}", router.projectHandler.Get).Methods(http.MethodGet)
	s.HandleFunc("/project", router.projectHandler.Create).Methods(http.MethodPost)
	s.HandleFunc("/project", router.projectHandler.Update).Methods(http.MethodPut)
	s.HandleFunc("/project/{id}", router.projectHandler.Delete).Methods(http.MethodDelete)

	s.HandleFunc("/columns/{projectId}", router.columnHandler.GetByProjectId).Methods(http.MethodGet)
	s.HandleFunc("/column/{id}", router.columnHandler.Get).Methods(http.MethodGet)
	s.HandleFunc("/column", router.columnHandler.Create).Methods(http.MethodPost)
	s.HandleFunc("/column", router.columnHandler.Update).Methods(http.MethodPut)
	s.HandleFunc("/column/{id}", router.columnHandler.Delete).Methods(http.MethodDelete)

	s.HandleFunc("/tasks/{columnId}", router.taskHandler.GetByColumnId).Methods(http.MethodGet)
	s.HandleFunc("/task/{id}", router.taskHandler.Get).Methods(http.MethodGet)
	s.HandleFunc("/task", router.taskHandler.Create).Methods(http.MethodPost)
	s.HandleFunc("/task", router.taskHandler.Update).Methods(http.MethodPut)
	s.HandleFunc("/task/{id}", router.taskHandler.Delete).Methods(http.MethodDelete)

	s.HandleFunc("/comments/{taskId}", router.commentHandler.GetByTaskId).Methods(http.MethodGet)
	s.HandleFunc("/comment/{id}", router.commentHandler.Get).Methods(http.MethodGet)
	s.HandleFunc("/comment", router.commentHandler.Create).Methods(http.MethodPost)
	s.HandleFunc("/comment", router.commentHandler.Update).Methods(http.MethodPut)
	s.HandleFunc("/comment/{id}", router.commentHandler.Delete).Methods(http.MethodDelete)

	return r
}
