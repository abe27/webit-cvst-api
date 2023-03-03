package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cvst/apiv1/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	// initial database
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", os.Getenv("MSSQL_USER"), os.Getenv("MSSQL_PASSWORD"), os.Getenv("MSSQL_HOST"), os.Getenv("MSSQL_PORT"), os.Getenv("MSSQL_DATABASE"))
	// db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	configs.Store, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		DisableAutomaticPing:                     true,
		DisableForeignKeyConstraintWhenMigrating: false,
		SkipDefaultTransaction:                   true,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tbt_", // table name prefix, table for `User` would be `t_users`
			SingularTable: false,  // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false,  // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "Cid"),
		},
	})

	if err != nil {
		panic("Failed to connect to database")
	}
}

func main() {
	// Create config variable
	config := fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "CVST Server API Service", // add custom server header
		AppName:       "API Version 1.0",
		BodyLimit:     10 * 1024 * 1024, // this is the default limit of 10MB
	}

	fmt.Println(config)
}
