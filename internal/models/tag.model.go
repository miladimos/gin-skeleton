package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Tag struct {
	Title string         `gorm:"column:title;unique" json:"title,omitempty" validate:"required"`
	Slug  sql.NullString `gorm:"column:slug;unique" json:"slug"`
	gorm.Model
}

func (u *Tag) Table() string {
	return "tags"
}
