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

// GenerateAccessToken
func GenerateAccessToken(user *models.User) (accessToken string, err error) {
	var ok bool
	var token *jwt.Token
	var claims jwt.MapClaims

	token = jwt.New(jwt.SigningMethodES256)

	if claims, ok = token.Claims.(jwt.MapClaims); !ok {
		return accessToken, errors.New("access token is not valid")
	}
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	accessToken, err = token.SignedString(services.PrivateKeyAccess)

	return accessToken, err
}
