package main

import (
	"log"

	"github.com/jodraarmiza/backend/auth"
	"github.com/jodraarmiza/backend/database"
	"github.com/jodraarmiza/backend/handlers" // ✅ Pastikan handler digunakan
	"github.com/jodraarmiza/backend/models"   // ✅ Gunakan model untuk migrasi
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// ✅ Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ✅ CORS agar bisa diakses dari jaringan (localhost & network 192.168.1.36)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", "https://192.168.1.36:5173"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// ✅ Koneksi ke Database
	database.ConnectDB()

	// ✅ Migrasi Database
	if err := database.DB.AutoMigrate(&models.User{}, &models.Todo{}); err != nil {
		log.Fatal("Gagal melakukan migrasi database:", err)
	}

	// ✅ Routes
	e.POST("/register", auth.Register)
	e.POST("/login", auth.Login)
	e.GET("/todos", handlers.GetToDos)
	e.POST("/todos", handlers.CreateToDo)

	// ✅ Jalankan server di semua network
	e.Logger.Fatal(e.Start("0.0.0.0:8080")) // Bisa diakses dari jaringan
}
