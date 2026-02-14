package main

import (
	"fmt"
	"html/template"
	"log"
	"strings"

	"portofolio-go/internal/config"
	"portofolio-go/internal/database"
	"portofolio-go/internal/handler"
	"portofolio-go/internal/middleware"
	"portofolio-go/internal/repository"
	"portofolio-go/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Muat file .env jika ada (untuk development)
	// Di production, environment variables diset langsung di sistem
	_ = godotenv.Load()

	// Muat konfigurasi dari environment variables
	cfg := config.LoadConfig()

	// Set mode Gin berdasarkan konfigurasi
	if cfg.AppMode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Inisialisasi database SQLite dan jalankan migration
	db, err := database.InitDB(cfg.DBPath)
	if err != nil {
		log.Fatalf("Gagal menginisialisasi database: %v", err)
	}
	defer db.Close()

	// Inisialisasi layer-layer arsitektur (dependency injection)
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)

	// Inisialisasi handler
	pageHandler := handler.NewPageHandler(svc)
	contactHandler := handler.NewContactHandler(svc)
	adminHandler := handler.NewAdminHandler(svc, cfg)

	// Setup router Gin
	r := gin.Default()

	// Daftarkan custom template functions dan muat template secara manual
	// (LoadHTMLGlob tidak mendukung nested directory dengan baik)
	funcMap := template.FuncMap{
		// split memecah string berdasarkan separator
		// Digunakan di template untuk memecah tech_used yang comma-separated
		"split": func(s, sep string) []string {
			return strings.Split(s, sep)
		},
		// trim menghapus whitespace di awal dan akhir string
		"trim": strings.TrimSpace,
		// add menjumlahkan dua angka (untuk kalkulasi di template)
		"add": func(a, b int) int {
			return a + b
		},
		// safe menandai string sebagai HTML yang aman (tidak di-escape)
		// Hanya gunakan untuk konten yang sudah disanitasi
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	}

	// Muat semua template HTML dari berbagai subdirektori
	tmpl := template.Must(template.New("").Funcs(funcMap).ParseFiles(
		"web/templates/pages/index.html",
		"web/templates/admin/login.html",
		"web/templates/admin/dashboard.html",
	))
	r.SetHTMLTemplate(tmpl)

	// Serve file statis (CSS, JS, gambar)
	r.Static("/static", "./web/static")

	// ============================================
	// ROUTES — Definisi rute aplikasi
	// ============================================

	// Halaman utama portofolio
	r.GET("/", pageHandler.Index)

	// API kontak form
	r.POST("/api/contact", contactHandler.SubmitContact)

	// ============================================
	// Admin routes — dilindungi middleware auth
	// ============================================

	// Login (tidak perlu auth)
	r.GET("/admin/login", adminHandler.ShowLogin)
	r.POST("/admin/login", adminHandler.Login)

	// Admin panel (perlu auth)
	admin := r.Group("/admin")
	admin.Use(middleware.AuthRequired())
	{
		// Dashboard utama
		admin.GET("", adminHandler.Dashboard)

		// Logout
		admin.POST("/logout", adminHandler.Logout)

		// CRUD Experience
		admin.POST("/experience", adminHandler.CreateExperience)
		admin.POST("/experience/:id", adminHandler.UpdateExperience)
		admin.POST("/experience/:id/delete", adminHandler.DeleteExperience)

		// CRUD Projects
		admin.POST("/project", adminHandler.CreateProject)
		admin.POST("/project/:id", adminHandler.UpdateProject)
		admin.POST("/project/:id/delete", adminHandler.DeleteProject)

		// CRUD Tech Stacks
		admin.POST("/techstack", adminHandler.CreateTechStack)
		admin.POST("/techstack/:id", adminHandler.UpdateTechStack)
		admin.POST("/techstack/:id/delete", adminHandler.DeleteTechStack)

		// Update konfigurasi situs
		admin.POST("/config", adminHandler.UpdateSiteConfig)

		// Pesan kontak
		admin.POST("/message/:id/read", adminHandler.MarkMessageRead)
		admin.POST("/message/:id/delete", adminHandler.DeleteMessage)
	}

	// Jalankan server
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("🚀 Server berjalan di http://localhost%s", addr)
	log.Printf("📖 Buku portofolio siap dibaca!")

	if err := r.Run(addr); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
