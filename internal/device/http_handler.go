package device

import (
	"encoding/json"
	"go_notifier/internal/common"
	"go_notifier/internal/http/helpers"
	"net/http"
)

type DeviceServiceInterface interface {
	CreateDevice(request *common.DeviceRequest) (int64, error)
}

type CreateDeviceResponse struct {
	Token string `json:"token"`
}

type DeviceHandler struct {
	service DeviceServiceInterface
}

func NewDeviceHandler(service DeviceServiceInterface) *DeviceHandler {
	return &DeviceHandler{
		service: service,
	}
}

func (h *DeviceHandler) CreateDevice(w http.ResponseWriter, r *http.Request) {
	// request
	var request common.DeviceRequest
	err := helpers.CreateAndValidateFromRequest(r, &request)
	if err != nil {
		helpers.HandleRequestError(err, w)
		return
	}

	// run business logic
	_, err = h.service.CreateDevice(&request)
	if err != nil {
		http.Error(w, "error while device creation", http.StatusInternalServerError)
		return
	}

	// response
	responseData := CreateDeviceResponse{Token: request.Token}
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
