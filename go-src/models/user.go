package models

import "time"

//User 用户信息
type User struct {
	ID            string
	Name          string
	Password      string
	Phone         string
	MobilPhone    string
	Email         string
	PostAddress   string
	LastLoginTime time.Time
	LoginIP       string
}
