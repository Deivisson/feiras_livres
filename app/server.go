package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Deivisson/free_fairs/db/migrations"
	"github.com/Deivisson/free_fairs/db/seeds"
	"github.com/Deivisson/free_fairs/domain"
	"github.com/Deivisson/free_fairs/service"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Run() {

	checkEnvironmentVariables()

	router := mux.NewRouter()

	dbClient := connectDB()
	fairRepositoryDb := domain.NewFairRepositoryDb(dbClient)
	fh := FairHandlers{service.NewFairService(fairRepositoryDb)}

	// Database initializations
	migrations.Load(dbClient)
	seeds.ImportCsvFile(fairRepositoryDb)

	// define routes
	router.HandleFunc("/fairs", fh.create).Methods(http.MethodPost)
	router.HandleFunc("/fairs/{id}", fh.update).Methods(http.MethodPut)
	router.HandleFunc("/fairs/search", fh.search).Methods(http.MethodPost)
	router.HandleFunc("/fairs/{id}", fh.getById).Methods(http.MethodGet)
	router.HandleFunc("/fairs/{id}", fh.delete).Methods(http.MethodDelete)

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	if IsDebugMode() {
		port = os.Getenv("DEBUG_PORT") // Debug port is defined on launch.json
	}
	log.Print(fmt.Sprintf("Starting server on %s:%s ...", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}

func checkEnvironmentVariables() {
	envProps := []string{
		"DB_DRIVER",
		"DB_USER",
		"DB_PASSWORD",
		"DB_PORT",
		"DB_HOST",
		"DB_NAME",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			log.Fatal(fmt.Sprintf("Variable %s not found. Shut down the app", k))
		}
	}
}

func connectDB() *gorm.DB {
	var err error
	Dbdriver := os.Getenv("DB_DRIVER")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbPort := os.Getenv("DB_PORT")
	DbHost := os.Getenv("DB_HOST")
	DbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := openDB(connectionString)
	if err != nil {
		setupDB(DbUser, DbPassword, DbPort, DbHost, DbName)
		db, err = openDB(connectionString)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		}
	}
	return db
}

func openDB(connectionString string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(3 * time.Minute)
	return db, err
}

func setupDB(DbUser, DbPassword, DbPort, DbHost, DbName string) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s application_name=upys_skill", DbHost, DbPort, DbUser, "postgres", DbPassword)
	db, _ := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	db.Exec(fmt.Sprintf("CREATE DATABASE %s;", DbName))
}

func IsDebugMode() bool {
	return os.Getenv("DEBUG_PORT") != ""
}
