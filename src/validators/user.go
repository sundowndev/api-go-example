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

// https://godoc.org/gopkg.in/go-playground/validator.v8

type userCreate struct {
	Name     string `form:"name" binding:"required,min=2,max=50"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=8,max=50"`
}

type userGetID struct {
	ID uint `uri:"id" binding:"required,min=0,max=9999"`
}

type userDeleteID struct {
	ID uint `uri:"id" binding:"required,min=0,max=9999"`
}

// ValidateUserCreate
func ValidateUserCreate(c *gin.Context) {
	var v userCreate

	if err := c.ShouldBindWith(&v, binding.FormPost); err != nil {
		errorHandling(c, err.Error())
	}
}

// ValidateUserGetID
func ValidateUserGetID(c *gin.Context) {
	var v userGetID

	if err := c.ShouldBindUri(&v); err != nil {
		errorHandling(c, err.Error())
	}
}

// ValidateUserDeleteID
func ValidateUserDeleteID(c *gin.Context) {
	var v userDeleteID

	if err := c.ShouldBindUri(&v); err != nil {
		errorHandling(c, err.Error())
	}
}
