package jwt

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestJwtClaims(t *testing.T) {
	t.Run("Test JWT Claims Success", func(t *testing.T) {
		token, err := JwtClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id":    "uuid",
			"expires_in": time.Now().Add(time.Duration(1) * time.Minute).Unix(),
		},
			[]byte("secret"))
		assert.Nil(t, err)
		assert.NotEqual(t, "", token)
	})
}

func TestJwtParse(t *testing.T) {
	t.Run("Test JWT Parse Success", func(t *testing.T) {
		token, err := JwtClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id":    "uuid",
			"expires_in": time.Now().Add(time.Duration(1) * time.Minute).Unix(),
		},
			[]byte("secret"))

		_, err = JwtParse(0, token, []byte("secret"))

		assert.Nil(t, err)
	})

	t.Run("Test JWT Parse Failed", func(t *testing.T) {
		token, err := JwtClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id":    "uuid",
			"expires_in": time.Now().Add(time.Duration(1) * time.Minute).Unix(),
		},
			[]byte("secret"))

		_, err = JwtParse(11, token, []byte("secret"))

		assert.NotNil(t, err)
		assert.Equal(t, "unknown method", err.Error())
	})
}

func TestJwtValidate(t *testing.T) {
	t.Run("Test JWT Validate Success", func(t *testing.T) {
		expiresIn := time.Now().Add(time.Duration(1) * time.Minute).Unix()
		token, err := JwtClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id":    "uuid",
			"expires_in": expiresIn,
		},
			[]byte("secret"))

		claims, err := JwtValidate(0, token, []byte("secret"))

		assert.Nil(t, err)
		assert.Equal(t, "uuid", claims["user_id"])
		assert.Equal(t, expiresIn, int64(claims["expires_in"].(float64)))
	})

	t.Run("Test JWT Validate Failed", func(t *testing.T) {
		expiresIn := time.Now().Add(time.Duration(1) * time.Minute).Unix()
		token, err := JwtClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id":    "uuid",
			"expires_in": expiresIn,
		},
			[]byte("secrets"))

		_, err = JwtValidate(0, token, []byte("secret"))

		assert.NotNil(t, err)
		assert.Equal(t, "signature is invalid", err.Error())
	})
}
