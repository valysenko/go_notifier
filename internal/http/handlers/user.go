package handlers

import (
	"encoding/json"
	"go_notifier/internal/dto"
	"go_notifier/internal/http/helpers"
	"go_notifier/internal/service"
	"net/http"
)

type CreateUserResponse struct {
	UUID string `json:"uuid"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// request
	var dto dto.User
	err := helpers.CreateAndValidateFromRequest(r, &dto)
	if err != nil {
		helpers.HandleRequestError(err, w)
		return
	}

	// run business logic
	_, err = service.CreateUser(&dto)
	if err != nil {
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
