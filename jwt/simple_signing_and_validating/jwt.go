package simple_signing_and_validating

import (
	"crypto/rand"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	mySecret []byte
)

func init() {
	mySecret = make([]byte, 32)
	rand.Read(mySecret)
}

type myCustomClaims struct {
	jwt.RegisteredClaims
	UserID int64 `json:"user_id,omitempty"`
}

var _ jwt.ClaimsValidator = (*myCustomClaims)(nil)

func (m *myCustomClaims) Validate() error {
	if m.UserID == 0 {
		return errors.New("must provide user ID")
	}
	return nil
}

func jwtNew() (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		myCustomClaims{
			UserID: 1,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "go-usage-examples",
				Subject:   "go-usage-examples-testing",
				Audience:  jwt.ClaimStrings{"go-usage-examples"},
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Second)),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ID:        "unique_jti",
			},
		},
	)
	return t.SignedString(mySecret)
}

func parse(tokenString string) (*myCustomClaims, error) {
	t, err := jwt.ParseWithClaims(tokenString, &myCustomClaims{},
		func(token *jwt.Token) (any, error) {
			return mySecret, nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithIssuedAt(),
		jwt.WithExpirationRequired(),
		jwt.WithAudience("go-usage-examples"),
		jwt.WithIssuer("go-usage-examples"),
	)
	if err != nil {
		return nil, err
	}
	return t.Claims.(*myCustomClaims), nil
}
