package main

import (
	"github.com/OdaDaisuke/go_utils/db"
	"log"
)

func main() {
	dbCtx, err := db.Init("mysql", &db.DataBaseSource{
		User: "root",
		Password: "",
		Host: "localhost",
		Port: "3306",
		DBName: "app",
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("db ctx -> ", dbCtx)
}
