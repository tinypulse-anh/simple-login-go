package routes

import "github.com/gin-gonic/gin"

func Init() *gin.Engine {
	router := gin.Default()
	registerApiV1Routes(router)

	return router
}
