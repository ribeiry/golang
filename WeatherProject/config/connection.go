package config

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/patrickmn/go-cache"
)

var DB *sql.DB
var Caching *cache.Cache

func Init() {
	var err error
	DB, err = sql.Open("mysql", DBConnection)

	if err != nil {
		panic(err)

	}
	err = DB.Ping()
	if err != nil {
		panic(err.Error())
	}
}

func InitCache() {
	Caching = cache.New(2*time.Minute, 10*time.Minute)
}
