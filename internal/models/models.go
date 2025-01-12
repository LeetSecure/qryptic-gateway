package models

type ConfigStore struct {
	InitSetupDone                    bool   `json:"initSetupDone"`
	ApplicationPort                  string `json:"applicationPort"`
	VpnGatewayUuid                   string `json:"vpnGatewayUuid"`
	VpnGatewayControllerJWTSecretKey string `json:"vpnGatewayControllerJWTSecretKey"`
	VpnGatewayControllerJWTAlgorithm string `json:"vpnGatewayControllerJWTAlgorithm"`
	ControllerVGWConfigUrlEndpoint   string `json:"controllerVGWConfigUrlEndpoint"`
	VpnGatewayServerIP               string `json:"vpnGatewayServerIP"`
	IPAddress                        string `json:"ipAddress"`
	WireguardPort                    int    `json:"wireguardPort"`
	PresharedKey                     string `json:"preSharedKey"`
	DnsServer                        string `json:"dnsServer"`
	VpnGatewayWgPublicKey            string `json:"vpnGatewayWgPublicKey"`
	VpnGatewayWgPrivateKey           string `json:"vpnGatewayWgPrivateKey"`
	WGPresharedKey                   string `json:"wgPresharedKey"`
	NetworkInterface                 string `json:"networkInterface"`
	PostUp                           string `json:"postUp"`
	PostDown                         string `json:"postDown"`
}

type WGServerInterfaceConfig struct {
	VpnGatewayUuid string `json:"vpnGatewayUuid"`
	PublicKey      string `json:"publicKey"`
	PrivateKey     string `json:"privateKey"`
	IPAddress      string `json:"ipAddress"`
	ListenPort     int    `json:"listenPort"`
	PostUp         string `json:"postUp"`
	PostDown       string `json:"postDown"`
	DnsServer      string `json:"dnsServer"`
}

type WGServerPeerConfig struct {
	ClientAllowedIPs string `json:"clientAllowedIPs"`
	ClientPublicKey  string `json:"clientPublicKey"`
	PresharedKey     string `json:"presharedKey"`
}

type WGServerConfigStore struct {
	WGServerInterfaceConfig WGServerInterfaceConfig `json:"wgServerInterfaceConfig"`
	WGServerPeerConfigs     []WGServerPeerConfig    `json:"wgServerPeerConfigs"`
}
