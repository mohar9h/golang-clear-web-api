package migrations

import (
	"github.com/mohar9h/golang-clear-web-api/config"
	"github.com/mohar9h/golang-clear-web-api/constants"
	"github.com/mohar9h/golang-clear-web-api/data/db"
	"github.com/mohar9h/golang-clear-web-api/data/models"
	logging2 "github.com/mohar9h/golang-clear-web-api/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = logging2.NewLogger(config.GetConfig())

func Up() {
	database := db.GetDBClient()

	var tables []interface{}

	country := models.Country{}
	city := models.City{}
	user := models.User{}
	role := models.Role{}
	userRole := models.UserRole{}

	tables = addNewMigrate(database, country, tables)
	tables = addNewMigrate(database, city, tables)
	tables = addNewMigrate(database, user, tables)
	tables = addNewMigrate(database, role, tables)
	tables = addNewMigrate(database, userRole, tables)

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		logger.Error(logging2.Postgres, logging2.Migration, err.Error(), nil)
	}
	createDefaultInformation(database)
	logger.Info(logging2.Postgres, logging2.Migration, "Tables created", nil)
}

func createDefaultInformation(database *gorm.DB) {

	adminRole := models.Role{Name: constants.AdminRoleName}
	createRoleIfNotExists(database, &adminRole)

	defaultRole := models.Role{Name: constants.DefaultRoleName}
	createRoleIfNotExists(database, &defaultRole)

	user := models.User{Username: constants.DefaultUserName, FirstName: "Test", LastName: "Test", Email: "test@test.com", MobileNumber: "09361562428"}
	password := "12345678"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	createUserIfNotExists(database, &user, adminRole.Id)
}

func createRoleIfNotExists(database *gorm.DB, role *models.Role) {
	exists := 0

	database.Model(&models.Role{}).
		Select("1").
		Where("name = ?", role.Name).
		First(&exists)

	if exists == 0 {
		database.Create(role)
	}
}

func createUserIfNotExists(database *gorm.DB, user *models.User, roleId int) {
	exists := 0

	database.Model(&models.User{}).
		Select("1").
		Where("username = ?", user.Username).
		First(&exists)

	if exists == 0 {
		database.Create(user)
		userRole := models.UserRole{UserId: user.Id, RoleId: roleId}
		database.Create(&userRole)
	}
}

func addNewMigrate(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func Down() {

}
