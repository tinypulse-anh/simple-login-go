package apiV1Sessions

import (
	"github.com/gin-gonic/gin"
	"github.com/tinypulse-anh/simple-login-go/models"
	"net/http"
)

const LOGIN_TYPE_SSO = "sso"
const LOGIN_TYPE_PASSWORD = "password"

// GET /api/v1/sessions/new
func New(c *gin.Context) {
	user := findUser(c.Param("username"))

	if user.UseSso() {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data": {
				"type":    LOGIN_TYPE_SSO,
				"sso_url": "https://www.tinypulse.com/",
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data": {
				"type": LOGIN_TYPE_PASSWORD,
			},
		})
	}
}

func findUser(username string) {
	var user models.User

	models.DB.Where("username = ? OR email = ?", username, username).First(&user)

	return user
}
