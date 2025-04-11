package models

import "gorm.io/gorm"

type Pastor struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Nome     string `json:"nome"`
	Telefone string `json:"telefone"`
	Email    string `json:"email"`
}

type Supervisor struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Nome     string `json:"nome"`
	Telefone string `json:"telefone"`
	Email    string `json:"email"`
}

type Lider struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Nome     string `json:"nome"`
	Telefone string `json:"telefone"`
	Email    string `json:"email"`
}

type Membro struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Nome     string `json:"nome"`
	Telefone string `json:"telefone"`
	Email    string `json:"email"`
}

type Aluno struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Nome     string `json:"nome"`
	Telefone string `json:"telefone"`
}
