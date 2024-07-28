package dbs

import (
	logs "auth/auth_back/pkg/logger"
	"log"
	"os"
	"time"

	notificationModels "auth/auth_back/models/notification"
	orgModels "auth/auth_back/models/organization"
	userModels "auth/auth_back/models/user"

	"github.com/spf13/viper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var l = logs.Logger{}

func InitDB() *gorm.DB {
	var errConnect error

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Error,
			Colorful: true, // Disable color
		},
	)

	db, errConnect := gorm.Open(postgres.New(postgres.Config{
		// DSN: viper.GetString("postgres.uri"),
		DSN:                  viper.GetString("postgres.uri_local"),
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		// DisableForeignKeyConstraintWhenMigrating: true,
		Logger: newLogger,
	})

	counter := 0

	for errConnect != nil {
		if counter > 5 {
			panic("Failed to connect database")
		}

		log.Println("Err connect DB, retrying after 15 sec")

		time.Sleep(15 * time.Second)

		db, errConnect = gorm.Open(postgres.Open(viper.GetString("postgres.uri")), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})

		counter++
	}

	log.Println("Connect to db success")

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)

	errMigrate := migrateSchemes(db)

	if errMigrate != nil {
		panic("Failed to migrate database")
	}

	return db
}

func migrateSchemes(db *gorm.DB) error {
	err := db.AutoMigrate(
		&orgModels.BusinessUniverse{},
		&orgModels.Company{},
		&orgModels.CompanyBank{},
		&orgModels.CompanyContactInfo{},
		&orgModels.Customer{},
		&orgModels.CustomerBank{},
		&orgModels.CustomerContactInfo{},
		&orgModels.CustomerUser{},
		&orgModels.ContactType{},
		&userModels.Role{},
		&userModels.User{},
		&notificationModels.MailConfrimationLinks{},
	)

	if err != nil {
		l.LogError(err.Error(), "pkg/dbs/dbs.go")
		return err
	}

	return nil
}
