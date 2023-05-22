package user

import (
	"MyLink_Server/server/internal/app/service/log"
	sql "MyLink_Server/server/internal/app/service/sqloperate"
)

func Login(inputUser User) (string, string) {
	msg := "select count(*) from user where account= ? and password=?;"
	db, err := sql.NewMySql(msg)
	if log.ErrorLog(err) != nil {
		return "", log.DatabaseConnFail
	}
	defer db.Close()
	result, err := db.Search(inputUser.Account, inputUser.Password)
	if log.ErrorLog(err) != nil {
		return "", log.DatabaseSearchFail
	}
	if result[0] == "1" {
		tokenString, err := GenerateToken(inputUser)
		if log.ErrorLog(err) != nil {
			return "", log.TokenGenerate
		}
		return tokenString, ""
	}
	return "", log.UserNotExist
}
