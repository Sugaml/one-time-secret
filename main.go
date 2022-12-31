package main

import (
	"database/sql"
	"log"

	"github.com/berrybytes/simplesecrets/docs"
	"github.com/berrybytes/simplesecrets/internal/controller/api"
	db "github.com/berrybytes/simplesecrets/internal/model/sqlc"
	"github.com/berrybytes/simplesecrets/util"
	_ "github.com/lib/pq"
)

// @Onetime-Secret-api
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email info.onetimesecret.com

// @host petstore.swagger.io:8080
// @BasePath /v2

func main() {
	// config, err := util.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("Cannot load config", err)
	// }
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	config, err := util.LoadEnvConfig()
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}
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
