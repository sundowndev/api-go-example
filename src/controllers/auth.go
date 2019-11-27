/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fbonhomm/api-go-example/source/libs"
	"github.com/fbonhomm/api-go-example/source/models"
	"github.com/fbonhomm/api-go-example/source/services"
)

// AuthLogin
func AuthLogin(c *gin.Context) {
	var accessToken string
	var refreshToken string
	var err error
	var user = models.User{
		Email: c.PostForm("email"),
	}

	if err = services.Db.First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else if err = user.Compare(c.PostForm("password")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else if accessToken, err = libs.GenerateAccessToken(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else if refreshToken, err = libs.GenerateRefreshToken(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	}
}

// AuthRefresh
func AuthRefresh(c *gin.Context) {
	var accessToken string
	var refreshToken string
	var token map[string]string
	var err error
	var user = models.User{}

	if token, err = GetToken(c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if err = services.Db.First(&user, token["id"]).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else if accessToken, err = libs.GenerateAccessToken(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else if refreshToken, err = libs.GenerateRefreshToken(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	}
}
