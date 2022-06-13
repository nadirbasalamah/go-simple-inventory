package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-simple-inventory/models"
	"github.com/nadirbasalamah/go-simple-inventory/services"
)

func GetAllItems(c *fiber.Ctx) error {
	var items []models.Item = services.GetAllItems()

	return c.JSON(fiber.Map{
		"items": items,
	})
}

func GetItemByID(c *fiber.Ctx) error {
	var itemID string = c.Params("id")

	item, err := services.GetItemByID(itemID)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"item": item,
	})
}

func CreateItem(c *fiber.Ctx) error {
	var itemInput *models.ItemRequest = new(models.ItemRequest)

	if err := c.BodyParser(itemInput); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var createdItem models.Item = services.CreateItem(*itemInput)

	return c.JSON(fiber.Map{
		"item": createdItem,
	})
}

func UpdateItem(c *fiber.Ctx) error {
	var itemInput *models.ItemRequest = new(models.ItemRequest)

	if err := c.BodyParser(itemInput); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var itemID string = c.Params("id")

	updatedItem, err := services.UpdateItem(*itemInput, itemID)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"item": updatedItem,
	})
}

func DeleteItem(c *fiber.Ctx) error {
	var itemID string = c.Params("id")

	var result bool = services.DeleteItem(itemID)

	if result {
		return c.JSON(fiber.Map{
			"message": "item deleted successfully",
		})
	}

	return c.Status(http.StatusNotFound).JSON(fiber.Map{
		"message": "item failed to delete",
	})
}
