package handlers

import (
	"encoding/json"
	"fmt"
	"go_notifier/internal/dto"
	"go_notifier/internal/http/helpers"
	"net/http"
)

type UserCampaignService interface {
	CreateUserCampaign(dto *dto.CampaignUser) (string, error)
}

type CreateUserCampaignResponse struct {
	Time string `json:"time"`
}

type UserCampaignHandler struct {
	service UserCampaignService
}

func NewUserCampaignHandler(s UserCampaignService) *UserCampaignHandler {
	return &UserCampaignHandler{
		service: s,
	}
}

func (h *UserCampaignHandler) CreateUserCampaign(w http.ResponseWriter, r *http.Request) {
	// request
	var dto dto.CampaignUser
	err := helpers.CreateAndValidateFromRequest(r, &dto)
	if err != nil {
		helpers.HandleRequestError(err, w)
		return
	}

	// run business logic
	time, err := h.service.CreateUserCampaign(&dto)
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
