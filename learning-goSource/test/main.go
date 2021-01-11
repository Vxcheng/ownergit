package main

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
)

func main() {
	var username = "root"
	var password = "root"
	var host = "localhost"
	var serviceName = "test"
	var port = 3309

	dsn := fmt.Sprintf("%s/%s@%s:%d/%s", username, password, host, port, serviceName)
	db, err := sql.Open("godror", dsn)
	if err != nil {
		fmt.Printf("sql.Open failed %v \r\n", err)
		return
	}
	if err = db.Ping(); err != nil {
		fmt.Printf("db.Ping failed %v \r\n", err)
		return
	}
	defer db.Close()

	querySQL := `select version from v$instance`
	var version string
	err = db.QueryRow(querySQL).Scan(&version)
	if err != nil {
		fmt.Printf("db.QueryRow failed %v \r\n", err)
		return
	}
	fmt.Printf("db version :%s \r\n", version)
}
