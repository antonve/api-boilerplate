package migrations

import (
	"fmt"
	"log"
	"os"

	"github.com/antonve/api-boilerplate/config"
	"github.com/antonve/api-boilerplate/models"

	"github.com/DavidHuie/gomigrate"
)

func getMigrator() (*gomigrate.Migrator, error) {
	// @TODO: change how migration path is handled
	appPath := os.Getenv("APP_PATH")
	if appPath == "" {
		appPath = fmt.Sprintf("%s/src/github.com/antonve/api-boilerplate", os.Getenv("GOPATH"))
	}

	return gomigrate.NewMigrator(models.GetSQLDatabase(), gomigrate.Postgres{}, fmt.Sprintf("%s/%s", appPath, config.GetConfig().MigrationsPath))
}

// Migrate migrates the database
func Migrate() error {
	log.Printf("Migrating in environment: %s", config.GetConfig().Environment)
	migrator, err := getMigrator()

	if err != nil {
		return err
	}

	err = migrator.Migrate()

	return err
}

// Destroy the current environment's database
func Destroy() error {
	if config.GetConfig().Environment == config.Environments["prod"] {
		return fmt.Errorf("Cannot destroy production.")
	}

	db := models.GetSQLConnection()
	defer db.Close()

	// Drop database
	_, err := db.Exec("DROP DATABASE IF EXISTS " + config.GetConfig().Database)
	if err != nil {
		return err
	}

	return nil
}

// Create a new database
func Create() error {
	db := models.GetSQLConnection()
	defer db.Close()

	// Create database
	_, err := db.Exec("CREATE DATABASE " + config.GetConfig().Database)
	if err != nil {
		return err
	}

	return nil
}
