package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(200)" json:"name"`
	Email    string `gorm:"type:varchar(200)" json:"email"`
	Role     string `gorm:"type:varchar(200)" json:"role"`
	Password string `gorm:"type:varchar(300)" json:"password"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
