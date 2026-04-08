package routes

import (
	handlers "github.com/amesupakorn/pocket-ledger/internal/handler"
	"github.com/gin-gonic/gin"
)

func TransactionRoute(r *gin.Engine) {

	h := handlers.NewTransactionHandler()

	r.POST("/create/transactions", h.CreateTransaction)
	r.GET("/transactions", h.GetTransactions)

}
