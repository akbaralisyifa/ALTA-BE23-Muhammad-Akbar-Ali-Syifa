package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type setting struct{
	Host, User, Password, DBName, Port string
}


func ImportSetting() setting{
	var result setting

	err := godotenv.Load(".env");

	if err != nil {
		return setting{}
	}

	result.Host 	= os.Getenv("DB_HOST");
	result.User		= os.Getenv("DB_USER");
	result.Password = os.Getenv("DB_PASSWORD");
	result.DBName 	= os.Getenv("DB_NAME");
	result.Port 	= os.Getenv("DB_PORT");

	return result;
};


func ConnectDB(s setting)(*gorm.DB, error){
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", s.Host, s.User, s.Password, s.DBName, s.Port)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{});

	if err != nil {
		return nil, err
	}

	return db, nil;
}