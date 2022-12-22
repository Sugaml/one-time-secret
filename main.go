package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/berrybytes/simplesecrets/internal/controller/api"
	db "github.com/berrybytes/simplesecrets/internal/model/sqlc"
	"github.com/berrybytes/simplesecrets/util"
	_ "github.com/lib/pq"
)

func main() {
	// config, err := util.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("Cannot load config", err)
	// }
	config, err := util.LoadEnvConfig()
	fmt.Println(config)
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server : ", err)
	}
	err = server.Start(":3000")
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
