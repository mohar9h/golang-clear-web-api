package models

type User struct {
	BaseModel
	Username     string `gorm:"TYPE:STRING;SIZE20;NOT NULL;UNIQUE"`
	FirstName    string `gorm:"TYPE:STRING;SIZE:15;NOT NULL"`
	LastName     string `gorm:"TYPE:STRING;SIZE:25;NOT NULL"`
	Email        string `gorm:"TYPE:STRING;SIZE:64;NULL;UNIQUE;DEFAULT:NULL"`
	MobileNumber string `gorm:"TYPE:STRING;SIZE:11;NULL;UNIQUE;DEFAULT:NULL"`
	Password     string `gorm:"TYPE:STRING;SIZE:64;NOT NULL"`
	Enabled      bool   `gorm:"TYPE:BOOLEAN;DEFAULT:true"`
	UserRoles    *[]UserRole
}
