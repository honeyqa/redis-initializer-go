package main

import (
	"database/sql"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
)

func connectRedis() (c redis.Conn) {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	return c
}

func connectMysql() (c *sql.DB) {
	db, err := sql.Open("mysql", "root:password@/dbname")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}

var redis_conn = connectRedis()
var mysql_conn = connectMysql()

func main() {

}
