package config_store_util

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/leetsecure/qryptic-gateway/internal/config"
	"github.com/leetsecure/qryptic-gateway/internal/models"
	"github.com/leetsecure/qryptic-gateway/internal/utils/logger"
)

func checkConfigFileExists(filepath string) (bool, error) {
	log := logger.Default()
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		log.Errorf("%s doesnot exists", filepath)
		return false, nil
	}
	if err != nil {
		log.Errorf("Error while checking file existence - %s", err)
		return false, err
	}
	return true, nil
}

func setConfigToFile(configstore models.ConfigStore) error {
	log := logger.Default()
	filepath := config.ConfigStoreFilePath
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Errorf("Error while opening file - %s", filepath)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(configstore)
	if err != nil {
		log.Errorf("Error while writing in file - %s", filepath)
		return err
	}
	return nil
}

func getConfigfromFile() (models.ConfigStore, error) {
	log := logger.Default()
	filepath := config.ConfigStoreFilePath
	exists, err := checkConfigFileExists(filepath)
	if err != nil {
		return models.ConfigStore{}, err
	}
	if !exists {
		return models.ConfigStore{}, errors.New("config file doesnot exists")
	}
	var configstore models.ConfigStore

	file, err := os.Open(filepath)
	if err != nil {
		log.Errorf("Error while opening file - %s", filepath)
		return models.ConfigStore{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configstore)
	if err != nil {
		log.Errorf("Error in decoding configs file to configstore - %s", filepath)
		return models.ConfigStore{}, err
	}
	return configstore, nil
}

func SyncConfigStoreFromFile() (models.ConfigStore, error) {
	configStore, err := getConfigfromFile()
	if err != nil {
		return models.ConfigStore{}, err
	}
	config.ConfigStore = configStore
	return configStore, nil
}

func UpdateConfigStore(configStore models.ConfigStore) error {
	err := setConfigToFile(configStore)
	if err != nil {
		return err
	}
	config.ConfigStore = configStore
	return nil
}

func InitialConfigStoreSetup() error {

	vpnGatewayUuid, uExists := os.LookupEnv("VpnGatewayUuid")
	vpnGatewayControllerJWTSecretKey, jpExists := os.LookupEnv("VpnGatewayControllerJWTSecretKey")
	vpnGatewayControllerJWTAlgorithm, jaExists := os.LookupEnv("VpnGatewayControllerJWTAlgorithm")
	controllerVGWConfigUrlEndpoint, vceExists := os.LookupEnv("ControllerVGWConfigUrlEndpoint")
	applicationPort, appPortExists := os.LookupEnv("ApplicationPort")

	if !uExists || !jpExists || !jaExists || !vceExists {
		return errors.New("required env variables not present")
	}

	if !appPortExists {
		applicationPort = "8080"
	}

	initConfigStore := models.ConfigStore{
		VpnGatewayUuid:                   vpnGatewayUuid,
		VpnGatewayControllerJWTSecretKey: vpnGatewayControllerJWTSecretKey,
		VpnGatewayControllerJWTAlgorithm: vpnGatewayControllerJWTAlgorithm,
		ControllerVGWConfigUrlEndpoint:   controllerVGWConfigUrlEndpoint,
		ApplicationPort:                  applicationPort,
	}
	err := UpdateConfigStore(initConfigStore)
	if err != nil {
		return err
	}
	return nil
}

func UpdateWgConfigStore(newWgServerConfigStore models.WGServerConfigStore) {
	config.WireguardConfig.WGServerInterfaceConfig = newWgServerConfigStore.WGServerInterfaceConfig
	config.WireguardConfig.WGServerPeerConfigs = append(config.WireguardConfig.WGServerPeerConfigs, newWgServerConfigStore.WGServerPeerConfigs...)
}
