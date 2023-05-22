package community

import (
	"MyLink_Server/server/internal/app/handler/httpRespone"
	"MyLink_Server/server/internal/app/service/community"
	"github.com/gin-gonic/gin"
)

func GetCommentIndex(c *gin.Context) {
	page := c.Query("page")
	username := c.Query("username") //发布动态的用户，被评论的人
	time := c.Query("time")         //动态发布的时间
	var pageInt int

	result, msg := community.GetCommentIndex(page, pageInt, username, time)
	if msg != "" {
		httpRespone.WriteFailed(c, msg)
		return
	}
	httpRespone.WriteOK(c, result)
}

func GetMyCommentIndex(c *gin.Context) {
	status := c.Value("status")
	if status == "false" {
		httpRespone.WriteFailed(c, "please login")
		return
	}
	account := c.Value("account").(string)
	// fmt.Println(account)
	page := c.Query("page")
	var pageInt int
	result, msg := community.GetMyCommentIndex(account, page, pageInt)
	if msg != "" {
		httpRespone.WriteFailed(c, msg)
		return
	}
	httpRespone.WriteOK(c, result)
}

func CreateCommentIndex(c *gin.Context) {
	status := c.Value("status")
	if status == "false" {
		httpRespone.WriteFailed(c, "please login")
		return
	}
	comment := c.Query("comment")
	date := c.Query("time")
	account := c.Value("account").(string)
	dateIndex := c.Query("date_index")
	accountIndex := c.Query("account_index")
	if msg := community.CreateCommentIndex(comment, date, account, accountIndex, dateIndex); msg != "" {
		httpRespone.WriteFailed(c, msg)
		return
	}
	httpRespone.WriteOK(c, nil)
}

func DeleteCommentIndex(c *gin.Context) {
	status := c.Value("status")
	if status == "false" {
		httpRespone.WriteFailed(c, "please login")
		return
	}
	date := c.Query("time")
	account := c.Value("account").(string)
	if msg := community.DeleteCommentIndex(date, account); msg != "" {
		httpRespone.WriteFailed(c, msg)
		return
	}
	httpRespone.WriteOK(c, nil)
}
