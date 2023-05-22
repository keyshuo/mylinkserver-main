package leaderboard

import (
	"MyLink_Server/server/internal/app/handler/httpRespone"
	"MyLink_Server/server/internal/app/service/leaderboard"
	"github.com/gin-gonic/gin"
)

// GetRankLow low difficulty
func GetRankLow(c *gin.Context) {
	result, msg := leaderboard.GetRankLow()
	if msg != "" {
		httpRespone.WriteFailed(c, msg)
		return
	}
	httpRespone.WriteOK(c, result)
}

// GetRankMedium medium difficulty
func GetRankMedium(c *gin.Context) {
	result, msg := leaderboard.GetRankMedium()
	if msg != "" {
		httpRespone.WriteFailed(c, msg)
		return
	}
	httpRespone.WriteOK(c, result)
}

// GetRankHigh high difficulty
func GetRankHigh(c *gin.Context) {
	result, msg := leaderboard.GetRankHigh()
	if msg != "" {
		httpRespone.WriteFailed(c, msg)
		return
	}
	httpRespone.WriteOK(c, result)
}
