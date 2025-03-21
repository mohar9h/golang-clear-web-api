package models

type City struct {
	BaseModel
	Name      string  `gorm:"SIZE:10;TYPE:STRING;NOT NULL"`
	CountryId int     `gorm:"TYPE:INT"`
	Country   Country `gorm:"foreignKey:CountryId"`
}
