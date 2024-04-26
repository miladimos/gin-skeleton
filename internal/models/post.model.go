package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Post struct {
	Title string         `gorm:"column:title;unique" json:"title"`
	Slug  sql.NullString `gorm:"column:slug;unique" json:"slug"`
	Body  string         `gorm:"column:body" json:"body"`

	gorm.Model
}

func (u *Post) Table() string {
	return "posts"
}
