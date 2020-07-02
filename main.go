package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"log"
	"net/http"
	"task-manager/adapters"
	"task-manager/handlers"
)

func main() {
	logger := zap.NewExample()
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Error(err.Error())
		}
	}()

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	connStr := "user=postgres password=manager dbname=task_manager host=localhost sslmode=disable"
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
