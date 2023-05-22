package httpRespone

import (
	"github.com/gin-gonic/gin"
)

// Ping test connection
func Ping(c *gin.Context) {
	WriteOK(c, "success")
}
