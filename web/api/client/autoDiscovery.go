package client

import (
	"github.com/gin-gonic/gin"
	"github.com/raymao96/komari/database/clients"
	"github.com/raymao96/komari/database/notification"
	"github.com/raymao96/komari/pkg/config"
	"github.com/raymao96/komari/utils"
	logger "github.com/raymao96/komari/utils/log"
	"github.com/raymao96/komari/web/api"
)

func RegisterClient(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		api.RespondError(c, 403, "Invalid AutoDiscovery Key")
		return
	}
	AutoDiscoveryKey, err := config.GetAs[string](config.AutoDiscoveryKeyKey, "")
	if err != nil {
		api.RespondError(c, 500, "Failed to get AutoDiscovery Key: "+err.Error())
		return
	}
	if AutoDiscoveryKey == "" ||
		len(AutoDiscoveryKey) < 12 ||
		"Bearer "+AutoDiscoveryKey != auth {

		api.RespondError(c, 403, "Invalid AutoDiscovery Key")
		return
	}
	name := c.Query("name")
	if name == "" {
		name = utils.GenerateRandomString(8)
	}
	name = "Auto-" + name
	uuid, token, err := clients.CreateClientWithName(name)
	if err != nil {
		api.RespondError(c, 500, "Failed to create client: "+err.Error())
		return
	}
	if err := notification.AddDefaultOnClientUUID(uuid); err != nil {
		logger.ErrorArgs("clients", "Failed to apply default-on load notifications to auto-discovered client:", err)
	}
	api.RespondSuccess(c, gin.H{"uuid": uuid, "token": token})
}
