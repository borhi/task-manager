package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"task-manager/adapters"
	"task-manager/handlers"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	logger, _ := zap.NewProduction()
	mode, exists := os.LookupEnv("MODE")
	if exists && mode == "development" {
		logger, _ = zap.NewDevelopment()
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Error(err.Error())
		}
	}()

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	postgresUser, exist := os.LookupEnv("POSTGRES_USER")
	if !exist {
		postgresUser = "postgres"
	}
	postgresPassword, exist := os.LookupEnv("POSTGRES_PASSWORD")
	if !exist {
		postgresPassword = "manager"
	}
	postgresDb, exist := os.LookupEnv("POSTGRES_DB")
	if !exist {
		postgresDb = "task_manager"
	}
	postgresHost, exist := os.LookupEnv("POSTGRES_HOST")
	if !exist {
		postgresHost = "localhost"
	}

	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s sslmode=disable",
		postgresUser, postgresPassword, postgresDb, postgresHost,
	)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	postgresAdapter := adapters.PostgresAdapter{Connection: conn}
	r := router{
		projectHandler: handlers.NewProjectHandler(postgresAdapter),
		columnHandler:  handlers.NewColumnHandler(postgresAdapter),
		taskHandler:    handlers.NewTaskHandler(postgresAdapter),
		commentHandler: handlers.NewCommentHandler(postgresAdapter),
	}
	log.Fatal(http.ListenAndServe(":8080", r.InitRouter()))
}
