package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// routes
	TransactionRoute(r)

	return r
}
