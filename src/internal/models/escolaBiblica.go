package models

import (
	"time"

	"gorm.io/gorm"
)

type Turma struct {
	gorm.Model
	ID         uint `gorm:"primaryKey;autoIncrement"`
	Nome       string
	Descricao  string
	Periodo    string
	DataInicio time.Time
	DataFim    time.Time
}

type Aula struct {
	gorm.Model
	ID           uint `gorm:"primaryKey;autoIncrement"`
	IDTurma      uint
	Data         time.Time
	Professor    string
	QTDPresentes int
}
