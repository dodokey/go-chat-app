package interfaces

import (
	"github.com/dodokey/go-chat-app/internal/application"
)

type LoginResult struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
    Token   string `json:"token,omitempty"`
}

type UserHandler struct {
    service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
    return &UserHandler{service: service}
}

func (h *UserHandler) Login(username, password string) (LoginResult, error) {
    _, err := h.service.Login(username, password)
    if err != nil {
        return LoginResult{}, err
    }
    return LoginResult{
        Success: true,
        Message: "Login successful",
    }, nil
}