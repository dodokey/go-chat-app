package usecase

import (
	"time"

	"github.com/dodokey/go-chat-app/internal/application"
	"github.com/dodokey/go-chat-app/internal/infrastructure"
	"github.com/dodokey/go-chat-app/internal/interfaces"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
)

type UserLoginInput struct {
    Username string
    Password string
}

type UserLoginOutput struct {
    Data interfaces.LoginResult `json:"data"`
    Error error `json:"error"`
}

func Userlogin(input UserLoginInput) (UserLoginOutput, error) {
    client := redis.NewClient(&redis.Options{Addr: "redis:6379"})
    userRepo := infrastructure.NewRedisUserRepository(client)
    userService := application.NewUserService(userRepo)
    userHandler := interfaces.NewUserHandler(userService)

    result, err := userHandler.Login(input.Username, input.Password)
    if(err != nil) {
		return UserLoginOutput{}, err
    }
    token, err := GenerateJWT(input.Username)
    result.Token = token
    return UserLoginOutput{Data: result , Error: err}, nil
}


var jwtKey = []byte("secret_key_5566")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}