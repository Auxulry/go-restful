package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MochamadAkbar/go-restful/common"
	"github.com/MochamadAkbar/go-restful/config"
)

const (
	Addr           = ":5000"
	ReadTimeout    = 10 * time.Second
	WriteTimeout   = 10 * time.Second
	MaxHeaderBytes = 1 << 20
)

const (
	Hostname = "localhost"
	Port     = "5432"
	Username = "postgres"
	Password = "postgres"
	Dbname   = "go_restful"
)

func main() {
	ctx := context.Background()
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", Username, Password, Hostname, Port, Dbname)

	db := config.InitDB(ctx, url)
	defer db.Close()

	router := config.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello World"))
	})

	server := &http.Server{
		Addr:           "localhost:5000",
		Handler:        router,
		ReadTimeout:    ReadTimeout,
		WriteTimeout:   WriteTimeout,
		MaxHeaderBytes: MaxHeaderBytes,
	}

	badge := common.MessageColorized(common.Green, "ready")
	msg := fmt.Sprintf("[%s] started serve on [::]%s", badge, Addr)
	fmt.Println(msg)
	if err := server.ListenAndServe(); err != nil {
		badge = common.MessageColorized(common.Red, "stop")
		msg = fmt.Sprintf("[%s] server failed to start : %s", badge, err.Error())
		log.Fatalln(msg)
	}
}
