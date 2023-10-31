package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginSession struct {
	ExpiresAt time.Time
	Uuid      string
	UserID    uint
}

type jwtCustomClaims struct {
	jwt.RegisteredClaims
	Admin bool
}

// ErrNoAuthHeaderIncluded -
var ErrNoAuthHeaderIncluded = errors.New("no auth header included in request")

// HashPassword -
func HashPassword(password string) (string, error) {
	dat, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

// CheckPasswordHash -
func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// MakeJWT -
func MakeJWT(userID string, tokenSecret string, expiresIn time.Duration, issuer string, admin bool) (string, error) {
	signingKey := []byte(tokenSecret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtCustomClaims{
		jwt.RegisteredClaims{
			Issuer:    issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
			Subject:   userID,
		},
		admin,
	})
	return token.SignedString(signingKey)
}

type claims struct {
	Subject string
	Issuer  string
	Admin   bool
}

// ValidateJWT -
func ValidateJWT(tokenString, tokenSecret string) (claims, error) {
	claimsStruct := struct {
		jwt.RegisteredClaims
		Admin bool
	}{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(tokenSecret), nil
		},
	)
	if err != nil {
		return claims{}, err
	}

	userIDString, err := token.Claims.GetSubject()
	if err != nil {
		return claims{}, err
	}

	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return claims{}, err
	}

	return claims{
		Subject: userIDString,
		Issuer:  issuer,
		Admin:   claimsStruct.Admin,
	}, nil
}
