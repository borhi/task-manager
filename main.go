package main

import (
	"log"
	"net/http"
	"go.uber.org/zap"
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

	r := router{}
	log.Fatal(http.ListenAndServe(":8080", r.InitRouter()))
}
