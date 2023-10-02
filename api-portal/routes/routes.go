package routes

import (
	"fmt"
	"log"
	"net/http"

	"stock-api/api-portal/routes/health_handler"
	"stock-api/api-portal/routes/stock_handler"
	"stock-api/global"

	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// gin-swagger middleware
	_ "stock-api/docs"
)

func Init() {
	var port = global.Config.ServerPort
	router := gin.New()
	router.Use()

	gin.SetMode(gin.DebugMode)
	// register our routes
	health_handler.RegisterRoutes(router)
	stock_handler.RegisterRoutes(router)

	// Serve Swagger UI at /swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fmt.Println("API gateway listening on port: " + port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("Fatal Can't ListenAndServe: ", err)
	}
}
