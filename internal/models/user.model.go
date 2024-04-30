package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	Firstname string         `gorm:"column:firstname" json:"firstname"`
	Lastname  string         `gorm:"column:lastname" json:"lastname"`
	Email     string         `gorm:"column:email;unique" json:"email" validate:"required,email"`
	Mobile    sql.NullString `gorm:"column:mobile;unique" json:"mobile"`
	Password  string         `gorm:"column:password" json:"-"`

	gorm.Model
}

func (u *User) Table() string {
	return "users"
}
