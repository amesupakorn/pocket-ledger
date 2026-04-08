package routes

import (
	handlers "github.com/amesupakorn/pocket-ledger/internal/handler"
	"github.com/gin-gonic/gin"
)

func CategoryRoute(r *gin.Engine) {

	h := handlers.NewCategoryHandler()

	r.GET("/categories", h.GetCategories)

}
