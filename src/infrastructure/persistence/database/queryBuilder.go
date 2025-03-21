package database

import "gorm.io/gorm"

type PreloadEntity struct {
	Entity string
}

func Preload(db *gorm.DB, preloads []PreloadEntity) *gorm.DB {
	for _, item := range preloads {
		db = db.Preload(item.Entity)
	}
	return db
}
