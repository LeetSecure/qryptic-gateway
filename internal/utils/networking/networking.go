package networking

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/leetsecure/qryptic-gateway/internal/utils/logger"
)

func InitialNetworkSetupForWireguard() error {
	log := logger.Default()

	// Step 1: Enable IP forwarding
	err := enableIPForwarding()
	if err != nil {
		log.Errorf("Failed to enable IP forwarding: %v", err)
		return err
	}

	// Step 2: Get the default network interface
	defaultInterface, err := getDefaultInterface()
	if err != nil {
		log.Errorf("Failed to get default interface: %v", err)
		return err
	}

	log.Infof("Default Interface: %s\n", defaultInterface)

	// Step 3: Configure iptables rules
	err = configureIptables(defaultInterface)
	if err != nil {
		log.Errorf("Failed to configure iptables: %v", err)
		return err
	}

	log.Info("Network configuration completed successfully.")
	return nil
}

// Function to enable IP forwarding
func enableIPForwarding() error {
	// Command to append to /etc/sysctl.conf
	cmd := `echo "net.ipv4.ip_forward=1" >> /etc/sysctl.conf && echo "net.ipv6.conf.all.forwarding=1" >> /etc/sysctl.conf && sysctl -p`
	// Run the command using bash
	err := exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		return fmt.Errorf("failed to enable IP forwarding: %v", err)
	}
	return nil
}

// Function to get the default network interface
func getDefaultInterface() (string, error) {
	log := logger.Default()
	// Command to get default interface
	cmd := `ip route | grep default | awk '{print $5}'`
	// Run the command using bash
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", fmt.Errorf("failed to get default interface: %v", err)
	}
	// Trim any trailing whitespace
	outString := strings.TrimSpace(string(out))
	log.Infof("default interface is : %s", outString)
	return outString, nil
}

// Function to configure iptables
func configureIptables(defaultInterface string) error {
	// List of iptables commands
	commands := []string{
		"iptables -A FORWARD -i wg0 -j ACCEPT",
		"iptables -A FORWARD -o wg0 -j ACCEPT",
		fmt.Sprintf("iptables -t nat -A POSTROUTING -o %s -j MASQUERADE", defaultInterface),
	}

	// Execute each iptables command
	for _, cmd := range commands {
		err := exec.Command("bash", "-c", cmd).Run()
		if err != nil {
			return fmt.Errorf("failed to run command '%s': %v", cmd, err)
		}
	}
	return nil
}
