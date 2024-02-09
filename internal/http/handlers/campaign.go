package handlers

import (
	"encoding/json"
	"go_notifier/internal/dto"
	"go_notifier/internal/http/helpers"
	"net/http"
)

type CampaignService interface {
	CreateCampaign(dto *dto.Campaign) (string, error)
}

type CreateCampaignResponse struct {
	UUID string `json:"uuid"`
}

type CampaignHandler struct {
	service CampaignService
}

func NewCampaignHandler(s CampaignService) *CampaignHandler {
	return &CampaignHandler{
		service: s,
	}
}

func (h *CampaignHandler) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	// request
	var dto dto.Campaign
	err := helpers.CreateAndValidateFromRequest(r, &dto)
	if err != nil {
		helpers.HandleRequestError(err, w)
		return
	}

	// run business logic
	campaignUUID, err := h.service.CreateCampaign(&dto)
	if err != nil {
		http.Error(w, "error while campaign creation", http.StatusInternalServerError)
		return
	}

	// response
	responseData := CreateCampaignResponse{UUID: campaignUUID}
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
