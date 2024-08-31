package main

import (
	"log"
	"net/http"

	"github.com/dodokey/go-chat-app/internal/httphandler"
	"github.com/dodokey/go-chat-app/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(httphandler.CORSMiddleware())

    if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
        log.Fatalf("could not set trusted proxies: %s", err)
    }

    router.GET("/", func(c *gin.Context) {c.String(http.StatusOK, "Hello, Go-Chat-Apps!")})

	router.POST("/login", httphandler.ResponseWrapper(func(c *gin.Context) (interface{}, error) {
		var input usecase.UserLoginInput
		if err := c.ShouldBindJSON(&input); err != nil {
			return nil, err
		}
		return usecase.Userlogin(input)
	}))

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("could not start server: %s", err)
	}
}