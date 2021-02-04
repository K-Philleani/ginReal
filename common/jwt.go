package common

import (
	"ginReal/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claim struct {
	UserID uint
	jwt.StandardClaims
}

var jwtKey = []byte("a_secret_create")

func ReleaseToken(user model.User) (string, error){
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claim := &Claim{
		UserID:         user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "issuer_now",
			Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}


func ParseToken(tokenString string) (*jwt.Token, *Claim, error){
	claim := &Claim{}
	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (i interface{}, err error){
		return jwtKey, nil
	})
	return token, claim, err
}