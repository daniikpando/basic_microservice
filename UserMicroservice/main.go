package main

import (
	"flag"
	"log"
	"github.com/daniel/PruebaBackend/UserMicroservice/routes"
	"github.com/daniel/PruebaBackend/UserMicroservice/migration"
)

// Puerto donde correrá la aplicación
var (
	SERVER_PORT = ":8080"
)

// Función principal

func main() {

	// Generar la migración a la base de datos si el desarrollador lo requiere
	// Para ello al momento de ejecutar la aplicación debe agregar la bandera --migrate yes
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera migración a la base de datos")

	flag.Parse()

	if migrate == "yes" {
		log.Println("Comenzo la migración")
		migration.Migrate()
		log.Println("Termino la migración")
	}

	// Configurar las rutas del servicio
	e := routes.InitRouter()

	// Ejecutar el servidor en el puerto establecido
	e.Logger.Fatal(e.Start(SERVER_PORT))

}
