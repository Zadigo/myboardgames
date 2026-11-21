package backend

import (
	"crypto/ed25519"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	issuer      = "minimalgo"
	audience    = "minimalgame"
	maxTokenTTL = 5 * time.Minute
)

type Claims struct {
	jwt.RegisteredClaims
	Scopes []string `json:"scopes"`
	// The player the action is performed for (asserted by minimalgo)
	OnBehalfOf string `json:"obo,omitempty"`
}

type Verifier struct {
	keys map[string]ed25519.PublicKey
}

func NewVerifier(keys map[string]ed25519.PublicKey) *Verifier {
	return &Verifier{keys: keys}
}

func (v *Verifier) Parse(value string) (*Claims, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(value, claims,
		func(t *jwt.Token) (any, error) {
			return claims, nil
		},
		jwt.WithValidMethods([]string{"EdDSA"}),
		jwt.WithIssuer(issuer),
		jwt.WithAudience(audience),
		jwt.WithExpirationRequired(),
		jwt.WithIssuedAt(),
		jwt.WithLeeway(5*time.Second),
	)

	if err != nil {
		return nil, err
	}

	if claims.IssuedAt == nil || claims.ExpiresAt.Sub(claims.IssuedAt.Time) > maxTokenTTL {
		return nil, errors.New("Token lifetime too long")
	}

	return claims, err
}

type Caller struct {
	Service    string
	OnBehalfOf string
	Scopes     map[string]struct{}
}

func (c *Caller) Has(scope string) bool {
	_, ok := c.Scopes[scope]
	return ok
}
