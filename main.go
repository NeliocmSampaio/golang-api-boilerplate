package main

import (
	"database/sql"
	"fmt"
	"gin-framework-test/basic-api/controllers"
	"gin-framework-test/basic-api/infrastructure"
	mysqlRepository "gin-framework-test/basic-api/infrastructure/db/mysql"
	"gin-framework-test/basic-api/router"
	"gin-framework-test/basic-api/services"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {

	cfg, err := infrastructure.NewConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	mysqlConfig := mysql.Config{
		User:   cfg.DB.User,
		Passwd: cfg.DB.Password,
		Net:    cfg.DB.Net,
		Addr:   fmt.Sprintf("%s:%d", cfg.DB.Host, cfg.DB.Port),
		DBName: cfg.DB.DBName,
	}

	db, err := sql.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	// MySQL repository
	mysqlConn := mysqlRepository.NewBookRepository(db)

	// TODO: Automate Dependency Injection
	// services
	bookService := services.NewBookService(mysqlConn)

	bookController := controllers.NewBookController(bookService)
	healthController := controllers.NewHealthController()

	r := router.NewRouter(bookController, healthController)
	r.SetupRouter()
	r.Run()
}
