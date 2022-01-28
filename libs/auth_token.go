package libs

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/quanghuy219/catalog-backend-golang/models"
)

func JwtEncode(u *models.User) string {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	claims := &jwt.StandardClaims{
		Audience: string(rune(u.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString(signingKey)
	return signedToken
}
