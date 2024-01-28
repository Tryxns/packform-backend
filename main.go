package main

import (
	"fmt"
	"packform/Config"
	"packform/Models"
	"packform/Routes"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func main() {
	err = connect()
	if err != nil {
		fmt.Println("Error:", err)
	}
	if err = Config.DB.AutoMigrate(&Models.Orders{}); err != nil {
		fmt.Println("Error:", err)
	}

	router := Routes.SetupRouter()

	router.Run(":9003")
}

func connect() error {
	Config.DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=admin dbname=packform port=5432 sslmode=disable TimeZone=Asia/Jakarta", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,                                                                                                          // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})

	if err != nil {
		return err
	}
	return nil
}
