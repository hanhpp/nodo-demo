package health_handler

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.Engine) {
	route.GET("/status", Health)
}
