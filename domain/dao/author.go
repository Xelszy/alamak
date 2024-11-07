package dao

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
}
