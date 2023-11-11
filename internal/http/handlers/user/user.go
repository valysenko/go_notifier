package user

import (
	"bytes"
	"encoding/json"
	"go_notifier/internal/dto"
	"go_notifier/internal/service/user"
	"io"
	"net/http"
)

type ResponseData struct {
	ID int64 `json:"id"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// request
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	requestData := buf.Bytes()
	r.Body.Close()

	var dto dto.User
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
	userId, err := user.CreateUser(&dto)
	if err != nil {
		http.Error(w, "error while user creation", http.StatusInternalServerError)
		return
	}

	// response
	responseData := ResponseData{ID: userId}
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
