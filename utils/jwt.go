package utils
import (
	"time"
	"task_manager/config"
	"github.com/dgrijalva/jwt-go"
)


// var Jwtkey = []byte(secret)

func GenerateTokens(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id" : userId,
		"exp" : time.Now().Add(48 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	})
	tokenString, err := token.SignedString(config.Jwtkey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}