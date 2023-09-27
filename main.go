package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	DB_PATH := os.Getenv("DB_PATH")
	if DB_PATH == "" {
		DB_PATH = "127.0.0.1:3306"
	}

	dbconf := "root:root@tcp(" + DB_PATH + ")/sample"

	db, err := sql.Open("mysql", dbconf)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	fmt.Println(db)

	err = db.Ping()

	if err != nil {
		fmt.Println("Ping 失敗")
		return
	} else {
		fmt.Println("Ping 成功")
	}

	ctx := context.TODO()
	cnn, err := db.Conn(ctx)

	if err != nil {
		fmt.Println("Connect 失敗")
		return
	} else {
		fmt.Println("Connect 成功")
	}

	fmt.Println("cnn", cnn)

	res, _ := db.Query("SHOW DATABASES")
	fmt.Println("----------------databases----------------")
	var database string
	for res.Next() {
		res.Scan(&database)
		fmt.Println(database)
	}
	res.Close()

}
