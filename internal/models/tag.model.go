package models

import (
	"gorm.io/gorm"
)

type Tag struct {
	Title string  `gorm:"column:title;uniqueIndex" json:"title,omitempty" validate:"required"`
	Slug  *string `gorm:"column:slug;uniqueIndex" json:"slug"`
	gorm.Model
}

func (u *Tag) Table() string {
	return "tags"
}
