package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	Completed   bool   `gorm:"default:false"`
	UserID      uint   // foreign key, references User table
	User        User   `gorm:"foreignKey:UserID" json:"-"`
}
