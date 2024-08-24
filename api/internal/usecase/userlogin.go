package usecase

import (
	"net/http"

	"github.com/dodokey/go-chat-app/internal/application"
	"github.com/dodokey/go-chat-app/internal/infrastructure"
	"github.com/dodokey/go-chat-app/internal/interfaces"

	"github.com/go-redis/redis/v8"
)

func Userlogin(w http.ResponseWriter, r *http.Request) {
	client := redis.NewClient(&redis.Options{Addr: "redis:6379"})
	userRepo := infrastructure.NewRedisUserRepository(client)
	userService := application.NewUserService(userRepo)
	userHandler := interfaces.NewUserHandler(userService)
	userHandler.Login(w, r)
}