package models

import "time"

type Task struct {
	ID       int       `gorm:"primary_key;AUTO_INCREMENT"` // Specify auto-increment
	TaskName string    `gorm:"not null"`
	Status   string    `gorm:"not null;default:'Pending'"`
	Date     time.Time `gorm:"not null;type:date"`
}
