package jwtfuncs

import (
	"byte-battle_backend/pkg/loggers"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	authorized bool
	username   string
	email      string
	role       int8
}

func VerifyJWT(tokenString string) (*Claims, error) {

	/**

	This function verifies the JWT and extracts claims from it.

	*/

	claims := jwt.MapClaims{}

	secretKey := os.Getenv("JWT_KEY")
	if secretKey == "" {
		loggers.VariableNotFound("JWT_KEY")
		return nil, fmt.Errorf("Could not find the JWT_KEY")
	}

	// Parse the JWT

	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("Could not parse the JWT")
	}

	// Assert the provided variable's types

	authorized, ok := claims["authorized"].(bool)
	if !ok {
		return nil, fmt.Errorf("AUTHORIZED claim did not pass the type assertion")
	}
	username, ok := claims["username"].(string)
	if !ok {
		return nil, fmt.Errorf("USERNAME claim did not pass the type assertion")
	}
	email, ok := claims["email"].(string)
	if !ok {
		return nil, fmt.Errorf("EMAIL claim did not pass the type assertion")
	}

	// Convert ROLE claim into int8

	extractedRole, ok := claims["role"].(int)
	var role int8
	if ok {
		role = int8(extractedRole)
	} else {
		return nil, fmt.Errorf("ROLE claim did not pass the type assertion")
	}

	extractedClaims := Claims{
		authorized: authorized,
		username:   username,
		email:      email,
		role:       role,
	}

	return &extractedClaims, nil
}
