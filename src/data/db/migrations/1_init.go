package migrations

import (
	"github.com/mohar9h/golang-clear-web-api/config"
	"github.com/mohar9h/golang-clear-web-api/data/db"
	"github.com/mohar9h/golang-clear-web-api/data/models"
	"github.com/mohar9h/golang-clear-web-api/logging"
)

var logger = logging.NewLogger(config.GetConfig())

func Up() {
	database := db.GetDBClient()

	var tables []interface{}

	country := models.Country{}
	city := models.City{}

	if !database.Migrator().HasTable(country) {
		tables = append(tables, country)
	}

	if !database.Migrator().HasTable(city) {
		tables = append(tables, city)
	}

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		logger.Error(logging.Postgres, logging.Migration, err.Error(), nil)
	}

	logger.Info(logging.Postgres, logging.Migration, "Tables created", nil)
}

func Down() {

}
