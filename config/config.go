package config

import (
	"fmt"
	"mini_project/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {

	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "root",
		"DB_Port":     "3306",
		"DB_Host":     "localhost",
		"DB_Name":     "mini_project",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"],
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	InitialMigration()
	return DB
}

func InitialMigration() {
	DB.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.Cart{},
		&model.CartProduct{},
	)
}
