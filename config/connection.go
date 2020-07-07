package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser,
	dbPassword,
	dbHost,
	dbPort,
	dbSchema string
)

func ConnectDB() *sql.DB {
	dbUser = os.Getenv("dbUser")
	dbPassword = os.Getenv("dbPassword")
	dbHost = os.Getenv("dbHost")
	dbPort = os.Getenv("dbPort")
	dbSchema = os.Getenv("dbSchema")

	//dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbSchema)

	connect := "root:Theresearcher_132@tcp(localhost:3306)/music_gallery"

	//masukan database dan address local host anda
	db, _ := sql.Open("mysql", connect)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db

}
