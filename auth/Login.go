package auth

import (
	"net/http"

	"github.com/jodraarmiza/backend/database"
	"github.com/jodraarmiza/backend/models"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Login Handler
func Login(c echo.Context) error {
	var input models.User
	var user models.User

	// Bind input dari frontend
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	// Cari user berdasarkan username
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Username tidak ditemukan"})
	}

	// Verifikasi password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Password salah"})
	}

	// Generate JWT Token
	token, err := GenerateToken(user.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal membuat token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
