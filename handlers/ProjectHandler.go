package handlers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"task-manager/adapters"
	"task-manager/models"
	"task-manager/repositories"
	"task-manager/services"
)

type ProjectHandler struct {
	service   services.ProjectService
	validator *validator.Validate
}

func NewProjectHandler(adapter adapters.IDbAdapter) *ProjectHandler {
	return &ProjectHandler{
		service: services.ProjectService{
			ProjectRepository: repositories.ProjectRepository{IDbAdapter: adapter},
			ColumnRepository:  repositories.ColumnRepository{IDbAdapter: adapter},
		},
		validator: validator.New(),
	}
}

func (handler ProjectHandler) GetList(w http.ResponseWriter, req *http.Request) {
	projects, err := handler.service.GetList()
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(projects); err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (handler ProjectHandler) Get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	project, err := handler.service.GetById(id)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return

	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(project); err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (handler ProjectHandler) Create(w http.ResponseWriter, req *http.Request) {
	var project models.ProjectModel
	err := json.NewDecoder(req.Body).Decode(&project)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.validator.Struct(&project)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newProject, err := handler.service.Create(project)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newProject); err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (handler ProjectHandler) Update(w http.ResponseWriter, req *http.Request) {
	var project models.ProjectModel
	err := json.NewDecoder(req.Body).Decode(&project)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.validator.Struct(&project)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedProject, err := handler.service.Update(project)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedProject); err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (handler ProjectHandler) Delete(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.service.DeleteById(id)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode("success"); err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
