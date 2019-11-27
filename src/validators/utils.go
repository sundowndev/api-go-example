/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package validators

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func errorHandling(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": msg})
	c.Abort()
}
