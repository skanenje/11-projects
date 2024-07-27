package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
func main(){
	const (
		username = "sql7722456"
		password = "zedolph@13"
		host = "sql7.freemysqlhosting.net"
		port = 8080
		dbname = "mydb"
	)
	connstr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbname)	

	db, err := sql.Open("mysql", connstr)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	_, err := db.Exec("")
}