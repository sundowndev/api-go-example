/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fbonhomm/api-go-example/source/models"
	"github.com/fbonhomm/api-go-example/source/services"
)

// UserCreate
func UserCreate(c *gin.Context) {
	user := models.User{
		Name:     c.PostForm("name"),
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}

	if err := services.Db.Create(&user).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"item": user})
	}
}

// UserGetID
func UserGetID(c *gin.Context) {
	user := models.User{}

	if err := services.Db.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"item": user})
	}
}

// UserDeleteID
func UserDeleteID(c *gin.Context) {
	user := models.User{}

	if err := services.Db.First(&user, c.Param("id")).Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"item": user})
	}
}
