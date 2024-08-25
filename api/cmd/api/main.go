package main

import (
	"log"
	"net/http"

	"github.com/dodokey/go-chat-app/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
	router.Use(corsMiddleware())

    if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
        log.Fatalf("could not set trusted proxies: %s", err)
    }

    router.GET("/", func(c *gin.Context) {c.String(http.StatusOK, "Hello, Go-Chat-Apps!")})

	router.POST("/login", ResponseWrapper(func(c *gin.Context) (interface{}, error) {
		var input usecase.UserLoginInput
		if err := c.ShouldBindJSON(&input); err != nil {
			return nil, err
		}

		result, err := usecase.Userlogin(input)
		if err != nil {
			return nil, err
		}

		return result, nil
	}))

	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("could not start server: %s", err)
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

func ResponseWrapper(handler func(*gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := handler(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
