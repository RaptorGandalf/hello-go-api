package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/RaptorGandalf/hello-go-api/repository"
	"github.com/RaptorGandalf/hello-go-api/router"

	"github.com/golang-migrate/migrate"

	// for postgres migrations
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	var connection *gorm.DB

	for {
		con, err := repository.GetConnection()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Retrying in 10 seconds.")
			time.Sleep(time.Second * 10)
		} else {
			fmt.Println("Database connection established")
			connection = con
			break
		}
	}

	files, err := ioutil.ReadDir("db/migrations")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(files)

	migration, err := migrate.New("file://db/migrations/", "postgres://postgres:postgres@localhost:5432/sample?sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println(err)
		return
	}

	http := gin.Default()
	http.Use(cors.Default())

	router.Setup(http, connection)

	err = http.Run("0.0.0.0:80")
	if err != nil {
		fmt.Println(err)
	}
}
