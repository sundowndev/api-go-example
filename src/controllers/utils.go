/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// GetToken
func GetToken(c *gin.Context) (map[string]string, error) {
	token, ok := c.Get("Token")

	if !ok {
		return nil, errors.New("token is not found")
	}

	return token.(map[string]string), nil
}
