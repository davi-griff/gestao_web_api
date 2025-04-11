package models

import (
	"time"

	"gorm.io/gorm"
)

type Culto struct {
	gorm.Model
	ID                  uint `gorm:"primaryKey;autoIncrement"`
	Nome                string
	Data                time.Time
	Pastor              string
	QTDJovens           int
	QTDAdultos          int
	QTDCriancas         int
	QTDVisitantes       int
	QTDBatismo          int
	QTDConversao        int
	QTDOfertaArrecadada int
}

type MetricasCulto struct {
	gorm.Model
	ID               uint `gorm:"primaryKey;autoIncrement"`
	Culto            uint
	Data             time.Time
	Pastor           string
	Jovens           int
	Mulheres         int
	Homens           int
	Adultos          int
	Criancas         int
	Visitantes       int
	OfertaArrecadada int
}
