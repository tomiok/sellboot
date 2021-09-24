package jwts

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
	"time"
)

func CreateToken(long bool, name string, role int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	var exp = time.Hour * 24 * 3
	if long {
		exp = time.Hour * 24 * 30
	}
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = name
	claims["role"] = role
	claims["exp"] = time.Now().Add(exp).Unix()

	t, err := token.SignedString([]byte("configs.Get().JWTSecret"))
	if err != nil {
		log.Error().Msgf("cannot sign JWT %v", err)
		return "", err
	}

	return t, nil
}
