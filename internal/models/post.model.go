package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Post struct {
	Title     string         `gorm:"column:uniqueIndex;unique" json:"title,omitempty" validate:"required"`
	Slug      sql.NullString `gorm:"column:slug;uniqueIndex" json:"slug"`
	Body      string         `gorm:"column:body" json:"body"`
	ViewCount uint           `gorm:"column:view_count" json:"view_count"`
	gorm.Model
}

func (u *Post) Table() string {
	return "posts"
}
