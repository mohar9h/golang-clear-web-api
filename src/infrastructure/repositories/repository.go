package repositories

import (
	"github.com/mohar9h/golang-clear-web-api/infrastructure/persistence/database"
	"github.com/mohar9h/golang-clear-web-api/pkg/logging"
	"gorm.io/gorm"
)

type BaseRepository[TEntity any] struct {
	database *gorm.DB
	logger   logging.Logger
	preloads []database.PreloadEntity
}
