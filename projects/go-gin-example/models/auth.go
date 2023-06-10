package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Auth struct {
	// ID       int    `gorm:"primary_key" json:"id"`
	Model

	Username string `json:"username"`
	Password string `json:"password"`
}

func (auth *Auth) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (auth *Auth) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func CheckAuth(username, password string) (bool, error) {
	var auth Auth
	err := db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if auth.ID > 0 {
		return true, nil
	}
	return false, nil
}

func RegistAuth(username, password string) bool {
	auth := Auth{
		Username: username,
		Password: password,
	}
	db.AutoMigrate(&Auth{})
	if err := db.Create(&auth).Error; err != nil {
		fmt.Println("xxxxxxfcf: RegistAuth 失败", err)
		return false
	}
	return true
}

// 判断Tag是否存在
func ExistAuthByName(username string) (bool, error) {
	var auth Auth
	err := db.Select("id").Where("username = ?", username).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if auth.ID > 0 {
		return true, nil
	}
	return false, nil
}
