package main

import (
	"fmt"
	"os"
	"github.com/Farhan-slurrp/go-notes/database"
	"github.com/Farhan-slurrp/go-notes/notes"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{}) 

	if err != nil {
		panic("Database error")
	}
	fmt.Println("Connected to db📦")

	database.DBConn.AutoMigrate(&notes.Note{})
	fmt.Println("DB Migrated🚌")
}

func notesMiddleware(app *fiber.App) {	
	app.Get("/", notes.GetNotes)
	app.Get("/:id", notes.GetNote)
	app.Post("/", notes.AddNote)
	app.Delete("/:id", notes.DeleteNote)
	app.Put("/:id", notes.UpdateNote)

}

func main() {
	app := fiber.New()

	initDB()
	mydb, _ := database.DBConn.DB()
	defer mydb.Close()

	notesMiddleware(app)

	port, ok := os.LookupEnv("PORT")

    	if ok == false {
        	port = "3000"
    	}

	app.Listen(":"+port)
	fmt.Println("Listening on port 4000🚀")
}