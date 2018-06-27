package migration

import (
	"github.com/daniel/PruebaBackend/UserMicroservice/connections"
	"github.com/daniel/PruebaBackend/UserMicroservice/models"
)

func Migrate() {
	db := connections.GetConnection()
	db.CreateTable(&models.User{})
}