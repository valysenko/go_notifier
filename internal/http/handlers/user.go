package handlers

import (
	"encoding/json"
	"errors"
	"go_notifier/internal/dto"
	"go_notifier/internal/http/helpers"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

type UserService interface {
	CreateUser(dto *dto.User) (int64, error)
}

type CreateUserResponse struct {
	UUID string `json:"uuid"`
}

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// request
	var dto dto.User
	err := helpers.CreateAndValidateFromRequest(r, &dto)
	if err != nil {
		helpers.HandleRequestError(err, w)
		return
	}

	// run business logic
	_, err = h.service.CreateUser(&dto)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			http.Error(w, "user already exists", http.StatusBadRequest)
			return
		}
		http.Error(w, "error while user creation", http.StatusInternalServerError)
		return
	}

	// response
	responseData := CreateUserResponse{UUID: dto.UUID}
	jsonResponse, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(jsonResponse)
	w.WriteHeader(http.StatusCreated)
}
