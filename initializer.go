package main

import (
	"database/sql"
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
)

func connectRedis() (c redis.Conn) {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	return c
}

func connectMysql() (c *sql.DB) {
	db, err := sql.Open("mysql", "root:password@/dbname")
	if err != nil {
		panic(err)
	}
	return db
}

var redis_conn = connectRedis()
var mysql_conn = connectMysql()

func main() {
}

func insertToRedis(i int) {
	n, err := redis_conn.Do("HSET", "honeyqa_shard", "key", "value")
	if err != nil {
		panic(err)
	}
  fmt.Println(n)
}
