package main

import (
	"gomx/config"
	"gomx/db"
	"gomx/router"
)

func main() {
	config.Configure()
	db.SetupDb()
	router.SetupAndRun()
}
