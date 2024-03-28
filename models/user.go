package models

import (
	"crypto/sha1"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"github.com/jinzhu/gorm"
	"myBulebell/pkg/utils"
	"strings"
)

const (
	Active = iota + 1
	NotActivated
	Baned
)

type User struct {
	gorm.Model
	UserId   int64  `gorm:"unique;not null"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Gender   int    `gorm:"default:0"`
	Status   int
}

func init() {
	gob.Register(User{})
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	salt := utils.RandomString(16)
	h := sha1.New()
	_, err := h.Write([]byte(password + salt))
	bs := hex.EncodeToString(h.Sum(nil))
	if err != nil {
		return err
	}
	// 存储 salt 值和摘要，":"分割
	user.Password = salt + ":" + bs
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) (bool, error) {
	passwordStore := strings.Split(user.Password, ":")
	if len(passwordStore) != 2 {
		return false, errors.New("unknown password format")
	}
	h := sha1.New()
	_, err := h.Write([]byte(password + passwordStore[0]))
	bs := hex.EncodeToString(h.Sum(nil))
	if err != nil {
		return false, err
	}
	return bs == passwordStore[1], nil
}

func (user *User) SetStatus(status int) {
	DB.Model(&user).Update("status", status)
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUserId(userId int64) (*User, error) {
	var user User
	if err := DB.Where("user_id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
