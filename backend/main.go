package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
