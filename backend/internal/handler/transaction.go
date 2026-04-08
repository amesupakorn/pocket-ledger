package handlers

import (
	"net/http"

	"github.com/amesupakorn/pocket-ledger/internal/dto"
	"github.com/amesupakorn/pocket-ledger/internal/services"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler() *TransactionHandler {
	return &TransactionHandler{
		service: services.NewTransactionService(),
	}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var req dto.CreateTransactionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := int64(1) // mock

	id, err := h.service.CreateTransaction(c.Request.Context(), userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "created",
	})
}

func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	data, err := h.service.GetTransactions(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}
