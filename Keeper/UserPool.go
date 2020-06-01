package Keeper

import (
	"fmt"
	"log"
)


type UserPool struct {
	mp map[string] *User
}

func NewUserPool() *UserPool {
	return &UserPool{map[string]*User{}}
}

// 增加/更新 用户连接
func (up *UserPool) AddUser(user *User) {
	log.Printf("user %s is login", user.name)
	up.mp[user.name] = user
}

// 用户断开连接
func (up *UserPool) DeleteUser(name string) error {

	if up.mp[name] == nil {
		return fmt.Errorf("%s is not connected" , name)
	}
	up.mp[name] = nil
	return nil
}

// 获取name的对应用户
func (up *UserPool) GetUser(name string) *User {
	return up.mp[name]
}
