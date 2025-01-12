package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leetsecure/qryptic-gateway/internal/config"
	"github.com/leetsecure/qryptic-gateway/internal/routes"
	"github.com/leetsecure/qryptic-gateway/internal/services"
	"github.com/leetsecure/qryptic-gateway/internal/utils/config_store_util"
	"github.com/leetsecure/qryptic-gateway/internal/utils/logger"
	"github.com/leetsecure/qryptic-gateway/internal/utils/networking"
)

func main() {
	log := logger.Default()
	err := config_store_util.InitialConfigStoreSetup()
	if err != nil {
		log.Error(err)
		return
	}
	err = services.SyncVpnGatewayConfig()
	if err != nil {
		log.Error(err)
		return
	}
	err = networking.InitialNetworkSetupForWireguard()
	if err != nil {
		log.Error(err)
		return
	}

	router := gin.Default()

	//Setup the routes in the router
	routes.SetupControllerRoutes(router)

	// Start the server
	log.Info("VPN Gateway Service Starting")
	router.Run(":" + config.ConfigStore.ApplicationPort)
}
