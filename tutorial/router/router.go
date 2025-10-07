package router

import (
	"github.com/adarshk007/tutorial/handler"
	"github.com/adarshk007/tutorial/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	searchSvc := &service.SearchService{}
	searchHandler := &handler.SearchHandler{Service: searchSvc}

	api := r.Group("/api")
	{
		api.GET("/search", searchHandler.SearchAll)
		// Add more routes: /index, /bulk, etc.
	}

	return r
}
