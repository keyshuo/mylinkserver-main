package app

import (
	"database/sql"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"k8s.io/klog"
)

type MySql struct {
	mysql *sql.DB
	stmt  *sql.Stmt
}

// NewMySql 新建连接，从连接池拿一个
func NewMySql(msg string) (*MySql, error) {
	mysqlTemp, err := sql.Open("mysql", "root:@Wx614481987@tcp(1.15.76.132:3306)/androidDatabase")
	if err != nil {
		return nil, err
	}
	err = mysqlTemp.Ping()
	if err != nil {
		return nil, err
	}
	stmtTemp, err := mysqlTemp.Prepare(msg)
	if err != nil {
		return nil, err
	}
	return &MySql{
		mysql: mysqlTemp,
		stmt:  stmtTemp,
	}, nil
}

// Search 用sql语句查询，返回一个结果字符串列表
func (sql *MySql) Search(args ...interface{}) ([]string, error) {
	var result []string
	var temp string
	rows, err := sql.stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&temp)
		if err != nil {
			return nil, err
		}
		result = append(result, temp)
	}
	return result, nil
}

func (sql *MySql) SearchRows(obj interface{}, args ...interface{}) ([]interface{}, error) {
	rows, err := sql.stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	s := reflect.ValueOf(obj).Elem()
	t := s.Type()
	fields := make(map[string]int)
	for i := 0; i < s.NumField(); i++ {
		fields[t.Field(i).Name] = i
	}

	// 遍历查询结果
	results := make([]interface{}, 0)
	for rows.Next() {
		// 创建结构体实例
		r := reflect.New(t).Elem()

		// 将查询结果映射到结构体
		values := make([]interface{}, s.NumField())
		for i := 0; i < s.NumField(); i++ {
			field := t.Field(i)
			values[i] = r.FieldByName(field.Name).Addr().Interface()
		}
		err := rows.Scan(values...)
		if err != nil {
			return nil, err
		}
		for i := 0; i < s.NumField(); i++ {
			field := t.Field(i)
			if field.Name == "Time" {
				if values[i] != nil {
					value := values[i].(*string)
					date, err := time.ParseInLocation("2006-01-02 15:04:05", *value, time.Local)
					if err != nil {
						return nil, err
					}
					r.FieldByName(field.Name).Set(reflect.ValueOf(date.Format("2006-01-02 15:04:05")))
				}
			}
		}
		results = append(results, r.Addr().Interface())
	}

	return results, nil
}

// Exec 执行UPDATE、INSERT、DELETE操作
func (sql *MySql) Exec(args ...interface{}) error {
	_, err := sql.stmt.Exec(args...)
	if err != nil {
		return err
	}
	return err
}

// Close 关闭数据库连接
func (sql *MySql) Close() {
	err := sql.mysql.Close()
	if err != nil {
		klog.Error(err)
		return
	}
	err = sql.stmt.Close()
	if err != nil {
		klog.Error(err)
		return
	}
}

func (newSql *MySql) UpdateMysql(msg string) error {
	mysqlTemp, err := sql.Open("mysql", "root:@Wx614481987@tcp(1.15.76.132:3306)/androidDatabase")
	if err != nil {
		return err
	}
	err = mysqlTemp.Ping()
	if err != nil {
		return err
	}
	stmtTemp, err := mysqlTemp.Prepare(msg)
	if err != nil {
		return err
	}
	newSql.mysql = mysqlTemp
	newSql.stmt = stmtTemp
	return nil
}
