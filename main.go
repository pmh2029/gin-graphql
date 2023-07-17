package main

import (
	"gin-graphql/app/migration"
	"gin-graphql/app/router"
	"gin-graphql/pkg/database"
	sharedLogger "gin-graphql/pkg/logger"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger := sharedLogger.NewLogger()

	err := godotenv.Load(filepath.Join(".env"))
	if err != nil {
		logger.Fatalln("Failed to load .env.")
		panic(err)
	}

	gin.SetMode(gin.DebugMode)

	dbconfig := database.DBConfig{
		HostMaster:  os.Getenv("DB_HOST"),
		HostSlaver:  os.Getenv("DB_HOST"),
		Name:        os.Getenv("DB_NAME"),
		User:        os.Getenv("DB_USER"),
		Pass:        os.Getenv("DB_PASS"),
		Port:        os.Getenv("DB_PORT"),
		Charset:     "utf8mb4",
		SSLMode:     os.Getenv("DB_SSLMODE"),
		SSLRootCert: os.Getenv("DB_SSLROOTCERT"),
	}

	logger.Info("Init Database")
	dbConn, err := database.NewDB(dbconfig, logger)
	if err != nil {
		logger.Fatalln("Failed to connect database.")
		panic(err)
	}
	logger.Info("Init Database Success")

	defer database.CloseDB(logger, dbConn)

	logger.Info("Migrate Database")
	err = migration.Migrate(dbConn)
	if err != nil {
		logger.Fatalln("Failed to migrate database.")
		panic(err)
	}
	logger.Info("Migrate Database Success")

	engine := gin.Default()
	router := &router.Router{
		Engine: engine,
		DBCon:  dbConn,
	}
	router.InitializeRouter(logger)
	router.SetupHandler()

	engine.Run(":" + os.Getenv("PORT"))
}
