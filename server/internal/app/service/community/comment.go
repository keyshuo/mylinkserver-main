package community

import (
	"MyLink_Server/server/internal/app/service/log"
	sqloperate "MyLink_Server/server/internal/app/service/sqloperate"
	"fmt"
)

func GetComment(page string, pageInt int) ([]interface{}, string) {
	if _, err := fmt.Sscan(page, &pageInt); log.ErrorLog(err) != nil {
		return nil, log.ServerError
	}
	page = string(rune((pageInt - 1) * 50))

	msg := "select user.username,comment.date,comment.comment from comment join user on user.account=comment.account order by comment.date desc limit 50 offset ?; "
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

func GetMyComment(account, page string, pageInt int) ([]interface{}, string) {
	if _, err := fmt.Sscan(page, &pageInt); log.ErrorLog(err) != nil {
		return nil, log.ServerError
	}
	page = string(rune((pageInt - 1) * 50))

	msg := "select user.username,comment.date,comment.comment from comment join user on user.account=comment.account where comment.account = ? order by comment.date desc limit 50 offset ? ;"
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

func CreateComment(comment, time, account string) string {

	msg := "insert into comment value (?,?,?);"
	db, err := sqloperate.NewMySql(msg)
	if log.ErrorLog(err) != nil {
		return log.DatabaseConnFail
	}
	defer db.Close()
	err = db.Exec(account, time, comment)
	if log.ErrorLog(err) != nil {
		return log.DatabaseExecFail
	}
	return ""
}

func DeleteComment(account, date string) string {
	msg := "delete from comment where account = ? and date = ?;"
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
