package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/MochamadAkbar/go-restful/common/colorize"
	"github.com/MochamadAkbar/go-restful/config"
	"github.com/MochamadAkbar/go-restful/exception"
	"github.com/MochamadAkbar/go-restful/injector"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env.local")
	if err != nil {
		panic("failed to load .env file")
	}
}

func main() {
	ctx := context.Background()
	urlString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db := config.NewDB(ctx, urlString)
	defer db.Close()

	router := config.NewRouter()
	router.Route("/api/v1", func(r chi.Router) {
		_ = injector.InitializeUserService(db, r)
	})
	server := config.NewServer(router)

	badge := colorize.MessageColorized(colorize.Green, "ready")
	msg := fmt.Sprintf("[%s] started serve on [::]%s", badge, ":5000")
	log.Println(msg)
	if err := server.ListenAndServe(); err != nil {
		panic(exception.NewException("[stop] server failed to start : " + err.Error()))
	}
}
