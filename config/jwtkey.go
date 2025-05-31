package config

import (
	"log"
	"os"
)

var Jwtkey []byte 

func GetJwtKey() {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		log.Fatalf("Environmental variable not set")
	}
	Jwtkey = []byte(secretKey)
}