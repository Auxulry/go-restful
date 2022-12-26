package common

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type JwtMethod int8

const (
	HMAC JwtMethod = iota
	Ed25519
	ECDSA
	RSA
	RSAPSS
)

func JwtClaims(method jwt.SigningMethod, claims jwt.MapClaims, key []byte) (string, error) {
	token := jwt.NewWithClaims(method, claims)

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func JwtParse(method JwtMethod, tokenString string, key []byte) (*jwt.Token, error) {
	switch method {
	case HMAC:
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return key, nil
		})
		return token, err
	case Ed25519:
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return key, nil
		})
		return token, err
	case ECDSA:
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return key, nil
		})
		return token, err
	case RSA:
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return key, nil
		})
		return token, err
	case RSAPSS:
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSAPSS); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return key, nil
		})
		return token, err
	default:
		return &jwt.Token{}, errors.New("unknown method")
	}
}

func JwtValidate(method JwtMethod, tokenString string, key []byte) (jwt.MapClaims, error) {
	token, err := JwtParse(method, tokenString, key)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if ok {
			return claims, nil
		} else {
			return jwt.MapClaims{}, errors.New("token not valid")
		}
	} else {
		return jwt.MapClaims{}, err
	}
}
