package routes

import (
	"exchange/services"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	api := router.Group("/apis/v1")
	{
		api.GET("/exchange-rates", services.ExchangeRate)
	}
	return router
}
