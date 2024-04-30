package models

import (
	"gorm.io/gorm"
)

type Activation struct {
	Code string `gorm:"column:code;uniqueIndex" json:"code,omitempty"`
	gorm.Model
}

func (u *Activation) Table() string {
	return "activations"
}
