package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func query(db *sql.DB) {
	var (
		id         int
		coursename string
		price      float64
		instructor string
	)
	for {
		var inputID int
		fmt.Print("enter your ID : ")
		fmt.Scan(&inputID)

		query := "SELECT id, coursename, price, instructor FROM onlinecourse WHERE id = ?"
		if err := db.QueryRow(query, inputID).Scan(&id, &coursename, &price, &instructor); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, coursename, price, instructor)
	}

}

func main() {
	db, err := sql.Open("mysql", "root:12345ter@tcp(127.0.0.1:3306)/coursebd")
	if err != nil {
		fmt.Println("failed to connect")
	} else {
		fmt.Println("connect successfuly")
	}
	query(db)
}
