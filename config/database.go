package config

import (
	"book-api/models"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	fmt.Println("Conectando ao banco de dados em:", dsn)
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

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Erro ao conectar ao banco de dados 'booksdb': " + err.Error())
	}

	database.AutoMigrate(&models.Book{})
	DB = database
	fmt.Println("Conexão com o banco de dados 'booksdb' estabelecida com sucesso!")
}
