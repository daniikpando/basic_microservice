package migration

import (
	"github.com/daniel/PruebaBackend/UserMicroservice/connections"
	"github.com/daniel/PruebaBackend/UserMicroservice/models"
)



// Migrate se encarga de generar la migración a la base de datos
func Migrate() {
	db := connections.GetConnection()
	db.CreateTable(&models.User{})
}