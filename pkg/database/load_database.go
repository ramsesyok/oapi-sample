package database

import (
	"bufio"
	"encoding/json"
	"io"
	"landmarks/pkg/openapi"
	"log/slog"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func LoadDatabase(dbFile string, inputFile string) (err error) {

	if contents, err := readDataFile(inputFile); err != nil {
		return err
	} else {
		return saveDatabase(contents, dbFile)
	}
}

func readDataFile(inputFile string) (io.Reader, error) {

	if file, err := os.Open(inputFile); err != nil {
		return file, err
	} else {
		return file, err
	}
}

func saveDatabase(contents io.Reader, dbFile string) (err error) {
	if db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{}); err != nil {
		return err
	} else {
		if err := db.Migrator().DropTable(&openapi.Landmark{}); err != nil {
			return err
		}
		if err := db.AutoMigrate(&openapi.Landmark{}); err != nil {
			return err
		}
		scanner := bufio.NewScanner(contents)
		for scanner.Scan() {
			line := scanner.Bytes()
			landmark := openapi.Landmark{}
			if err := json.Unmarshal(line, &landmark); err != nil {
				slog.Warn(err.Error())
			} else {
				if result := db.Create(&landmark); result.Error != nil {
					slog.Warn(result.Error.Error())
				}
			}
		}
		if err := scanner.Err(); err != nil {
			slog.Warn(err.Error())
		}
	}
	return nil
}
