package user

import (
	"context"
	"encoding/json"
	"net/http"
)

type UserController struct {
	repository UserRepositoryInterface
}

func (u *UserController) getUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user := User{
		ID:       "1",
		Name:     "John Doe",
		Email:    "john@example.com",
		password: "hashed_password",
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (u *UserController) getUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := context.Background()
	users, err := u.repository.GetAllUsers(ctx)
	if err != nil {
		http.Error(w, "Failed to get all users", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (u *UserController) Listen() {
	http.HandleFunc("/user", u.getUser)
	http.HandleFunc("/user/list", u.getUsers)
}
