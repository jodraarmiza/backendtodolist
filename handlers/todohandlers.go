package handlers

import (
	"net/http"

	"github.com/jodraarmiza/backend/database"
	"github.com/jodraarmiza/backend/models"
	"github.com/labstack/echo/v4"
)

// GetToDos ✅ (Ambil semua todo dari database)
func GetToDos(c echo.Context) error {
	var todos []models.Todo
	if err := database.DB.Find(&todos).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengambil data"})
	}
	return c.JSON(http.StatusOK, todos)
}

// CreateToDo ✅ (Tambahkan todo baru)
func CreateToDo(c echo.Context) error {
	var todo models.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Input tidak valid"})
	}

	if err := database.DB.Create(&todo).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal menyimpan data"})
	}

	return c.JSON(http.StatusCreated, todo)
}
