package main

import (
	"context"
	"fmt"
	"log"

	"github.com/MochamadAkbar/go-restful/common"
	"github.com/MochamadAkbar/go-restful/config"
	"github.com/MochamadAkbar/go-restful/injector"
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

	db := config.NewDB(ctx, url)
	defer db.Close()

	server := injector.InitializeServer(db)

	badge := common.MessageColorized(common.Green, "ready")
	msg := fmt.Sprintf("[%s] started serve on [::]%s", badge, ":5000")
	fmt.Println(msg)
	if err := server.ListenAndServe(); err != nil {
		badge = common.MessageColorized(common.Red, "stop")
		msg = fmt.Sprintf("[%s] server failed to start : %s", badge, err.Error())
		log.Fatalln(msg)
	}
}
