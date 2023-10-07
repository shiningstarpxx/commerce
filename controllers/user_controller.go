package controllers

import (
	"encoding/json"
	"net/http"

	"commerce/datalayer"
	"commerce/metrics"
)

type UserController struct {
	Datalayer datalayer.UserDatalayerInterface
}

// HttpRequestCreateUser is the data that should be provided in a request to create a new user.
type HttpRequestCreateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type HttpResponseUpdateUser struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "Method not allowed"}`))
		return
	}

	var requestData HttpRequestCreateUser
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid request data"}`))
		return
	}

	user, err := u.Datalayer.CreateUser(requestData.Username, requestData.Password, requestData.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Failed to create user"}`))
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (u *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	metrics.GetUserCount.WithLabelValues(string(status)).Inc()
}

func (u *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
}

func (u *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
}

func (u *UserController) UpdateUserEmailByIDWithTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "Method not allowed"}`))
		return
	}

	var requestData HttpResponseUpdateUser
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid request data"}`))
		return
	}

	err = u.Datalayer.UpdateUserEmailByIDWithTransaction(requestData.ID, requestData.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Failed to create user"}`))
		return
	}
}

func (u *UserController) UpdateUserEmailByIDWithCAS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "Method not allowed"}`))
		return
	}

	var requestData HttpResponseUpdateUser
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid request data"}`))
		return
	}

	err = u.Datalayer.UpdateUserEmailByIDWithTransaction(requestData.ID, requestData.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Failed to create user"}`))
		return
	}
}
