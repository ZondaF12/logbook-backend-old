package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ZondaF12/logbook-backend/internal/models"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

type (
	Service interface {
		Health() map[string]string
		Handlers
	}

	Handlers interface {
		Authentication
		Users
	}

	Authentication interface {
		AddAuthenticationToDB(params models.UserAuth) map[string]string
		GetAuthenticationByEmail(email string) models.UserAuth
		GetAuthenticationByID(id uuid.UUID) models.UserAuth
	}

	Users interface {
		GetUsers() []models.User
		GetUserByID(id uuid.UUID) (models.User, error)
		AddNewUserToDB(user models.User) map[string]string
		UpdateUserByID(id uuid.UUID, user models.User) (models.User, error)
	}

	service struct {
		db *sql.DB
	}
)

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

func New() Service {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	s := &service{db: db}
	return s
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}
