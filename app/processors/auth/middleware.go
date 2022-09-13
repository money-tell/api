package auth

import (
	"context"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/lestrrat-go/jwx/jwt"
)

func NewAuthenticator(ja *JWTAuth) openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		return Authenticator(ja, input)
	}
}

func Authenticator(ja *JWTAuth, input *openapi3filter.AuthenticationInput) error {
	if input.SecuritySchemeName != "BearerAuth" {
		return fmt.Errorf("security scheme %s != 'BearerAuth'", input.SecuritySchemeName)
	}

	token, err := VerifyRequest(ja, input.RequestValidationInput.Request, TokenFromHeader, TokenFromCookie)
	if err != nil || token == nil {
		return fmt.Errorf("getting token: %w", err)
	}

	if err := jwt.Validate(token); err != nil {
		return fmt.Errorf("getting token: %w", err)
	}

	_, err = UserIDFromToken(token)
	if err != nil {
		return fmt.Errorf("user is not found: %w", err)
	}

	return nil
}
