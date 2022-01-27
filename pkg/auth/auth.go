package auth

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"io/ioutil"
	"time"
)


// Claims represents the authorization claims transmitted via a JWT.
type Claims struct {
	jwt.StandardClaims
}

func GetRSAPublicKey() (*rsa.PublicKey, error) {
	keyData, err := ioutil.ReadFile("keys/id_rsa.pub.pkcs8")

	if err != nil {
		return nil, err
	}

	return jwt.ParseRSAPublicKeyFromPEM(keyData)
}

// GenerateToken generates a signed JWT token string representing the user Claims.
func GenerateToken(claims Claims) (string, error) {
	keyData, err := ioutil.ReadFile("keys/id_rsa")

	if err != nil {
		return "", err
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)

	if err != nil {
		return "", err
	}

	method := jwt.GetSigningMethod("RS256")
	if method == nil {
		return "", errors.New("configuring algorithm RS256")
	}

	tkn := jwt.NewWithClaims(method, claims)

	str, err := tkn.SignedString(key)
	if err != nil {
		return "", errors.Wrap(err, "signing token")
	}

	return str, nil
}


// NewClaims constructs a Claims value for the identified user. The Claims
// expire within a specified duration of the provided time. Additional fields
// of the Claims can be set after calling NewClaims is desired.
func NewClaims(subject string,  now time.Time, expires time.Duration) Claims {
	c := Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   subject,
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(expires).Unix(),
		},
	}

	return c
}

