package models

type Role struct {
	BaseModel
	Name      string `gorm:"type:varchar(20);not null" json:"name"`
	UserRoles *[]UserRole
}
