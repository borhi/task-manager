package handlers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"task-manager/models"
	"task-manager/repositories"
	"task-manager/services"
)

type ColumnHandler struct {
	service   services.ColumnService
	validator *validator.Validate
}

func NewColumnHandler() *ColumnHandler {
	return &ColumnHandler{
		service: services.ColumnService{
			Repository: repositories.ColumnRepository{},
		},
		validator: validator.New(),
	}
}

func (handler ColumnHandler) GetByProjectId(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	projectId, err := strconv.ParseUint(vars["projectId"], 10, 64)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	columns, err := handler.service.GetByProjectId(uint(projectId))
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(columns); err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (handler ColumnHandler) Get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	project, err := handler.service.GetById(uint(id))
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

func (handler ColumnHandler) Create(w http.ResponseWriter, req *http.Request) {
	var column models.ColumnModel
	err := json.NewDecoder(req.Body).Decode(&column)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.validator.Struct(&column)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newColumn, err := handler.service.Create(column)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newColumn); err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (handler ColumnHandler) Update(w http.ResponseWriter, req *http.Request) {
	var column models.ColumnModel
	err := json.NewDecoder(req.Body).Decode(&column)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.validator.Struct(&column)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedColumn, err := handler.service.Create(column)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedColumn); err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (handler ColumnHandler) Delete(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.service.DeleteById(uint(id))
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
