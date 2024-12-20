package adminToken

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

const SecretAdminKey = "HJsGYdUDF!DN^Bdd$%asj*_da#sd!has$$ash#dasd&%^$@"

type AdminClaims struct {
	AdminID uint `json:"admin_id"`
	jwt.StandardClaims
}

func GenerateAdminToken(adminID uint) (string, error) {
	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &AdminClaims{
		AdminID: adminID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretAdminKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateAdminToken(tokenString string) (*AdminClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretAdminKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*AdminClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
