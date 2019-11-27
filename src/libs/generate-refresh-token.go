/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package libs

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/fbonhomm/api-go-example/source/models"
	"github.com/fbonhomm/api-go-example/source/services"
)

// GenerateRefreshToken
func GenerateRefreshToken(user *models.User) (refreshToken string, err error) {
	var ok bool
	var token *jwt.Token
	var claims jwt.MapClaims

	token = jwt.New(jwt.SigningMethodES384)

	if claims, ok = token.Claims.(jwt.MapClaims); !ok {
		return refreshToken, errors.New("refresh token is not valid")
	}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	refreshToken, err = token.SignedString(services.PrivateKeyRefresh)

	return refreshToken, err
}
