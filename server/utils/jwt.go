package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const secretKey = "supersecret"

func GenerateToken(id int64, role string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":  id,
        "role": role,
        "exp": time.Now().Add(time.Hour * 24).Unix(), 
    })
    return token.SignedString([]byte(secretKey))
}



func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Beklenmeyen imzalama yöntemi: %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("Geçersiz token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("token doğrulanamadı")
}
