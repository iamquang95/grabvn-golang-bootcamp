package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"time"
)

type TodoItem struct {
	Id          int
	Title       string
	Completed   bool
	CreatedDate time.Time
}

var db *gorm.DB

func main() {

	// Open connection to db
	var err error
	// If you use := here, it will caused an error because db is created before. TIL lol =))
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=quangle dbname=todolist password=password sslmode=disable")
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}
	db.LogMode(true)
	defer db.Close()

	err = db.AutoMigrate(TodoItem{}).Error
	if err != nil {
		log.Fatal("failed to migrate table todolist")
	}

	router := gin.Default()

	router.GET("/todos", listTodos)
	router.POST("/create", createTodo)
	err = router.Run(":8088")
	if err != nil {
		log.Fatal("failed to start server")
	}
}

func listTodos(c *gin.Context) {
	var todos []TodoItem
	err := db.Find(&todos).Error
	if err != nil {
		c.String(500, "failed to get list of todo items")
		return
	}
	c.JSON(200, todos)
}

func createTodo(c *gin.Context) {
	var arg struct {
		Title string
	}
	err := c.BindJSON(&arg)
	if err != nil {
		c.String(400, "invalid paramerter")
		return
	}

	todo := TodoItem{
		Title: arg.Title,
	}

	err = db.Create(&todo).Error
	if err != nil {
		c.String(500, "failed to create new todo")
	}

	c.JSON(200, todo)
}
