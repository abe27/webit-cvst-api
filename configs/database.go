package configs

import "gorm.io/gorm"

var (
	Store           *gorm.DB
	API_TRIGGER_URL string
	APP_NAME        string
	APP_DESCRIPTION string
	DB_HOST         string
	DB_PORT         int
	DB_NAME         string
	DB_USERNAME     string
	DB_PASSWORD     string
)
