package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type Post struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Mail   string `json:"mail"`
	AreaID int    `json:"area_id"`
}

func main() {
	// Loading .env file for confing database.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database.
	cfg := mysql.Config{
		User:      os.Getenv("MYSQL_USER"),
		Passwd:    os.Getenv("MYSQL_PASSWORD"),
		Net:       "tcp",
		Addr:      os.Getenv("MYSQL_HOST"),
		DBName:    os.Getenv("MYSQL_DATABASE"),
		ParseTime: true,
	}
	db, err := sqlx.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Server setting.
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Hello work!"})
	})

	router.GET("/customer", func(ctx *gin.Context) {
		var id string
		var name string
		err := db.QueryRow("SELECT id, name from customer order by id desc limit 1").Scan(&id, &name)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusOK, gin.H{"id": id, "name": name})
	})

	router.GET("/greeting", func(ctx *gin.Context) {
		name := ctx.Query("name")
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Hello " + name})
	})

	router.POST("/", func(ctx *gin.Context) {
		var post Post
		if err := ctx.Bind(&post); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		sql := fmt.Sprintf("insert into customer values (%d, '%s', '%s', %d)", post.ID, post.Name, post.Mail, post.AreaID)
		_, err := db.Exec(sql)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "ok"})
	})

	// Run!!
	router.Run(":9000")
}