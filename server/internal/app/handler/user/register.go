package user

import (
	"MyLink_Server/server/internal/app/handler/httpRespone"
	"MyLink_Server/server/internal/app/service/log"
	"MyLink_Server/server/internal/app/service/user"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var inputUser user.User
	err := c.ShouldBindJSON(&inputUser)
	if log.ErrorLog(err) != nil {
		httpRespone.WriteFailed(c, "data acquisition failed ")
		return
	}
	if msg := user.Register(inputUser); msg != "" {
		httpRespone.WriteFailed(c, msg)
		return
	}
	httpRespone.WriteOK(c, nil)
}
