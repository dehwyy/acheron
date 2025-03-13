package extractors

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	headerKeyAuthorization = "Authorization"
)

var (
	ErrNoAuthorizationToken     = errors.New("Authorization token is not provided in headers")
	ErrInvalidAuthorizatioToken = errors.New("Invalid authorization token - valid form is 'Bearer <token>'")
)

func GetAuthorizationToken(ctx *gin.Context) (string, error) {
	token := ctx.GetHeader(headerKeyAuthorization)

	if token == "" {
		return "", ErrNoAuthorizationToken
	}

	tokens := strings.Split(token, " ")
	if len(tokens) != 2 || tokens[0] != "Bearer" {
		return "", ErrInvalidAuthorizatioToken
	}

	return tokens[1], nil
}
