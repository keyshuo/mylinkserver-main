package user

import (
	"MyLink_Server/server/internal/app/service/log"
	app "MyLink_Server/server/internal/app/service/sqloperate"
)

func Register(inputUser User) string {

	msg := "select count(*) from user where account = ? ;"
	db, err := app.NewMySql(msg)
	if log.ErrorLog(err) != nil {
		return log.DatabaseConnFail
	}
	defer db.Close()

	//这里可能存在BUG，count语句使用string数组保存
	result, err := db.Search(inputUser.Account)

	if err == nil {
		if result[0] >= "1" {
			return log.UserExist
		}
	} else {
		log.ErrorLog(err)
		return log.DatabaseSearchFail
	}

	msg = "select count(*) from user where username = ? ;"
	err = db.UpdateMysql(msg)
	if log.ErrorLog(err) != nil {
		return log.DatabaseUpdateConn
	}

	result1, err := db.Search(inputUser.Username)

	if err == nil {
		if result1[0] >= "1" {
			return log.NameRepeat
		} else {
			msg = "insert into user value ( ?, ?, ?);"
			err = db.UpdateMysql(msg)
			if log.ErrorLog(err) != nil {
				return log.DatabaseUpdateConn
			}
			err = db.Exec(inputUser.Account, inputUser.Username, inputUser.Password)
			if err != nil {
				return log.DatabaseExecFail
			}
			return ""
		}
	} else {
		return log.DatabaseSearchFail
	}
}
