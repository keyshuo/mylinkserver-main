package httpRespone

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// WriteOK connect and visit successfully
func WriteOK(c *gin.Context, msg interface{}) {
	if msg != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{})
	}
}

// WriteFailed connect and visit failed
func WriteFailed(c *gin.Context, err interface{}) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":  401,
		"error": err,
	})
}

func WriteMusic(c *gin.Context, music interface{}) {
	if music != nil {
		c.Data(http.StatusOK, "audio/mpeg", music.([]byte))
	} else {
		c.JSON(http.StatusOK, gin.H{})
	}
}
