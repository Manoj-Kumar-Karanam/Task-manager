package utils

import (
	"errors"
	"fmt"
	"task_manager/config"

	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Validate the signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return config.Jwtkey, nil
    })

    if err != nil {
        return nil, fmt.Errorf("token parse error: %v", err)
    }

    // Check token validity
    if !token.Valid {
        return nil, errors.New("invalid token")
    }

    // Assert and return claims
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return nil, errors.New("could not parse claims")
    }

    return claims, nil
}
