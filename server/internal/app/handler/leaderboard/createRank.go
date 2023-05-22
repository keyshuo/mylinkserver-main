package leaderboard

import (
	"MyLink_Server/server/internal/app/handler/httpRespone"
	"MyLink_Server/server/internal/app/service/leaderboard"
	"github.com/gin-gonic/gin"
)

func CreateRank(c *gin.Context) {
	status := c.Value("status")
	if status == "false" {
		httpRespone.WriteFailed(c, "please login")
		return
	}

	account := c.Value("account").(string)
	time := c.Query("time")
	date := c.Query("date")
	if msg := leaderboard.CreateRank(account, time, date); msg != "" {
		httpRespone.WriteFailed(c, msg)
		return
	}
	httpRespone.WriteOK(c, nil)
}
