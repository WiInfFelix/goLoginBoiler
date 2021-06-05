package auth

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTWrapper struct {
	SecretKey   string
	Issuer      string
	Expirations int
}

type JWTClaim struct {
	Email string
	jwt.StandardClaims
}

func (j *JWTWrapper) GenerateJWT(email string) (signedToken string, err error) {
	claims := &JWTClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 2).Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(j.SecretKey))

	if err != nil {
		return
	}

	return
}

func ValidateToken(signedToken string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte("weyo"), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)

	if !ok {
		log.Println("Couldnt parse claims")
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return nil, err
	}
	return claims, err
}
