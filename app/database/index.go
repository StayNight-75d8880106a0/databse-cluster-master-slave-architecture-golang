package database

import (
	"databse-cluster-master-slave-architecture-golang/app/config/db_config"
	"databse-cluster-master-slave-architecture-golang/app/models"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB_Cluster struct {
	Master         *gorm.DB
	SlaveCases     *gorm.DB
	SlaveSuspects  *gorm.DB
	SlaveDetective *gorm.DB
}

var dbCluster *DB_Cluster

func Connect() *DB_Cluster {

	var ErrorConnect error

	dsnMaster := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s client_encoding=UTF8",
		db_config.DB_Config().MASTER_HOST, db_config.DB_Config().DB_USER, db_config.DB_Config().DB_PASSWORD, db_config.DB_Config().DB_NAME,
		db_config.DB_Config().DB_PORT, db_config.DB_Config().DB_SSLMODE, db_config.DB_Config().DB_TIMEZONE,
	)

	dsnSlave1 := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s client_encoding=UTF8",
		db_config.DB_Config().SLAVE1_HOST, db_config.DB_Config().DB_USER, db_config.DB_Config().DB_PASSWORD, db_config.DB_Config().DB_NAME,
		db_config.DB_Config().DB_PORT, db_config.DB_Config().DB_SSLMODE, db_config.DB_Config().DB_TIMEZONE,
	)

	dsnSlave2 := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s client_encoding=UTF8",
		db_config.DB_Config().SLAVE2_HOST, db_config.DB_Config().DB_USER, db_config.DB_Config().DB_PASSWORD, db_config.DB_Config().DB_NAME,
		db_config.DB_Config().DB_PORT, db_config.DB_Config().DB_SSLMODE, db_config.DB_Config().DB_TIMEZONE,
	)

	dsnSlave3 := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s client_encoding=UTF8",
		db_config.DB_Config().SLAVE3_HOST, db_config.DB_Config().DB_USER, db_config.DB_Config().DB_PASSWORD, db_config.DB_Config().DB_NAME,
		db_config.DB_Config().DB_PORT, db_config.DB_Config().DB_SSLMODE, db_config.DB_Config().DB_TIMEZONE,
	)

	master, ErrorConnect := gorm.Open(postgres.Open(dsnMaster), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	slave1, ErrorConnect := gorm.Open(postgres.Open(dsnSlave1), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	slave2, ErrorConnect := gorm.Open(postgres.Open(dsnSlave2), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	slave3, ErrorConnect := gorm.Open(postgres.Open(dsnSlave3), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if ErrorConnect != nil {
		panic("An error occurred while trying to connect to the database!! " + ErrorConnect.Error())
	}

	masterSQLDB, err := master.DB()
	if err != nil {
		panic("failed to get master sql.DB: " + err.Error())
	}
	masterSQLDB.SetMaxOpenConns(100)
	masterSQLDB.SetMaxIdleConns(25)
	masterSQLDB.SetConnMaxLifetime(30 * time.Minute)
	masterSQLDB.SetConnMaxIdleTime(10 * time.Minute)

	slave1SQLDB, err := slave1.DB()
	if err != nil {
		panic("failed to get slave1 sql.DB: " + err.Error())
	}
	slave1SQLDB.SetMaxOpenConns(150)
	slave1SQLDB.SetMaxIdleConns(40)
	slave1SQLDB.SetConnMaxLifetime(30 * time.Minute)
	slave1SQLDB.SetConnMaxIdleTime(10 * time.Minute)

	slave2SQLDB, err := slave2.DB()
	if err != nil {
		panic("failed to get slave2 sql.DB: " + err.Error())
	}
	slave2SQLDB.SetMaxOpenConns(150)
	slave2SQLDB.SetMaxIdleConns(40)
	slave2SQLDB.SetConnMaxLifetime(30 * time.Minute)
	slave2SQLDB.SetConnMaxIdleTime(10 * time.Minute)

	slave3SQLDB, err := slave3.DB()
	if err != nil {
		panic("failed to get slave2 sql.DB: " + err.Error())
	}
	slave3SQLDB.SetMaxOpenConns(150)
	slave3SQLDB.SetMaxIdleConns(40)
	slave3SQLDB.SetConnMaxLifetime(30 * time.Minute)
	slave3SQLDB.SetConnMaxIdleTime(10 * time.Minute)

	master.SetupJoinTable(&models.Cases{}, "Detective", &models.CaseDetective{})

	errMigrate := master.AutoMigrate(&models.Cases{}, &models.Suspects{}, &models.Detective{}, &models.CaseDetective{})

	if errMigrate != nil {
		panic("Failed to Migrate Database!! " + errMigrate.Error())
	}

	fmt.Println("=========================================")
	fmt.Println("🚀 Database Cluster Status:")
	fmt.Println("✅ Master Connection: OK!")
	fmt.Println("✅ Slave 1 Connection: OK!")
	fmt.Println("✅ Slave 2 Connection: OK!")
	fmt.Println("✅ Slave 3 Connection: OK!")
	fmt.Println("✅ Auto Migration: Success!")
	fmt.Println("=========================================")

	dbCluster = &DB_Cluster{
		Master:         master,
		SlaveCases:     slave1,
		SlaveSuspects:  slave2,
		SlaveDetective: slave3,
	}

	return dbCluster
}

func GetInstanceDbCluster() *DB_Cluster {
	if dbCluster == nil {
		panic("Database cluster is not initialized. Please call Connect() first.")
	}
	return dbCluster
}
