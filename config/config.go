package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type setting struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func ImportSetting() setting {
	var result setting;
	err := godotenv.Load(".env");
	if err != nil {
		return setting{};
	}

	result.Host 	= os.Getenv("host");
	result.User 	= os.Getenv("user");
	result.Password = os.Getenv("password");
	result.DBName 	= os.Getenv("dbname");
	result.Port 	= os.Getenv("port");

	return result;
}

func ConnectDB(s setting) (*gorm.DB, error) {
	cnnStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", s.Host, s.User, s.Password, s.DBName, s.Port)
	db, err := gorm.Open(postgres.Open(cnnStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "try.",
		},
	});

	if err != nil {
		return nil, err
	}

	return db, nil;
}