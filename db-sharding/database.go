package dbsharding

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/lib/pq"
)

const (
	userDB = "postgres"
	passDB = "mysecretpassword"
	hostDB = "localhost"
	portDB = 5432
	nameDB = "sharding-db"
)

func connectDB() (*gorm.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		userDB, passDB, hostDB, portDB, nameDB))
	if err != nil {
		fmt.Println("connect error: ", err)
		return nil, err
	}

	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	return gorm.Open(postgres.New(postgres.Config{
		Conn: db, // Berikan koneksi `*sql.DB` yang sudah dibuka tadi ke GORM
	}), &gorm.Config{
		Logger:      logger.Default.LogMode(2),
		PrepareStmt: true,
	})
}

func ConnectAndMigratePostgres() (*gorm.DB, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}

	isExists := db.Migrator().HasTable(ProductCategory{})
	if !isExists {
		err = db.AutoMigrate(ProductCategory{})
		if err != nil {
			return nil, err
		}
	}

	isExists = db.Migrator().HasTable(Product{})
	if !isExists {
		err = db.AutoMigrate(Product{})
		if err != nil {
			return nil, err
		}
	}

	isExists = db.Migrator().HasTable(Customer{})
	if !isExists {
		err = db.AutoMigrate(Customer{})
		if err != nil {
			return nil, err
		}
	}

	isExists = db.Migrator().HasTable(&SalesOrder{})
	if !isExists {
		err = db.AutoMigrate(&SalesOrder{})
		if err != nil {
			return nil, err
		}
	}

	isExists = db.Migrator().HasTable(&SalesOrderItem{})
	if !isExists {
		err = db.AutoMigrate(&SalesOrderItem{})
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
