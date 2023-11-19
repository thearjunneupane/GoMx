package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func SetupDb() {

	var err error

	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// name := os.Getenv("DB_NAME")
	// user := os.Getenv("DB_USER")
	// pass := os.Getenv("DB_PASS")

	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, name)

	psqlInfo := os.Getenv("PSQL_connStr")

	// fullInfo := true
	// if len(host) < 1 {
	// 	log.Println("Failed !! Got Empty host")
	// 	fullInfo = false

	// } else {
	// 	log.Println("Success!! Got Host")
	// }

	// if len(port) < 1 {
	// 	log.Println("Failed !! Got Empty port")
	// 	fullInfo = false

	// } else {
	// 	log.Println("Success!! Got Port")
	// }

	// if len(user) < 1 {
	// 	log.Println("Failed !! Got Empty user")
	// 	fullInfo = false

	// } else {
	// 	log.Println("Success!! Got User")
	// }

	// if len(pass) < 1 {
	// 	log.Println("Failed !! Got Empty password")
	// 	fullInfo = false

	// } else {
	// 	log.Println("Success!! Got Password")
	// }

	// if len(name) < 1 {
	// 	log.Println("Failed !! Got Empty name")
	// 	fullInfo = false

	// } else {
	// 	log.Println("Success!! Got DBName")
	// }

	// if !fullInfo {
	// 	log.Fatalln("Failed !! Not Got Full psqlInfo:")
	// }

	// log.Println("Success!! Got Full psqlInfo")

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln("Failed !! Error While Connecting to Database: ", err)
	}

	log.Println("Success!! DB Opened")

	if err = DB.Ping(); err != nil {
		log.Fatalln("Failed !! Error While Pinging to Database: ", err)
	}

	log.Println("Success!! Database Connected Successfully")

	// sqlFile, err := os.ReadFile("db/init.sql")
	// if err != nil {
	// 	log.Fatalln("Failed !! Error While Reading SQL File: ", err)
	// }

	// script := string(sqlFile)

	// if _, err := DB.Exec(script); err != nil {
	// 	log.Fatalln("Failed !! Error While Executing SQL Script: ", err)
	// }

	// log.Println("Success!! SQL Executed Successfully")

}
