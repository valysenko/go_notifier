package user_campaign

import (
	"encoding/json"
	"fmt"
	"go_notifier/internal/app/http/helpers"
	"go_notifier/internal/common"
	"net/http"
)

type UserCampaignServiceInterface interface {
	CreateUserCampaign(request *common.CampaignUserRequest) (string, error)
}

type CreateUserCampaignResponse struct {
	Time string `json:"time"`
}

type UserCampaignHandler struct {
	service UserCampaignServiceInterface
}

func NewUserCampaignHandler(s UserCampaignServiceInterface) *UserCampaignHandler {
	return &UserCampaignHandler{
		service: s,
	}
}

func (h *UserCampaignHandler) CreateUserCampaign(w http.ResponseWriter, r *http.Request) {
	// request
	var request common.CampaignUserRequest
	err := helpers.CreateAndValidateFromRequest(r, &request)
	if err != nil {
		helpers.HandleRequestError(err, w)
		return
	}

	// run business logic
	time, err := h.service.CreateUserCampaign(&request)
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
