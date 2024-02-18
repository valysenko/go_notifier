package user

import (
	"encoding/json"
	"errors"
	"go_notifier/internal/app/http/helpers"
	"go_notifier/internal/common"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

type UserServiceInterface interface {
	CreateUser(request *common.UserRequest) (int64, error)
}

type CreateUserResponse struct {
	UUID string `json:"uuid"`
}

type UserHandler struct {
	service UserServiceInterface
}

func NewUserHandler(service UserServiceInterface) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// request
	var request common.UserRequest
	err := helpers.CreateAndValidateFromRequest(r, &request)
	if err != nil {
		helpers.HandleRequestError(err, w)
		return
	}

	// run business logic
	_, err = h.service.CreateUser(&request)
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
	responseData := CreateUserResponse{UUID: request.UUID}
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
