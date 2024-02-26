package user_app

import (
	"encoding/json"
	"go_notifier/internal/app/http/helpers"
	"go_notifier/internal/common"
	"net/http"
)

type UserAppServiceInterface interface {
	CreateUserApp(request *common.UserAppRequest) (int64, error)
}

type CreateUserAppResponse struct {
	Identifier string `json:"identitier"`
	Type       string `json:"type"`
}

type UserAppHandler struct {
	service UserAppServiceInterface
}

func NewUserAppHandler(service UserAppServiceInterface) *UserAppHandler {
	return &UserAppHandler{
		service: service,
	}
}

func (h *UserAppHandler) CreateUserApp(w http.ResponseWriter, r *http.Request) {
	// request
	var request common.UserAppRequest
	err := helpers.CreateAndValidateFromRequest(r, &request)
	if err != nil {
		helpers.HandleRequestError(err, w)
		return
	}

	// run business logic
	_, err = h.service.CreateUserApp(&request)
	if err != nil {
		http.Error(w, "error while device creation", http.StatusInternalServerError)
		return
	}

	// response
	responseData := CreateUserAppResponse{Identifier: request.Identifier, Type: request.Type}
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
