package dbaccess

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/akhilmk/GoRESTAPI/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DB_SCHEMA_FILE = "config/db_schema.sql"

// DB session for db operation.
var dbSession *gorm.DB

func init() {
	connectAndInitDB()
}

func connectAndInitDB() {
	db := GetDBSession()

	if config.AppConfig.DbInt {
		fileByte, err := ioutil.ReadFile(DB_SCHEMA_FILE)
		if err != nil {
			log.Fatal("Error schema file reading")
		}

		result := db.Exec(string(fileByte))
		if result.Error != nil {
			log.Fatal("Error table creation")
		}
		log.Println("DB Tables created")
	}
}

// Injecting db session for Unit testing purpose
func SetDBSession(dbSesn *gorm.DB) {
	dbSession = dbSesn
}

func GetDBSession() *gorm.DB {
	if dbSession == nil {
		connectToDBSession()
	}
	return dbSession
}

func connectToDBSession() {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.AppConfig.DbHost, config.AppConfig.DbUser, "postgres", config.AppConfig.DbName, 5432, "disable ", "Asia/Kolkata")

	pgConf := postgres.New(postgres.Config{DSN: dns})
	gormConf := &gorm.Config{}

	db, err := gorm.Open(pgConf, gormConf)
	if err != nil {
		log.Fatal("DB Connection failed")
	}

	dbSession = db
}
