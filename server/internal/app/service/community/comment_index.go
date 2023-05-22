package community

import (
	"MyLink_Server/server/internal/app/service/log"
	sqloperate "MyLink_Server/server/internal/app/service/sqloperate"
	"fmt"
)

func GetCommentIndex(page string, pageInt int, username, date string) ([]interface{}, string) {
	if _, err := fmt.Sscan(page, &pageInt); log.ErrorLog(err) != nil {
		return nil, log.ServerError
	}
	page = string(rune((pageInt - 1) * 50))

	//SQL逻辑：
	//根据发送过来的username在用户表中找到被评论人的account，
	//根据这个account和发送过来的date确定这条动态底下所有的评论，即account、date和comment
	//最后根据查找到的account在user表中找到username
	//最终返回username
	msg := `
	SELECT  u.username,ci.comment, ci.date
	FROM user u
	JOIN comment_index ci ON u.account = ci.account
	WHERE ci.account_index = (SELECT account FROM user WHERE username = ?)
	AND ci.date_index = ?
	ORDER BY ci.date DESC
	LIMIT 50 offset ?;
	`
	db, err := sqloperate.NewMySql(msg)
	if log.ErrorLog(err) != nil {
		return nil, log.DatabaseConnFail
	}

	defer db.Close()

	result, err := db.SearchRows(&UserComment{}, page)
	if log.ErrorLog(err) != nil {
		return nil, log.DatabaseSearchFail
	}

	return result, ""
}

func GetMyCommentIndex(account, page string, pageInt int) ([]interface{}, string) {
	if _, err := fmt.Sscan(page, &pageInt); log.ErrorLog(err) != nil {
		return nil, log.ServerError
	}
	page = string(rune((pageInt - 1) * 50))

	msg := `select user.username,comment_index.date,comment_index.comment
	from comment_index
	join user on user.account=comment_index.account 
	where comment_index.account = ?
	order by comment_index.date desc 
	limit 50 offset ? ;
	`
	db, err := sqloperate.NewMySql(msg)
	if log.ErrorLog(err) != nil {
		return nil, log.DatabaseConnFail
	}
	defer db.Close()

	result, err := db.SearchRows(&UserComment{}, account, page)
	if log.ErrorLog(err) != nil {
		return nil, log.DatabaseSearchFail
	}

	return result, ""
}

func CreateCommentIndex(comment, time, account, accountIndex, timeIndex string) string {

	//
	msg := "insert into comment_index value (?,?,?,?,?);"
	db, err := sqloperate.NewMySql(msg)
	if log.ErrorLog(err) != nil {
		return log.DatabaseConnFail
	}
	defer db.Close()
	err = db.Exec(account, accountIndex, timeIndex, time, comment)
	if log.ErrorLog(err) != nil {
		return log.DatabaseExecFail
	}
	return ""
}

func DeleteCommentIndex(date, account string) string {
	msg := "delete from comment where account = ? and date = ? ;"
	db, err := sqloperate.NewMySql(msg)
	if log.ErrorLog(err) != nil {
		return log.DatabaseConnFail
	}
	defer db.Close()
	err = db.Exec(account, date)
	if log.ErrorLog(err) != nil {
		return log.DatabaseExecFail
	}
	return ""
}
