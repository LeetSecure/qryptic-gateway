package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/leetsecure/qryptic-gateway/internal/config"
	"github.com/leetsecure/qryptic-gateway/internal/utils/logger"
)

func CreateVpnGatewayToken() (string, error) {
	log := logger.Default()
	vpnGatewayUuid := config.ConfigStore.VpnGatewayUuid
	jwtAuthSecretKey := config.ConfigStore.VpnGatewayControllerJWTSecretKey
	jwtVpnGatewayAuthSecretKeyBytes := []byte(jwtAuthSecretKey)
	timeNow := time.Now()

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": vpnGatewayUuid,                             // Subject (user identifier)
		"iss": "gateway",                                  // Issuer
		"aud": "controller",                               // Audience (user role)
		"exp": timeNow.Add(config.JwtTokenTimeout).Unix(), // Expiration time
		"iat": timeNow.Unix(),                             // Issued at
	})

	tokenString, err := claims.SignedString(jwtVpnGatewayAuthSecretKeyBytes)
	if err != nil {
		log.Error("Error in creating signed jwt token for connecting to Controller")
		return "", err
	}
	return tokenString, nil
}

func VerifyControllerAuthToken(controllerAuthToken string) (string, error) {
	log := logger.Default()
	jwtAuthSecretKey := config.ConfigStore.VpnGatewayControllerJWTSecretKey
	jwtControllerAuthSecretKeyBytes := []byte(jwtAuthSecretKey)
	token, err := jwt.Parse(controllerAuthToken, func(token *jwt.Token) (interface{}, error) {
		return jwtControllerAuthSecretKeyBytes, nil
	})

	// Check for verification errors
	if err != nil {
		log.Error("Error in verifying jwt token")
		return "", err
	}

	// Check if the token is valid
	if !token.Valid {
		log.Info("Invalid jwt token received")
		return "", err
	}

	return "", nil
}
