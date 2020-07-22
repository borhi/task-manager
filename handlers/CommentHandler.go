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

type CommentHandler struct {
	service   services.CommentService
	validator *validator.Validate
}

func NewCommentHandler(adapter adapters.IDbAdapter) *CommentHandler {
	return &CommentHandler{
		service: services.CommentService{
			Repository: repositories.CommentRepository{IDbAdapter: adapter},
		},
		validator: validator.New(),
	}
}

func (handler CommentHandler) GetByTaskId(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	commentId, err := strconv.ParseInt(vars["taskId"], 10, 64)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comments, err := handler.service.GetByTaskId(commentId)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(comments); err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (handler CommentHandler) Get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comment, err := handler.service.GetById(id)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return

	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (handler CommentHandler) Create(w http.ResponseWriter, req *http.Request) {
	var comment models.CommentModel
	err := json.NewDecoder(req.Body).Decode(&comment)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.validator.Struct(&comment)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newComment, err := handler.service.Create(comment)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newComment); err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (handler CommentHandler) Update(w http.ResponseWriter, req *http.Request) {
	var comment models.CommentModel
	err := json.NewDecoder(req.Body).Decode(&comment)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.validator.Struct(&comment)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedComment, err := handler.service.Update(comment)
	if err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedComment); err != nil {
		zap.L().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (handler CommentHandler) Delete(w http.ResponseWriter, req *http.Request) {
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
