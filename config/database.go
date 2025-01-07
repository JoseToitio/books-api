package config

import (
	"book-api/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=root port=5432 sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("Erro ao conectar ao PostgreSQL: " + err.Error())
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE booksdb")
	if err != nil {
		log.Println("Banco de dados já existe ou erro ao criar:", err)
	} else {
		fmt.Println("Banco de dados 'booksdb' criado com sucesso!")
	}

	dsn = "host=localhost user=postgres password=root dbname=booksdb port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Erro ao conectar ao banco de dados 'booksdb': " + err.Error())
	}

	database.AutoMigrate(&models.Book{})
	DB = database
	fmt.Println("Conexão com o banco de dados 'booksdb' estabelecida com sucesso!")
}
