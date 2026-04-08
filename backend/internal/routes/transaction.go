package routes

import (
	handlers "github.com/amesupakorn/pocket-ledger/internal/handler"
	"github.com/gin-gonic/gin"
)

func TransactionRoute(r *gin.Engine) {

	h := handlers.NewTransactionHandler()

	r.POST("/create/transaction", h.CreateTransaction)
	r.GET("/transactions", h.GetTransactions)
	r.PUT("/transaction/:id", h.UpdateTransaction)
	r.DELETE("/transaction/:id", h.DeleteTransaction)
}
