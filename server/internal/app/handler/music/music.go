package music

import (
	"MyLink_Server/server/internal/app/handler/httpRespone"
	"MyLink_Server/server/internal/app/service/music"
	"github.com/gin-gonic/gin"
)

func GetMusicList(c *gin.Context) {
	result, msg := music.GetMusicList()
	if msg != "" {
		httpRespone.WriteFailed(c, msg)
		return
	}
	httpRespone.WriteOK(c, result)
}

func GetMusic(c *gin.Context) {
	name := c.Query("name")
	musicData, msg := music.GetMusic(name)
	if msg != "" {
		httpRespone.WriteFailed(c, msg)
		return
	}
	httpRespone.WriteMusic(c, musicData)
}
