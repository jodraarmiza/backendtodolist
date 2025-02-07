package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// ✅ Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Warning: .env file tidak ditemukan, menggunakan environment default")
	}

	// ✅ Ambil nilai dari .env
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// ✅ Validasi jika ada environment yang kosong
	if dbHost == "" || dbPort == "" || dbUser == "" || dbName == "" {
		log.Fatal("❌ Environment variables tidak lengkap, pastikan file .env sudah diisi dengan benar")
	}

	// ✅ Format DSN (Data Source Name)
	var dsn string
	if dbPassword == "" {
		dsn = fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable",
			dbHost, dbUser, dbName, dbPort)
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			dbHost, dbUser, dbPassword, dbName, dbPort)
	}

	// ✅ Koneksi ke database
	var errDB error
	DB, errDB = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errDB != nil {
		log.Fatal("❌ Gagal konek ke database:", errDB)
	}

	log.Println("✅ Database connected successfully!")
}
