package model

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/gostudy/project2/common/message"
)

type UserDao struct {
	pool *redis.Pool
}

var (
	MyUserDao *UserDao
)

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (this *UserDao) getUserById(conn redis.Conn, id string) (user *User, err error) {

	res, err := redis.String(conn.Do("hget", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ErrorUserNoExitsts
		}
		return
	}

	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json u err", err)
		return
	}
	return
}

func (this *UserDao) Login(userId string, userPwd string) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	if user.UserPwd != userPwd {
		err = ErrorUserPwd
		return
	}
	return
}

func (this *UserDao) Register(user *message.User) (err error) {
	conn := this.pool.Get()
	defer conn.Close()

	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ErrorUserExitsts
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		return
	}

	_, err = conn.Do("hset", "users", user.UserId, data)
	if err != nil {
		fmt.Println("conn error: ", err)
		return
	}
	return

}
