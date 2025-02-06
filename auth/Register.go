package auth

import (
	"net/http"

	"github.com/jodraarmiza/backend/database"
	"github.com/jodraarmiza/backend/models"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Register Handler
func Register(c echo.Context) error {
	var user models.User

	// Bind input dari frontend
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	// Cek apakah username sudah digunakan
	var existingUser models.User
	if err := database.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return c.JSON(http.StatusConflict, map[string]string{"message": "Username sudah digunakan"})
	}

	// Hash password sebelum menyimpan ke database
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	// Simpan user baru ke database
	result := database.DB.Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal menyimpan user ke database"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Registrasi berhasil! Silakan login."})
}
