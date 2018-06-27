package main

import (
	"github.com/daniel/PruebaBackend/UserMicroservice/routes"
	"flag"
	"log"
	"github.com/daniel/PruebaBackend/UserMicroservice/migration"
)

var (
	SERVER_PORT = ":8080"
)
func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera migración a la base de datos")

	flag.Parse()

	if migrate == "yes" {
		log.Println("Comenzo la migración")
		migration.Migrate()
		log.Println("Termino la migración")
	}

	e := routes.InitRouter()
	e.Logger.Fatal(e.Start(SERVER_PORT))

}
