package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
	UserId string
}

const SECRET_KEY = "secretkey123"
const TOKEN_EXP = int64(time.Hour * 3) // 3hours

func main() {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: TOKEN_EXP,
		},
		UserId: "User1",
	})
	token, err := t.SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(token)
	fmt.Println(getUserId(token))
	fmt.Println(tokenIsValid(token))

}

func getUserId(tokenStr string) string {
	claims := &Claims{}
	jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	return claims.UserId
}

func tokenIsValid(tokenStr string) bool {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}
