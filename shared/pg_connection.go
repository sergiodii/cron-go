package shared

import (
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PGDB struct {
	Database *gorm.DB
	logger   *logger
}

func NewPGDB(loggerFileName string) *PGDB {
	db := new(PGDB)
	if len(loggerFileName) >= 1 {
		db.logger = NewLogger(loggerFileName)
	} else {
		db.logger = NewLogger("")
	}
	return db
}

func init() {
	envFileName := GetEnvFileName()
	err := godotenv.Load(envFileName)
	if err != nil {
		fmt.Println(err)
		// logger := NewLogger("")
		// defer logger.CloseFile()
		// logger.Warning("Shared Package dont founded ", envFileName, " file")
	}
}

func (db *PGDB) makeDsnString() string {
	dbHost := GetEnvOrFail("DB_HOST")
	dbPort := GetEnvOrFail("DB_PORT")
	dbUser := GetEnvOrFail("DB_USER")
	dbPass := GetEnvOrFail("DB_PASS")
	dbDatabase := GetEnvOrFail("DB_DATABASE")
	// url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
	// 	dbUser,
	// 	dbPass,
	// 	dbHost,
	// 	dbPort,
	// 	dbDatabase)
	// return url
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", dbHost, dbUser, dbPass, dbDatabase, dbPort)
}

func (db *PGDB) Connect() {
	dsn := db.makeDsnString()
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		db.logger.Fatal("ERROR ON OPEN DB: ", err)
	}
	db.Database = database
}

func (db *PGDB) Disconnect() {
	datab, err := db.Database.DB()
	if err != nil {
		db.logger.Fatal("ERROR ON CLOSE DB: ", err)
	}
	datab.Close()
}

func (db *PGDB) StartMigration(models ...interface{}) {
	db.Database.AutoMigrate(models...)
}
