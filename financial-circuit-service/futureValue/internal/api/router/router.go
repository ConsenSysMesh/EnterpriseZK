package router

import (
	"futurevalue/internal/api/controllers"
	"futurevalue/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	app := gin.New()

	// Middlewares
	app.Use(gin.Recovery())
	app.NoRoute(middleware.NoRouteHandler())

	// Routes
	// ================== Proof Route
	app.POST("/prove", controllers.Proof)

	// ================== Verify Route
	app.POST("/verify", controllers.Verify)

	return app
}
