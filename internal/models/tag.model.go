package models

import (
	"gorm.io/gorm"
)

type Tag struct {
	Title string  `gorm:"column:title;unique" json:"title,omitempty" validate:"required" binding:"required,min=3,max=150"`
	Slug  *string `gorm:"column:slug;unique" json:"slug"`
	gorm.Model
}

func (u *Tag) Table() string {
	return "tags"
}
