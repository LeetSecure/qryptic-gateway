package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leetsecure/qryptic-gateway/internal/models"
	"github.com/leetsecure/qryptic-gateway/internal/services"
)

func SyncVpnGatewayConfig(c *gin.Context) {
	err := services.SyncVpnGatewayConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func FullRestart(c *gin.Context) {
	err := services.FullRestart()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func AddPeers(c *gin.Context) {
	var wgServerPeerConfigs []models.WGServerPeerConfig
	if err := c.ShouldBindJSON(&wgServerPeerConfigs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := services.AddPeers(wgServerPeerConfigs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})

}

func DeletePeers(c *gin.Context) {
	var wgServerPeerConfigs []models.WGServerPeerConfig
	if err := c.ShouldBindJSON(&wgServerPeerConfigs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := services.DeletePeers(wgServerPeerConfigs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}
