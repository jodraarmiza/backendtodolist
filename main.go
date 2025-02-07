package main

import (
	"log"

	"github.com/jodraarmiza/backend/auth"
	"github.com/jodraarmiza/backend/database"
	"github.com/jodraarmiza/backend/handlers"
	"github.com/jodraarmiza/backend/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// ✅ Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ✅ CORS Config agar frontend bisa akses backend
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173", "http://192.168.1.36:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// ✅ Koneksi ke Database
	database.ConnectDB()

	// ✅ Migrasi Database
	if err := database.DB.AutoMigrate(&models.User{}, &models.Todo{}); err != nil {
		log.Fatal("Gagal melakukan migrasi database:", err)
	}

	// ✅ Routes
	e.POST("/register", auth.Register) // Pastikan ini ada
	e.POST("/login", auth.Login)
	e.GET("/todos", handlers.GetToDos)
	e.POST("/todos", handlers.CreateToDo)

	// ✅ Jalankan server
	log.Println("✅ Server berjalan di: http://192.168.1.36:8080")
	e.Logger.Fatal(e.Start("0.0.0.0:8080")) // Bisa diakses dari jaringan
}
