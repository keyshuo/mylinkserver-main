package music

import (
	"MyLink_Server/server/internal/app/service/log"
	sql "MyLink_Server/server/internal/app/service/sqloperate"
	"fmt"
	"io/ioutil"
	"os"
)

func GetMusicList() ([]string, string) {
	msg := "select name from music"
	db, err := sql.NewMySql(msg)
	defer db.Close()
	if log.ErrorLog(err) != nil {
		return nil, log.DatabaseConnFail
	}
	result, err := db.Search()
	if log.ErrorLog(err) != nil {
		return nil, log.DatabaseSearchFail
	}
	return result, ""
}

func GetMusic(name string) ([]byte, string) {
	msg := "select path from music where name = ? ;"
	fmt.Println(name)
	db, err := sql.NewMySql(msg)
	defer db.Close()
	if log.ErrorLog(err) != nil {
		return nil, log.DatabaseConnFail
	}
	result, err := db.Search(name)
	if log.ErrorLog(err) != nil {
		return nil, log.DatabaseSearchFail
	}
	path := result[0]
	file, err := os.Open(path)
	if log.ErrorLog(err) != nil {
		return nil, log.FileOPenFail
	}
	defer file.Close()
	musicData, err := ioutil.ReadAll(file)
	if log.ErrorLog(err) != nil {
		return nil, log.FileReadFail
	}
	return musicData, ""
}
