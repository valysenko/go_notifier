package handlers

import (
	"encoding/json"
	"fmt"
	"go_notifier/internal/dto"
	"go_notifier/internal/http/helpers"
	"go_notifier/internal/service"
	"net/http"
)

type CreateUserCampaignResponse struct {
	Time string `json:"time"`
}

func CreateUserCampaignHandler(w http.ResponseWriter, r *http.Request) {
	// request
	var dto dto.CampaignUser
	err := helpers.CreateAndValidateFromRequest(r, &dto)
	if err != nil {
		helpers.HandleRequestError(err, w)
		return
	}

	// run business logic
	time, err := service.CreateUserCampaign(&dto)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error while campaign-user creation", http.StatusInternalServerError)
		return
	}

	// response
	responseData := CreateUserCampaignResponse{Time: time}
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
