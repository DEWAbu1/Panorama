package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	rootDNS = "root:@/"
	dbName  = "panorama" // database name to be decided
)

var DNS = rootDNS + dbName + "?multiStatements=true&charset=utf8&parseTime=true"

func OpenDB() *sql.DB {

	db, err := sql.Open("mysql", DNS)
	if err != nil {
		log.Println(err)
	}

	return db
}

func CheckCount(rows *sql.Rows) (count int) {
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			log.Println(err)
		}
	}
	return count
}
