package models

import (
	"gorm.io/gorm"
)

type Activation struct {
	Code string `gorm:"column:code;uniqueIndex" json:"code,omitempty" validate:"required"`
	gorm.Model
}

func (u *Activation) Table() string {
	return "activations"
}
