package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leetsecure/qryptic-gateway/internal/handlers"
	"github.com/leetsecure/qryptic-gateway/internal/middlewares"
)

func SetupControllerRoutes(r *gin.Engine) {

	r.GET("/health", handlers.HealthCheck)
	controllerGroup := r.Group("/controller")
	{
		controllerGroup.POST("/add-peers", middlewares.ControllerAuthCheckMiddleware, handlers.AddPeers)
		controllerGroup.POST("/delete-peers", middlewares.ControllerAuthCheckMiddleware, handlers.DeletePeers)
		controllerGroup.POST("/sync-vpn-gateway-config", middlewares.ControllerAuthCheckMiddleware, handlers.SyncVpnGatewayConfig)
		controllerGroup.POST("/restart", middlewares.ControllerAuthCheckMiddleware, handlers.FullRestart)
	}

}
