package handler

import (
	"net/http"

	"github.com/adarshk007/tutorial/service"
	"github.com/gin-gonic/gin"
)

type SearchHandler struct {
	Service *service.SearchService
}

func (h *SearchHandler) SearchAll(c *gin.Context) {
	index := c.Query("index")
	if index == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "index query param is required"})
		return
	}

	result, err := h.Service.SearchAll(index)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
