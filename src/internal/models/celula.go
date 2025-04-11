package models

import (
	"time"

	"gorm.io/gorm"
)

type Celula struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"ID"`
	IDRede      int    `json:"id_rede"`
	Nome        string `json:"nome"`
	Lider       string `json:"lider"`
	Supervisor  int    `json:"supervisor"`
	QtdMembros  int    `json:"qtd_membros"`
	Local       string `json:"local"`
	Rede        string `json:"rede"`
	DiaDaSemana string `json:"dia_da_semana"`
	Horario     string `json:"horario"`
}

type Encontro struct {
	gorm.Model
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"ID"`
	IDCelula         int       `json:"id_celula"`
	Data             time.Time `json:"data"`
	Pregador         string    `json:"pregador"`
	QtdPresentes     int       `json:"qtd_presentes"`
	QtdVisitantes    int       `json:"qtd_visitantes"`
	OfertaArrecadada float64   `json:"oferta_arrecadada"`
}

type EncontroBody struct {
	ID               uint      `json:"ID"`
	Data             time.Time `json:"data"`
	Pregador         string    `json:"pregador"`
	QtdPresentes     int       `json:"qtd_presentes"`
	QtdVisitantes    int       `json:"qtd_visitantes"`
	OfertaArrecadada float64   `json:"oferta_arrecadada"`
	MembrosPresentes []int     `json:"membros_presentes"`
}

type MembroCelula struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey;autoIncrement" json:"ID"`
	IDCelula int       `json:"id_celula"`
	Nome     string    `json:"nome"`
	Telefone string    `json:"telefone"`
	Email    string    `json:"email"`
	DataNasc time.Time `json:"data_nascimento"`
	Endereco string    `json:"endereco"`
	Batizado bool      `json:"batizado"`
}

type MembroCelulaEncontro struct {
	gorm.Model
	ID         uint `gorm:"primaryKey;autoIncrement" json:"ID"`
	IDCelula   int  `json:"id_celula"`
	IDEncontro int  `json:"id_encontro"`
	IDMembro   int  `json:"id_membro"`
}
