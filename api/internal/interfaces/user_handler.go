package interfaces

import (
	"encoding/json"
	"net/http"

	"github.com/dodokey/go-chat-app/internal/application"
)

type UserHandler struct {
    service *application.UserService
}

type LoginResult struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
    Token   string `json:"token,omitempty"`
}

func NewUserHandler(service *application.UserService) *UserHandler {
    return &UserHandler{service: service}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
    var creds struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err := h.service.Login(creds.Username, creds.Password)
    if err != nil {
        result:= LoginResult{
            Success: false,
            Message: err.Error(),
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(result)
        return
    }

    result := LoginResult{
        Success: true,
        Message: "login successful",
        Token: "123",
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}