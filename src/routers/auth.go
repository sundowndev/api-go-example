/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/fbonhomm/api-go-example/source/controllers"
	"github.com/fbonhomm/api-go-example/source/middlewares"
	"github.com/fbonhomm/api-go-example/source/validators"
)

// main
func Auth(router *gin.Engine) {
	route := router.Group("/auth")

	route.POST("", validators.ValidateAuthLogin, controllers.AuthLogin)
	route.POST("/refresh", middlewares.AuthRefresh, validators.ValidateAuthRefresh, controllers.AuthRefresh)
}
