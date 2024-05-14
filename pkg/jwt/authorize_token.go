package jwt

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/radityacandra/go-project-mongodb/pkg/jwt/types"
)

const (
	ISSUER      = "Kecilin-Challenge"
	AUDIENCE    = "Kecilin-Challenge"
	CONTEXT_KEY = "token_data"
	TOKEN_DATA  = tokenData("token_data")
)

func AuthorizeToken(authorizationStr string) (map[string]string, error) {
	authPart := strings.Split(authorizationStr, " ")
	if len(authPart) != 2 && authPart[0] != "Bearer" {
		return nil, types.NewAuthorizationError("invalid authorization header")
	}
	tokenStr := authPart[1]

	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(os.Getenv("JWT_PUBLIC_KEY")))
	if err != nil {
		return nil, types.NewAuthorizationError(err.Error())
	}

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, types.NewAuthorizationError(fmt.Sprintf("unexpected signing method: %v", t.Header["alg"]))
		}

		return key, nil
	})
	if err != nil {
		return nil, types.NewAuthorizationError(err.Error())
	}

	if !token.Valid {
		return nil, types.NewAuthorizationError("failed to verify signature")
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok {
		return nil, types.NewAuthorizationError("invalid token")
	}

	claim := token.Claims.(jwt.MapClaims)
	// validate nbf claim
	if nbf, err := claim.GetNotBefore(); err != nil || nbf.Unix() > time.Now().Unix() {
		return nil, types.NewAuthorizationError("invalid token")
	}

	// validate expired claim
	if exp, err := claim.GetExpirationTime(); err != nil || exp.Unix() < time.Now().Unix() {
		return nil, types.NewAuthorizationError("token has been expired")
	}

	// validate issuer
	if issuer, err := claim.GetIssuer(); err != nil || issuer != ISSUER {
		return nil, types.NewAuthorizationError("invalid token")
	}

	// validate audience
	if audience, err := claim.GetAudience(); err != nil || audience[0] != AUDIENCE {
		return nil, types.NewAuthorizationError("invalid token")
	}

	mapData := make(map[string]string)
	for key, value := range claim {
		mapData[key] = fmt.Sprint(value)
	}

	return mapData, nil
}
