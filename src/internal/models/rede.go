package models

import "gorm.io/gorm"

type Rede struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Nome  string `json:"nome"`
	Label string `json:"label"`
}
