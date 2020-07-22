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

type TaskHandler struct {
	service   services.TaskService
	validator *validator.Validate
}

func NewTaskHandler(adapter adapters.IDbAdapter) *TaskHandler {
	return &TaskHandler{
		service: services.TaskService{
			Repository: repositories.TaskRepository{IDbAdapter: adapter},
		},
		validator: validator.New(),
	}
}

func (handler TaskHandler) GetByColumnId(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	columnId, err := strconv.ParseInt(vars["columnId"], 10, 64)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tasks, err := handler.service.GetByColumnId(columnId)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (handler TaskHandler) Get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := handler.service.GetById(id)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return

	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (handler TaskHandler) Create(w http.ResponseWriter, req *http.Request) {
	var task models.TaskModel
	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.validator.Struct(&task)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTask, err := handler.service.Create(task)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newTask); err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (handler TaskHandler) Update(w http.ResponseWriter, req *http.Request) {
	var task models.TaskModel
	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.validator.Struct(&task)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedTask, err := handler.service.Update(task)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedTask); err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (handler TaskHandler) Delete(w http.ResponseWriter, req *http.Request) {
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
