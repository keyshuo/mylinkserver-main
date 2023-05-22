package community

import (
	"MyLink_Server/server/internal/app/handler/httpRespone"
	"MyLink_Server/server/internal/app/service/community"
	"github.com/gin-gonic/gin"
)

func GetComment(c *gin.Context) {
	page := c.Query("page")
	var pageInt int

	result, msg := community.GetComment(page, pageInt)
	if msg != "" {
		httpRespone.WriteFailed(c, msg)
		return
	}
	httpRespone.WriteOK(c, result)
}

func GetMyComment(c *gin.Context) {
	status := c.Value("status")
	if status == "false" {
		httpRespone.WriteFailed(c, "please login")
		return
	}
	account := c.Value("account").(string)
	// fmt.Println(account)
	page := c.Query("page")
	var pageInt int
	result, msg := community.GetMyComment(account, page, pageInt)
	if msg != "" {
		httpRespone.WriteFailed(c, msg)
		return
	}
	httpRespone.WriteOK(c, result)
}

func CreateComment(c *gin.Context) {
	status := c.Value("status")
	if status == "false" {
		httpRespone.WriteFailed(c, "please login")
		return
	}
	comment := c.Query("comment")
	time := c.Query("time")
	account := c.Value("account").(string)
	if msg := community.CreateComment(comment, time, account); msg != "" {
		httpRespone.WriteFailed(c, msg)
		return
	}
	httpRespone.WriteOK(c, nil)
}

func DeleteComment(c *gin.Context) {
	status := c.Value("status")
	if status == "false" {
		httpRespone.WriteFailed(c, "please login")
		return
	}
	time := c.Query("time")
	account := c.Value("account").(string)
	if msg := community.DeleteComment(account, time); msg != "" {
		httpRespone.WriteFailed(c, msg)
		return
	}
	httpRespone.WriteOK(c, nil)
}
