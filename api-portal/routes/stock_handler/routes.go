package stock_handler

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	router.GET("/api/stocks", GetStocks)
	router.POST("/api/stocks", CreateStock)
	router.GET("/api/stocks/:id", GetStockByID)
	router.PATCH("/api/stocks/:id", UpdateStock)
	router.DELETE("/api/stocks/:id", DeleteStock)
}
