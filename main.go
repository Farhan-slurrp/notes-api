package main

import (
	"fmt"
	"os"

	"github.com/Farhan-slurrp/go-notes/database"
	"github.com/Farhan-slurrp/go-notes/notes"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() {
	var err error
	database.DBConn, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{}) 

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
        	port = "4000"
    	}

	app.Listen(":"+port)
	fmt.Println("Listening on port", port, "🚀")
}