package models

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        int          `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time    `gorm:"TYPE:TIMESTAMP WITH TIME ZONE;DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt sql.NullTime `gorm:"TYPE:TIMESTAMP WITH TIME ZONE;"`
	DeleteAt  sql.NullTime `gorm:"TYPE:TIMESTAMP WITH TIME ZONE;"`

	CreatedBy int            `gorm:"NOT NULL"`
	UpdatedBy *sql.NullInt64 `gorm:"NULL"`
	DeletedBy *sql.NullInt64 `gorm:"NULL"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var userId = -1
	if value != nil {
		userId = int(value.(float64))
	}
	m.CreatedAt = time.Now().UTC()
	m.CreatedBy = userId
	return
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var userId = &sql.NullInt64{Valid: true}
	if value != nil {
		userId = &sql.NullInt64{Valid: true, Int64: int64(value.(float64))}
	}
	m.UpdatedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	m.UpdatedBy = userId
	return
}

func (m *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var userId = &sql.NullInt64{Valid: true}
	if value != nil {
		userId = &sql.NullInt64{Valid: true, Int64: int64(value.(float64))}
	}
	m.DeleteAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	m.DeletedBy = userId
	return
}
