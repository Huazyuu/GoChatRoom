package server

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var MySQLDB *sql.DB
var dbErr error

const (
	UserName = "root"
	Password = "your own pwd"
	Host     = "127.0.0.1"
	Port     = "3306"
	Database = "go_chat"
	Charset  = "utf8"
)

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		UserName, Password, Host, Port, Database)
	MySQLDB, dbErr = sql.Open("mysql", dsn)
	if dbErr != nil {
		fmt.Println("Mysql open err: ", dbErr)
	}
	MySQLDB.SetMaxOpenConns(50)
	MySQLDB.SetMaxIdleConns(10)
	MySQLDB.SetConnMaxLifetime(30 * time.Second)
	if dbErr = MySQLDB.Ping(); dbErr != nil {
		panic("Mysql open err: " + dbErr.Error())
	}

}
func LogToDb(msg string, address string) int64 {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("LogToDb: ", err)
		}
	}()

	ret, err := MySQLDB.Exec("insert into chat_logs (message,address) values (?,?)",
		msg, address)
	if err != nil {
		fmt.Println(" MySQLDB.Exec error: ", err)
		return -1
	}
	rows, _ := ret.RowsAffected()
	return rows
}
