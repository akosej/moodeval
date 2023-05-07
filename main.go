package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var (
	db  *gorm.DB
	err error
)

func main() {
	db, err = gorm.Open("sqlite3", "./db/data.db")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.AutoMigrate(Facultades{}, Dpto{}, Profesores{}, Planes{}, Anos{}, Semestres{}, Asignaturas{}, Reporte{})

	gin.SetMode(gin.ReleaseMode)
	r := engine()
	r.Use(gin.Logger())
	if err := engine().Run(":8080"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}
