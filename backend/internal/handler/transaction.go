package handlers

import (
	"net/http"
	"strconv"

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

func (h *TransactionHandler) DeleteTransaction(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.ParseInt(idParam, 10, 64)

	err := h.service.DeleteTransaction(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "deleted"})
}

func (h *TransactionHandler) UpdateTransaction(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.ParseInt(idParam, 10, 64)

	var req dto.UpdateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.service.UpdateTransaction(c.Request.Context(), id, req.Amount, req.Note)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"message": "updated"})
}
