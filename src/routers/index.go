/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package routers

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// RouterInitialize
// Initialize all api route
func RouterInitialize() *gin.Engine {
	routerEngine := gin.Default()

	routerEngine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("HOST") + ":" + os.Getenv("PORT")},
		AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	User(routerEngine)
	Auth(routerEngine)

	return routerEngine
}
