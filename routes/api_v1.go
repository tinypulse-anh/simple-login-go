package routes

import (
	"github.com/gin-gonic/gin"
	sessionsApi "github.com/tinypulse-anh/simple-login-go/controllers/api/v1/sessions"
)

func registerApiV1Routes(router *gin.Engine) {
	root := router.Group("/api/v1")

	sessions_group := root.Group("/sessions")
	sessions_group.GET("/new", sessionsApi.New)
}
