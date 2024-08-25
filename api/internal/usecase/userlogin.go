package usecase

import (
	"github.com/dodokey/go-chat-app/internal/application"
	"github.com/dodokey/go-chat-app/internal/infrastructure"
	"github.com/dodokey/go-chat-app/internal/interfaces"
	"github.com/go-redis/redis/v8"
)

type UserLoginInput struct {
    Username string
    Password string
}

type UserLoginOutput struct {
    Token string
    Result interfaces.LoginResult
    Error error
}

func Userlogin(input UserLoginInput) (UserLoginOutput, error) {
    client := redis.NewClient(&redis.Options{Addr: "redis:6379"})
    userRepo := infrastructure.NewRedisUserRepository(client)
    userService := application.NewUserService(userRepo)
    userHandler := interfaces.NewUserHandler(userService)

    // Call the Login function with the correct arguments
    result, err := userHandler.Login(input.Username, input.Password)
    return UserLoginOutput{Token: "token1", Result: result , Error: err}, nil
}