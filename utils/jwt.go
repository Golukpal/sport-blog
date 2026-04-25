package utils

import(
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("mysecret")

func GenerateJWT(id int )( string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(SecretKey)
}