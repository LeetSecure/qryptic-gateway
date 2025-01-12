package externalcomms

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/leetsecure/qryptic-gateway/internal/models"
	"github.com/leetsecure/qryptic-gateway/internal/utils/logger"
)

func FetchWgConfigFromController(vpnGatewayUuid, controllerConfigUrl, authToken string) (models.WGServerConfigStore, error) {

	log := logger.Default()

	spaceClient := http.Client{
		Timeout: time.Second * 5,
	}

	wgServerConfigStore := models.WGServerConfigStore{}

	req, err := http.NewRequest(http.MethodGet, controllerConfigUrl, nil)
	if err != nil {
		log.Errorf("Error in creating request for %s", controllerConfigUrl)
		return wgServerConfigStore, err
	}

	authTokenWithBearer := fmt.Sprintf("Bearer %s", authToken)
	req.Header.Set("Authorization", authTokenWithBearer)
	req.Header.Set("VPN-Gateway-UUID", vpnGatewayUuid)

	res, err := spaceClient.Do(req)
	if err != nil {
		log.Errorf("Error in executing request for %s", controllerConfigUrl)
		return wgServerConfigStore, err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Errorf("Error in reading response body of %s", controllerConfigUrl)
		return wgServerConfigStore, err
	}

	err = json.Unmarshal(body, &wgServerConfigStore)
	if err != nil {
		log.Errorf("Error in unmarshalling response body of %s", controllerConfigUrl)
		return wgServerConfigStore, err
	}

	return wgServerConfigStore, nil

}
