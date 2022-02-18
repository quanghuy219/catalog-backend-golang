package libs

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/quanghuy219/catalog-backend-golang/models"
)

func JwtEncode(u *models.User) string {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	claims := &jwt.StandardClaims{
		Audience:  strconv.Itoa(int(u.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString(signingKey)
	return signedToken
}

func ParseJwtToken(tokenString string) (*jwt.Token, error) {
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("invalid token format")
			} else if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
				return nil, errors.New("token has expired")
			} else {
				return nil, errors.New("couldn't handle token")
			}
		}
	} else if token.Valid {
		return token, nil
	}
	return nil, errors.New("couldn't handle token")
}
