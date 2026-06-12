package main

import (
	"log"

	conf "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/config"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/database"
)

func main() {
	db := conf.MySQLConn()

	database.Migrate()
	database.Seeder(db)

	log.Println("Migration & Seeder Success")
}
