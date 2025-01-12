package wireguard

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/leetsecure/qryptic-gateway/internal/models"
	"github.com/leetsecure/qryptic-gateway/internal/utils/logger"
)

const wgInterfaceName = "wg0"
const wgConfigDirectory = "/etc/wireguard"
const wgConfigPath = "/etc/wireguard/wg0.conf"

func createWireguardConfFile() error {
	log := logger.Default()
	err := os.MkdirAll(wgConfigDirectory, 0700)
	if err != nil {
		log.Errorf("Error in creating directory %s", wgConfigDirectory)
		return err
	}
	_, err = os.Create(wgConfigPath)
	if err != nil {
		log.Errorf("Error in creating file %s", wgConfigPath)
		return err
	}
	return nil
}

func WireguardSetInterface(wgServerInterfaceConfig models.WGServerInterfaceConfig) error {
	log := logger.Default()
	createWireguardConfFile()
	wgServerInterfaceConfigText := fmt.Sprintf(`
[Interface]
PrivateKey = %s
Address = %s
ListenPort = %d
DNS = %s	
# Add peers below
`, wgServerInterfaceConfig.PrivateKey, wgServerInterfaceConfig.IPAddress, wgServerInterfaceConfig.ListenPort, wgServerInterfaceConfig.DnsServer)

	wgServerInterfaceConfigTextBytes := []byte(wgServerInterfaceConfigText)

	err := os.WriteFile(wgConfigPath, wgServerInterfaceConfigTextBytes, 0644)
	if err != nil {
		log.Errorf("Error in writing to [%s] - %s", wgConfigPath, err)
		return err
	}

	return nil
}

func WireguardStart() error {
	log := logger.Default()
	cmd := exec.Command("bash", "-c", "wg-quick up "+wgInterfaceName)
	if err := cmd.Run(); err != nil {
		log.Errorf("Error in starting Wireguard - %s", err)
		return err
	}
	return nil
}

func WireguardAddPeer(peerConfig models.WGServerPeerConfig) error {
	log := logger.Default()
	cmd := exec.Command("wg", "set", wgInterfaceName, "peer", peerConfig.ClientPublicKey, "allowed-ips", peerConfig.ClientAllowedIPs)
	if err := cmd.Run(); err != nil {
		log.Errorf("Error in adding Peer [%s] - %s", peerConfig.ClientPublicKey, err)
		return err
	}
	return nil
}

func WireguardDeletePeer(peerConfig models.WGServerPeerConfig) error {
	log := logger.Default()
	cmd := exec.Command("wg", "set", wgInterfaceName, "peer", peerConfig.ClientPublicKey, "remove")
	if err := cmd.Run(); err != nil {
		log.Errorf("Error in deleting Peer [%s] - %s", peerConfig.ClientPublicKey, err)
		return err
	}
	return nil
}

func GetNetworkInterface() (string, error) {
	return "", nil
}

func WireguardRestart() error {

	err := WireguardStop()
	if err != nil {
		return err
	}
	err = WireguardStart()
	if err != nil {
		return err
	}
	return nil
}

func WireguardStop() error {
	log := logger.Default()
	cmd := exec.Command("bash", "-c", "wg-quick down "+wgInterfaceName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		// Check for specific error messages or codes if needed
		if string(output) == "wg-quick: "+wgInterfaceName+" does not exist" {
			log.Errorf("WG interface [%s] does not exist", wgInterfaceName)
			return nil
		} else if strings.Contains(string(output), "is not a WireGuard interface") {
			log.Errorf("WG interface [%s] is already stopped", wgInterfaceName)
			return nil
		} else {
			log.Errorf("Error in stopping Wireguard - %s", err)
			return err
		}
	}
	return nil
}
