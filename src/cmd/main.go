package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Celula struct {
	gorm.Model
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Nome        string
	Lider       string
	Supervisor  string
	QtdMembros  int
	Local       string
	Rede        string
	DiaDaSemana string
	Horario     string
}

func main() {
	dsn := "host=localhost user=postgres password=admin dbname=gestao port=5433 sslmode=disable"

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	db.AutoMigrate(&Celula{})

	hmacSecret := os.Getenv("HMAC_SECRET")

	if hmacSecret == "" {
		log.Fatal("HMAC secret is not set")
	}

	router := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}

	router.Use(cors.New(corsConfig))

	// router.POST("/secret", authMiddleware(hmacSecret), secretRouteHandler())
	router.GET("/celulas", getCelulas)
	router.POST("/celulas", createCelula)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}

}

var emailCtxKey = "email"

func authMiddleware(hmacSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		email, err := parseJWTToken(token, []byte(hmacSecret))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		log.Printf("Authenticated user: %s", email)
		ctx := context.WithValue(c, emailCtxKey, email)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}

}

func secretRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.GetString(emailCtxKey)
		c.JSON(http.StatusOK, gin.H{"message": "Secret message", "email": email})
	}
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func parseJWTToken(token string, hmacSecret []byte) (string, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})

	if err != nil {
		return "", err
	} else if claims, ok := t.Claims.(*Claims); ok {
		return claims.Email, nil
	} else {
		return "", fmt.Errorf("invalid token")
	}

}

type CelulaJSON struct {
	ID          uint   `json:"id"`
	Nome        string `json:"nome"`
	Lider       string `json:"lider"`
	Supervisor  string `json:"supervisor"`
	QtdMembros  int    `json:"qtdMembros"`
	Local       string `json:"local"`
	Rede        string `json:"rede"`
	DiaDaSemana string `json:"diaDaSemana"`
	Horario     string `json:"horario"`
}

func getCelulas(c *gin.Context) {

	var celulas []Celula
	db.Order("nome DESC").Find(&celulas)

	var celulasJSON []CelulaJSON
	for _, celula := range celulas {
		celulasJSON = append(celulasJSON, CelulaJSON{
			ID:          celula.ID,
			Nome:        celula.Nome,
			Lider:       celula.Lider,
			Supervisor:  celula.Supervisor,
			QtdMembros:  celula.QtdMembros,
			Local:       celula.Local,
			Rede:        celula.Rede,
			DiaDaSemana: celula.DiaDaSemana,
			Horario:     celula.Horario,
		})
	}

	c.JSON(http.StatusOK, celulasJSON)
}

func createCelula(c *gin.Context) {
	var celula CelulaJSON
	if err := c.ShouldBindJSON(&celula); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	celulaDB := Celula{
		ID:          celula.ID,
		Nome:        celula.Nome,
		Lider:       celula.Lider,
		Supervisor:  celula.Supervisor,
		QtdMembros:  celula.QtdMembros,
		Local:       celula.Local,
		Rede:        celula.Rede,
		DiaDaSemana: celula.DiaDaSemana,
		Horario:     celula.Horario,
	}

	result := db.Create(&celulaDB)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, celulaDB)
}
