package tests

import (
	"crypto/ed25519"
	"crypto/rand"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"myboardgames/minimalgame/internal/backend"
)

const (
	testKID      = "key"
	testIssuer   = "minimalgo"
	testAudience = "minimalgame"
)

func newKeyPair(t *testing.T) (ed25519.PublicKey, ed25519.PrivateKey) {
	t.Helper()
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	require.NoError(t, err)
	return pub, priv
}

// validClaims returns a fresh, valid set of claims. Each test case mutates it.
func validClaims() jwt.MapClaims {
	now := time.Now()
	return jwt.MapClaims{
		"iss":    testIssuer,
		"aud":    testAudience,
		"sub":    "minimalgo",
		"iat":    now.Unix(),
		"exp":    now.Add(time.Minute).Unix(),
		"jti":    "test-jti",
		"scopes": []string{"table:create", "table:read"},
		"obo":    "player-1",
	}
}

func sign(t *testing.T, method jwt.SigningMethod, kid string, key any, claims jwt.MapClaims) string {
	t.Helper()
	tok := jwt.NewWithClaims(method, claims)

	if kid != "" {
		tok.Header["kid"] = kid
	}
	
	signed, err := tok.SignedString(key)
	require.NoError(t, err)
	return signed
}

func TestVerifierParse(t *testing.T) {
	pub, priv := newKeyPair(t)
	_, otherPriv := newKeyPair(t) // attacker key, not registered in the verifier

	v := backend.NewVerifier(map[string]ed25519.PublicKey{testKID: pub})

	// mutate builds claims from validClaims and applies a modification.
	mutate := func(fn func(c jwt.MapClaims)) jwt.MapClaims {
		c := validClaims()
		fn(c)
		return c
	}

	type testCase struct {
		title      string
		token      func(t *testing.T) string
		wantErr    error  // matched with errors.Is (jwt sentinel errors)
		wantErrMsg string // matched with a substring (custom errors)
	}

	testCases := []testCase{
		{
			title: "valid jwt",
			token: func(t *testing.T) string {
				return sign(t, jwt.SigningMethodEdDSA, testKID, priv, validClaims())
			},
		},
		{
			title: "expired jwt",
			token: func(t *testing.T) string {
				c := mutate(func(c jwt.MapClaims) {
					c["iat"] = time.Now().Add(-2 * time.Hour).Unix()
					c["exp"] = time.Now().Add(-time.Hour).Unix()
				})
				return sign(t, jwt.SigningMethodEdDSA, testKID, priv, c)
			},
			wantErr: jwt.ErrTokenExpired,
		},
		{
			title: "missing exp",
			token: func(t *testing.T) string {
				c := mutate(func(c jwt.MapClaims) { delete(c, "exp") })
				return sign(t, jwt.SigningMethodEdDSA, testKID, priv, c)
			},
			wantErr: jwt.ErrTokenRequiredClaimMissing,
		},
		{
			title: "nil issued at",
			token: func(t *testing.T) string {
				c := mutate(func(c jwt.MapClaims) { delete(c, "iat") })
				return sign(t, jwt.SigningMethodEdDSA, testKID, priv, c)
			},
			wantErrMsg: "lifetime",
		},
		{
			title: "lifetime too long",
			token: func(t *testing.T) string {
				c := mutate(func(c jwt.MapClaims) {
					c["exp"] = time.Now().Add(time.Hour).Unix()
				})
				return sign(t, jwt.SigningMethodEdDSA, testKID, priv, c)
			},
			wantErrMsg: "lifetime",
		},
		{
			title: "wrong issuer",
			token: func(t *testing.T) string {
				c := mutate(func(c jwt.MapClaims) { c["iss"] = "wrong" })
				return sign(t, jwt.SigningMethodEdDSA, testKID, priv, c)
			},
			wantErr: jwt.ErrTokenInvalidIssuer,
		},
		{
			title: "wrong audience",
			token: func(t *testing.T) string {
				c := mutate(func(c jwt.MapClaims) { c["aud"] = "wrong" })
				return sign(t, jwt.SigningMethodEdDSA, testKID, priv, c)
			},
			wantErr: jwt.ErrTokenInvalidAudience,
		},
		{
			title: "unknown kid",
			token: func(t *testing.T) string {
				return sign(t, jwt.SigningMethodEdDSA, "unknown", priv, validClaims())
			},
			wantErr: jwt.ErrTokenUnverifiable,
		},
		{
			title: "missing kid",
			token: func(t *testing.T) string {
				return sign(t, jwt.SigningMethodEdDSA, "", priv, validClaims())
			},
			wantErr: jwt.ErrTokenUnverifiable,
		},
		{
			title: "signed with another key",
			token: func(t *testing.T) string {
				return sign(t, jwt.SigningMethodEdDSA, testKID, otherPriv, validClaims())
			},
			wantErr: jwt.ErrTokenSignatureInvalid,
		},
		{
			title: "alg none",
			token: func(t *testing.T) string {
				return sign(t, jwt.SigningMethodNone, testKID,
					jwt.UnsafeAllowNoneSignatureType, validClaims())
			},
			wantErr: jwt.ErrTokenSignatureInvalid,
		},
		{
			title: "hs256 key confusion (public key used as HMAC secret)",
			token: func(t *testing.T) string {
				return sign(t, jwt.SigningMethodHS256, testKID, []byte(pub), validClaims())
			},
			wantErr: jwt.ErrTokenSignatureInvalid,
		},
		{
			title: "hardcoded hs256 token with wrong claims",
			token: func(*testing.T) string {
				return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwiaXNzdWVyIjoid3JvbmciLCJhdWRpZW5jZSI6Indyb25nIn0.NcVGaIwY0dWrTtYBONJM7mF7Ba1iHAF_5FJZWKqK4oE"
			},
			wantErr: jwt.ErrTokenSignatureInvalid,
		},
		{
			title:   "garbage",
			token:   func(*testing.T) string { return "not-a-jwt" },
			wantErr: jwt.ErrTokenMalformed,
		},
		{
			title:   "empty",
			token:   func(*testing.T) string { return "" },
			wantErr: jwt.ErrTokenMalformed,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			claims, err := v.Parse(tc.token(t))

			if tc.wantErr == nil && tc.wantErrMsg == "" {
				require.NoError(t, err)
				require.NotNil(t, claims)
				assert.Equal(t, "minimalgo", claims.Subject)
				assert.Equal(t, "player-1", claims.OnBehalfOf)
				assert.ElementsMatch(t, []string{"table:create", "table:read"}, claims.Scopes)
				return
			}

			require.Error(t, err)
			assert.Nil(t, claims)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
			}
			if tc.wantErrMsg != "" {
				assert.ErrorContains(t, err, tc.wantErrMsg)
			}
		})
	}
}
