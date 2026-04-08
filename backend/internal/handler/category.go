package handler

import (
	"net/http"

	"github.com/amesupakorn/pocket-ledger/internal/services"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{
		service: services.NewCategoryService(),
	}
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {
	data, err := h.service.GetCategories(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
