package model

import (
	"httpproj/db"
	"time"
)

//规则：表名是结构体名字小写加s或者es，字段名称首字母小写，后面的大写字母使用_代替
type User struct {
	ID          int
	Name        string
	Password    string
	Identifier  string
	CreatedTime time.Time
}

func (p *User) Add() bool {
	tx := db.DB.Begin()
	if err := tx.Create(p).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}
func GetUserCount(n string) int {
	var count int
	db.DB.Model(&User{}).
		Where("Name = ?", n).
		Count(&count)
	return count
}
func GetUserLoginCount(name string, pwd string) int {
	var count int
	db.DB.Model(&User{}).
		Where("Name = ? and Password = ?", name, pwd).
		Count(&count)
	return count
}
func GetUserByName(name string) *User {
	var pro User
	var count int
	db.DB.Model(&User{}).
		Where("name = ?", name).
		First(&pro).
		Count(&count)
	if count == 0 {
		return nil
	} else {
		return &pro
	}
}
