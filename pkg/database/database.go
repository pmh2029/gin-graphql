package database

import (
	"fmt"

	sharedLogger "gin-graphql/pkg/logger"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

// DBConfig config for DB
type DBConfig struct {
	HostMaster  string
	HostSlaver  string
	Name        string
	User        string
	Pass        string
	Port        string
	Charset     string
	SSLMode     string
	SSLRootCert string
}

// NewDB initialize database
func NewDB(
	config DBConfig,
	logger *logrus.Logger,
) (
	*gorm.DB,
	error,
) {
	var dbConn *gorm.DB
	var err error

	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=prefer"
	dsnMaster := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s sslrootcert=%s", config.HostMaster, config.User, config.Pass, config.Name, config.Port, config.SSLMode, config.SSLRootCert)
	dsnSlave := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s sslrootcert=%s", config.HostSlaver, config.User, config.Pass, config.Name, config.Port, config.SSLMode, config.SSLRootCert)

	dbConn, err = gorm.Open(postgres.Open(dsnMaster), &gorm.Config{
		Logger: sharedLogger.NewGormLogger(logger),
	})
	if err != nil {
		return nil, err
	}
	resolver := dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{postgres.Open(dsnSlave)},
	})
	err = dbConn.Use(resolver)
	if err != nil {
		return nil, err
	}

	err = Ping(dbConn)
	return dbConn, err
}

func CloseDB(
	logger *logrus.Logger,
	db *gorm.DB,
) {
	myDB, err := db.DB()
	if err != nil {
		logger.Errorf("Error while returning *sql.DB: %v", err)
	}

	logger.Info("Closing the DB connection pool")
	if err := myDB.Close(); err != nil {
		logger.Errorf("Error while closing the master DB connection pool: %v", err)
	}
}

func Ping(db *gorm.DB) error {
	myDB, err := db.DB()
	if err != nil {
		return err
	}

	return myDB.Ping()
}
