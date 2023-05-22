package leaderboard

import (
	"MyLink_Server/server/internal/app/service/log"
	sqloperate "MyLink_Server/server/internal/app/service/sqloperate"
)

// GetRankLow low difficulty
func GetRankLow() ([]interface{}, string) {
	msg := "select user.username,ranklow.date,ranklow.score from ranklow join user on user.account=ranklow.account order by score desc limit 50;"
	db, err := sqloperate.NewMySql(msg)
	if log.ErrorLog(err) != nil {
		return nil, log.DatabaseConnFail
	}

	defer db.Close()

	result, err := db.SearchRows(&UserRank{})
	if log.ErrorLog(err) != nil {
		return nil, log.DatabaseSearchFail
	}

	return result, ""
}

// GetRankMedium medium difficulty
func GetRankMedium() ([]interface{}, string) {
	msg := "select user.username,rankmedium.date,rankmedium.score from rankmedium join user on user.account=rankmedium.account order by score desc limit 50;"
	db, err := sqloperate.NewMySql(msg)
	if log.ErrorLog(err) != nil {
		return nil, log.DatabaseConnFail
	}

	defer db.Close()

	result, err := db.SearchRows(&UserRank{})
	if log.ErrorLog(err) != nil {
		return nil, log.DatabaseSearchFail
	}

	return result, ""
}

// GetRankHigh high difficulty
func GetRankHigh() ([]interface{}, string) {
	msg := "select user.username,rankhigh.date,rankhigh.score from rankhigh join user on user.account=rankhigh.account order by score desc limit 50;"
	db, err := sqloperate.NewMySql(msg)
	if log.ErrorLog(err) != nil {
		return nil, log.DatabaseConnFail
	}

	defer db.Close()

	result, err := db.SearchRows(&UserRank{})
	if log.ErrorLog(err) != nil {
		return nil, log.DatabaseSearchFail
	}

	return result, ""
}
