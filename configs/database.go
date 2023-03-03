package configs

import "gorm.io/gorm"

var (
	Store           *gorm.DB
	API_TRIGGER_URL string
	APP_NAME        string
	APP_DESCRIPTION string
)
