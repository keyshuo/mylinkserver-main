package leaderboard

import (
	"MyLink_Server/server/internal/app/service/log"
	sqloperate "MyLink_Server/server/internal/app/service/sqloperate"
)

func CreateRank(account, time, date string) string {
	msg := "insert into ranktable value ( ?, ?, ?);"
	db, err := sqloperate.NewMySql(msg)
	if log.ErrorLog(err) != nil {
		return log.DatabaseConnFail
	}
	defer db.Close()
	err = db.Exec(account, time, date)
	if log.ErrorLog(err) != nil {
		return log.DatabaseExecFail
	}
	return ""
}
