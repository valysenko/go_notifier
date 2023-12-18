package handlers

import (
	"encoding/json"
	"go_notifier/internal/dto"
	"go_notifier/internal/http/helpers"
	"go_notifier/internal/service"
	"net/http"
)

type CreateDeviceResponse struct {
	Token string `json:"token"`
}

func CreateDeviceHandler(w http.ResponseWriter, r *http.Request) {
	// request
	var dto dto.Device
	err := helpers.CreateAndValidateFromRequest(r, &dto)
	if err != nil {
		helpers.HandleRequestError(err, w)
		return
	}

	// run business logic
	_, err = service.CreateDevice(&dto)
	if err != nil {
		http.Error(w, "error while device creation", http.StatusInternalServerError)
		return
	}

	// response
	responseData := CreateDeviceResponse{Token: dto.Token}
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
