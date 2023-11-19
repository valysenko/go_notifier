package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_notifier/internal/dto"
	"go_notifier/internal/service"
	"io"
	"net/http"
)

type CreateUserCampaignResponse struct {
	Time string `json:"time"`
}

func CreateUserCampaignHandler(w http.ResponseWriter, r *http.Request) {
	// request
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	requestData := buf.Bytes()
	r.Body.Close()

	var dto dto.CampaignUser
	if err := json.Unmarshal(requestData, &dto); err != nil {
		http.Error(w, "ivalid json unmarshal", http.StatusInternalServerError)
		return
	}

	err = dto.Validate()
	if err != nil {
		http.Error(w, "not valid", http.StatusBadRequest)
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
