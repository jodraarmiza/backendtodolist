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

	// ✅ CORS Config (Menambahkan domain Netlify & Railway)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:5173",                                  // Saat development
			"http://192.168.1.36:5173",                               // Saat testing di jaringan lokal
			"https://todolistjotam.netlify.app",                      // Frontend yang sudah dideploy di Netlify
			"https://backendtodolist-production-e715.up.railway.app", // Backend di Railway
		},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	// ✅ Koneksi ke Database
	database.ConnectDB()

	// ✅ Migrasi Database
	if err := database.DB.AutoMigrate(&models.User{}, &models.Todo{}); err != nil {
		log.Fatalf("❌ Gagal melakukan migrasi database: %v", err)
	}

	// ✅ Routes
	e.POST("/register", auth.Register)
	e.POST("/login", auth.Login)
	e.GET("/todos", handlers.GetToDos)
	e.POST("/todos", handlers.CreateToDo)

	// ✅ Jalankan server di semua network (localhost & jaringan)
	serverAddress := "0.0.0.0:8080"
	log.Printf("✅ Server berjalan di: http://%s", serverAddress)
	e.Logger.Fatal(e.Start(serverAddress))
}
