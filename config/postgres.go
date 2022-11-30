package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	host, user, password, dbname, port, dsn string
)

func ConnectDB() *gorm.DB {
	host = os.Getenv("DB_HOST")
	user = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	dbname = os.Getenv("DB_NAME")
	port = os.Getenv("DB_PORT")
	dsn = "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "public.",
		},
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print("Postgres server listening on port " + port + "...")

	return db
}
