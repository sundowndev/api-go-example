/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50);not null" json:"name"`
	Email    string `gorm:"type:varchar(100);unique_index;not null" json:"email"`
	Password string `gorm:"type:text;not null" json:"-"`
}

func (u *User) BeforeSave() (err error) {
	var hash []byte

	if u.Password != "" {
		hash, err = bcrypt.GenerateFromPassword([]byte(u.Password), 12)

		if err == nil {
			u.Password = string(hash)
		}
	}
	return err
}

func (u *User) Compare(attempt string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(attempt))
}
