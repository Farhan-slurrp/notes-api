package notes

import (
	"github.com/Farhan-slurrp/go-notes/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title string `json:"title"`
	Content string `json:"content"`
}

func GetNotes(c *fiber.Ctx) error {
	db := database.DBConn
	notes := []Note{}

	db.Find(&notes)

	err := c.JSON(notes)

	return err
}

func GetNote(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var note Note

	db.Find(&note, id)

	err := c.JSON(note)

	return err
}

func AddNote(c *fiber.Ctx) error {
	db := database.DBConn
	note := new(Note)

	if err := c.BodyParser(note); err != nil {
		c.Status(503).Send([]byte("Cannot add note"))
	}

	db.Create(&note)
	err := c.JSON("book added")

	return err
}

func UpdateNote(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	note := Note{}

	db.First(&note, id)

	if note.Title == "" {
		c.Status(500).JSON("Cannot find data with the given ID")
	}

	if err := c.BodyParser(&note); err != nil {
		c.Status(503).JSON("Please provide data")
	}

	db.Save(&note)


	err := c.JSON("Data updated")
	return err
}

func DeleteNote(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	note := Note{}

	db.First(&note, id)

	if note.Title == "" {
		c.Status(500).JSON("Cannot fine the name")
	}

	db.Delete(&note, id)

	err := c.JSON("Data deleted")
	
	return err
}