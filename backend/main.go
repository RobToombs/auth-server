package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AuthCode struct {
	Code string `json:"code"`
}

type AuthToken struct {
	IdToken string `json:"id_token"`
}

func main() {
	router := gin.Default()
	router.POST("/auth", authenticate)

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	router.Run(":8080")
}

func authenticate(c *gin.Context) {
	newAlbum := AuthCode{}
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	if newAlbum.Code != "" {
		c.IndentedJSON(http.StatusOK, AuthToken{IdToken: "DUMMY_TOKEN"})
	} else {
		c.IndentedJSON(http.StatusBadRequest, AuthToken{})
	}
}
