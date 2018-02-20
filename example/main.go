package main

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"database/sql"
	"github.com/mattn/go-sqlite3"
)

func main() {
	ver, _, _ := sqlite3.Version()
	fmt.Println("SQLite3 Version : ", ver)
	db, err := sql.Open("sqlite3", `D:\golang\src\github.com\iamGreedy\Inspire\test.db`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Open DB : ", db.Driver())
	sql, arg, err := sq.Select("*").From("KVDB").Where("Key=?", "Hello").ToSql()
	if err != nil {
		panic(err)
	}
	fmt.Println(sql, arg)
	stmt, err  := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	res, err := stmt.Query(arg...)
	if err != nil {
		panic(err)
	}
	defer res.Close()
	var i = -
	for res.Next(){
		dat, err := res.Columns()
		if err != nil {
			fmt.Println("Error : ", err)
		}
		fmt.Println(dat)
	}

}
