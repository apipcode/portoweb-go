package config

import (
	"os"
)

// AppConfig menyimpan seluruh konfigurasi aplikasi
// yang diambil dari environment variables
type AppConfig struct {
	Port          string // Port server HTTP
	DBPath        string // Path ke file database SQLite
	AdminUsername  string // Username untuk login admin panel
	AdminPassword  string // Password untuk login admin panel
	SessionSecret string // Secret key untuk session cookie
	AppMode       string // Mode aplikasi (development/production)
}

// LoadConfig membaca konfigurasi dari environment variables
// dan mengembalikan struct AppConfig dengan nilai default jika tidak diset
func LoadConfig() *AppConfig {
	return &AppConfig{
		Port:          getEnv("PORT", "8080"),
		DBPath:        getEnv("DB_PATH", "./data/portfolio.db"),
		AdminUsername:  getEnv("ADMIN_USERNAME", "admin"),
		AdminPassword:  getEnv("ADMIN_PASSWORD", "changeme"),
		SessionSecret: getEnv("SESSION_SECRET", "default-secret-ganti-ini"),
		AppMode:       getEnv("APP_MODE", "development"),
	}
}

// getEnv mengambil nilai environment variable
// Jika tidak ditemukan, kembalikan nilai default (fallback)
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
