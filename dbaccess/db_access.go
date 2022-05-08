package dbaccess

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB session for db operation.
var dbSession *gorm.DB

func init() {
	createAndInitDB()
}

func createAndInitDB() {
	db := GetDBSession()

	// read query file to create tables
	result := db.Exec("")
	if result.Error != nil {
		log.Fatal("Error table creation")
	}
	log.Println("DB Tables created")
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
		"localhost", "postgres", "postgres", "userdb", 5432, "disable ", "Asia/Kolkata")

	pgConf := postgres.New(postgres.Config{DSN: dns})
	gormConf := &gorm.Config{}

	db, err := gorm.Open(pgConf, gormConf)
	if err != nil {
		log.Fatal("DB Connection failed")
	}

	dbSession = db
}
