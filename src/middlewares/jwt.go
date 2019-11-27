/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/fbonhomm/api-go-example/source/libs"
	"github.com/fbonhomm/api-go-example/source/services"
)

func Auth(c *gin.Context) {
	tokenString, err := libs.GetToken(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return services.PublicKeyAccess, nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not conform."})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Set("Token", claims)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not conform."})
		c.Abort()
	}
}

func AuthRefresh(c *gin.Context) {
	tokenString := c.PostForm("refresh_token")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return services.PublicKeyRefresh, nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not conform."})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var decode = make(map[string]string, 10)

		for key, value := range claims {
			decode[key] = fmt.Sprintf("%v", value)
		}
		c.Set("Token", decode)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not conform."})
		c.Abort()
	}
}
