package handlers

import (
	"net/http"
	"strconv"

	"github.com/jodraarmiza/backend/database"
	"github.com/jodraarmiza/backend/models"
	"github.com/labstack/echo/v4"
)

// ✅ GetToDos (Ambil semua todo untuk user yang sedang login)
func GetToDos(c echo.Context) error {
	userID, ok := c.Get("user_id").(uint) // Ambil user_id dari JWT
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "User tidak valid"})
	}

	var todos []models.Todo
	if err := database.DB.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengambil data"})
	}
	return c.JSON(http.StatusOK, todos)
}

// ✅ CreateToDo (Tambahkan todo baru untuk user yang sedang login)
func CreateToDo(c echo.Context) error {
	userID, ok := c.Get("user_id").(uint) // Ambil user_id dari JWT
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "User tidak valid"})
	}

	var todo models.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Format input tidak valid"})
	}

	// Pastikan data minimal yang diperlukan ada
	if todo.Text == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Teks tugas tidak boleh kosong"})
	}

	// Set UserID untuk todo baru
	todo.UserID = userID

	// Gunakan transaksi untuk memastikan penyimpanan yang aman
	if err := database.DB.Create(&todo).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal menyimpan tugas"})
	}

	return c.JSON(http.StatusCreated, todo)
}

// ✅ UpdateToDo (Perbarui todo berdasarkan ID, hanya jika milik user yang login)
func UpdateToDo(c echo.Context) error {
	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "User tidak valid"})
	}

	// Konversi ID dari string ke uint
	todoID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID tugas tidak valid"})
	}

	var todo models.Todo
	if err := database.DB.Where("id = ? AND user_id = ?", todoID, userID).First(&todo).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Tugas tidak ditemukan"})
	}

	// Bind input ke struct, tapi hanya memperbarui `Text` dan `Completed`
	var updateData struct {
		Text      string `json:"text"`
		Completed bool   `json:"completed"`
	}
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Format input tidak valid"})
	}

	// Perbarui data yang diizinkan saja
	todo.Text = updateData.Text
	todo.Completed = updateData.Completed

	// Gunakan `Save` untuk menyimpan perubahan
	if err := database.DB.Save(&todo).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal memperbarui tugas"})
	}

	return c.JSON(http.StatusOK, todo)
}

// ✅ DeleteToDo (Hapus todo berdasarkan ID, hanya jika milik user yang login)
func DeleteToDo(c echo.Context) error {
	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "User tidak valid"})
	}

	// Konversi ID dari string ke uint
	todoID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID tugas tidak valid"})
	}

	var todo models.Todo
	if err := database.DB.Where("id = ? AND user_id = ?", todoID, userID).First(&todo).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Tugas tidak ditemukan"})
	}

	// Hapus tugas
	if err := database.DB.Delete(&todo).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus tugas"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Tugas berhasil dihapus"})
}
