package models

type Country struct {
	BaseModel
	Name   string `gorm:"SIZE:15;TYPE:STRING;NOT NULL"`
	Cities *[]City
}
