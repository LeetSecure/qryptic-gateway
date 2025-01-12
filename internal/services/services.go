package services

import (
	"github.com/leetsecure/qryptic-gateway/internal/config"
	"github.com/leetsecure/qryptic-gateway/internal/externalcomms"
	"github.com/leetsecure/qryptic-gateway/internal/models"
	"github.com/leetsecure/qryptic-gateway/internal/utils/auth"
	"github.com/leetsecure/qryptic-gateway/internal/utils/config_store_util"
	"github.com/leetsecure/qryptic-gateway/internal/utils/wireguard"
)

func AddPeers(wgServerPeerConfigs []models.WGServerPeerConfig) error {
	for _, wgServerPeerConfig := range wgServerPeerConfigs {
		err := wireguard.WireguardAddPeer(wgServerPeerConfig)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeletePeers(wgServerPeerConfigs []models.WGServerPeerConfig) error {
	for _, wgServerPeerConfig := range wgServerPeerConfigs {
		err := wireguard.WireguardDeletePeer(wgServerPeerConfig)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateInterface(wgServerConfigStore models.WGServerConfigStore) error {
	var err error
	interfaceUpdated := changesInWireguardInterface(wgServerConfigStore.WGServerInterfaceConfig)
	if interfaceUpdated {
		// if config.ConfigStore.InitSetupDone {
		// 	err = wireguard.WireguardStop()
		// 	if err != nil {
		// 		return err
		// 	}
		// }
		err = wireguard.WireguardSetInterface(wgServerConfigStore.WGServerInterfaceConfig)
		if err != nil {
			return err
		}
		err = wireguard.WireguardRestart()
		if err != nil {
			return err
		}
	}
	return nil
}

func FullRestart() error {
	config.ConfigStore.InitSetupDone = false
	err := SyncVpnGatewayConfig()
	if err != nil {
		return err
	}
	return nil
}

func SyncVpnGatewayConfig() error {
	vpnGatewayUuid := config.ConfigStore.VpnGatewayUuid
	controllerConfigUrl := config.ConfigStore.ControllerVGWConfigUrlEndpoint
	authToken, err := auth.CreateVpnGatewayToken()

	if err != nil {
		return err
	}
	wgServerConfigStore, err := externalcomms.FetchWgConfigFromController(vpnGatewayUuid, controllerConfigUrl, authToken)
	if err != nil {
		return err
	}
	config_store_util.UpdateWgConfigStore(wgServerConfigStore)
	UpdateInterface(wgServerConfigStore)
	AddPeers(wgServerConfigStore.WGServerPeerConfigs)
	config.ConfigStore.InitSetupDone = true
	config_store_util.UpdateConfigStore(config.ConfigStore)
	return nil
}

func changesInWireguardInterface(wgServerInterfaceConfig models.WGServerInterfaceConfig) bool {
	if !config.ConfigStore.InitSetupDone {
		return true
	}
	return false
}
