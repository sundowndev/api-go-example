/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package libs

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetToken
func GetToken(c *gin.Context) (string, error) {
	var reqToken = c.Request.Header.Get("Authorization")
	var splitToken = strings.Split(reqToken, "Bearer")

	if len(splitToken) != 2 {
		return "", errors.New("token not provided")
	}

	return strings.TrimSpace(splitToken[1]), nil
}
