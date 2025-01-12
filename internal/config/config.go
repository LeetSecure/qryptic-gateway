package config

import (
	"time"

	"github.com/leetsecure/qryptic-gateway/internal/models"
)

const JwtTokenTimeout = time.Hour
const ConfigStoreFilePath = "./config.json"

var ConfigStore = models.ConfigStore{}

var WireguardConfig = models.WGServerConfigStore{}
