/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package validators

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type authLogin struct {
	Email    string `form:"email"       binding:"required,email"`
	Password string `form:"password"    binding:"required,min=8,max=50"`
}

type authRefresh struct {
	RefreshToken string `form:"refresh_token" binding:"required,min=100,max=1000"`
}

// ValidateAuthLogin
func ValidateAuthLogin(c *gin.Context) {
	var v authLogin

	if err := c.ShouldBindWith(&v, binding.FormPost); err != nil {
		errorHandling(c, err.Error())
		return
	}
}

// ValidateAuthRefresh
func ValidateAuthRefresh(c *gin.Context) {
	var v authRefresh

	if err := c.ShouldBindWith(&v, binding.FormPost); err != nil {
		errorHandling(c, err.Error())
		return
	}
}
