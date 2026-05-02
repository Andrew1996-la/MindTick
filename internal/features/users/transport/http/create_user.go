package users_transport_http

import (
	"encoding/json"
	"net/http"
)

type createUserRequest struct {
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

type createUserResponse struct {
	ID          int    `json:"id"`
	Version     int    `json:"version"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

func (u *UserHttpHandler) CreateUser(rw http.ResponseWriter, req http.Request) {
	var user createUserResponse

	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		
	}
}
