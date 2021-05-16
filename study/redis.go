package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {
	conn := pool.Get()

	r, err := redis.Strings(conn.Do("lrange", "city", "0", "-1"))
	if err != nil {
		fmt.Println("lrange error:", err)
		return
	}
	fmt.Println(r)

}
