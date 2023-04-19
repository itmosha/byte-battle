package jwtfuncs

import (
	"byte-battle_backend/pkg/loggers"
	"os"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(username string, email string, role int8) (string, error) {

	/**

	This function makes a JWT with username, email and role provided in the payload.

	*/

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["email"] = email
	claims["role"] = role

	secretKey := os.Getenv("JWT_KEY")
	if secretKey == "" {
		loggers.VariableNotFound("JWT_KEY")
		return "", nil
	}

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
