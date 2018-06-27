package connections

import (
	"github.com/jinzhu/gorm"
	"sync"
	"fmt"
	"github.com/daniel/PruebaBackend/UserMicroservice/configuration"
	"log"
	_"github.com/jinzhu/gorm/dialects/postgres"
)

var (
	connect sync.Once
	db      *gorm.DB
)

func GetConnection() *gorm.DB {
	connect.Do(func() {
		var err error
		config := configuration.GetConfiguration()

		stringConnection := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
			config.Server,
			config.User,
			config.Database,
			config.Password,
		)

		db, err = gorm.Open(config.Engine, stringConnection)

		if err != nil {
			log.Fatal(err)
		}
	})

	return db
}