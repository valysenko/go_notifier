package campaign

import (
	"encoding/json"
	"go_notifier/internal/common"
	"go_notifier/internal/http/helpers"
	"net/http"
)

type CampaignServiceInterface interface {
	CreateCampaign(request *common.CampaignRequest) (string, error)
}

type CreateCampaignResponse struct {
	UUID string `json:"uuid"`
}

type CampaignHandler struct {
	service CampaignServiceInterface
}

func NewCampaignHandler(s CampaignServiceInterface) *CampaignHandler {
	return &CampaignHandler{
		service: s,
	}
}

func (h *CampaignHandler) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	// request
	var request common.CampaignRequest
	err := helpers.CreateAndValidateFromRequest(r, &request)
	if err != nil {
		helpers.HandleRequestError(err, w)
		return
	}

	// run business logic
	campaignUUID, err := h.service.CreateCampaign(&request)
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
