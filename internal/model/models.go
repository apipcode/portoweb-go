package model

import "time"

// Experience merepresentasikan pengalaman kerja
// Data ini ditampilkan di halaman Experience dalam format timeline
type Experience struct {
	ID          int       `json:"id"`
	Company     string    `json:"company"`     // Nama perusahaan
	Role        string    `json:"role"`        // Posisi/jabatan
	Period      string    `json:"period"`      // Periode kerja
	Description string    `json:"description"` // Deskripsi narasi pengalaman
	SortOrder   int       `json:"sort_order"`  // Urutan tampil
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Project merepresentasikan proyek portofolio
// Ditampilkan dengan konteks bisnis dan dampak, bukan sekadar daftar fitur
type Project struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`       // Judul proyek
	Description string    `json:"description"` // Deskripsi dengan konteks bisnis & impact
	TechUsed    string    `json:"tech_used"`   // Teknologi yang digunakan (comma-separated)
	Link        string    `json:"link"`        // Link ke demo/repo
	ImageURL    string    `json:"image_url"`   // URL gambar proyek
	SortOrder   int       `json:"sort_order"`  // Urutan tampil
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TechStack merepresentasikan teknologi yang dikuasai
// Dideskripsikan dalam konteks penggunaan, BUKAN level persentase
type TechStack struct {
	ID          int       `json:"id"`
	Category    string    `json:"category"`    // Kategori (Backend, Frontend, DevOps, dll)
	Name        string    `json:"name"`        // Nama teknologi
	Description string    `json:"description"` // Konteks penggunaan
	SortOrder   int       `json:"sort_order"`  // Urutan tampil
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ContactMessage merepresentasikan pesan dari pengunjung
// melalui form kontak di website
type ContactMessage struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`    // Nama pengirim
	Email     string    `json:"email"`   // Email pengirim
	Message   string    `json:"message"` // Isi pesan
	IsRead    bool      `json:"is_read"` // Status sudah dibaca atau belum
	CreatedAt time.Time `json:"created_at"`
}

// SiteConfig merepresentasikan konfigurasi situs (key-value)
// Digunakan untuk menyimpan data seperti nama, tagline, about, dll
type SiteConfig struct {
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PortfolioData adalah kumpulan semua data yang dibutuhkan
// untuk merender halaman utama portofolio
type PortfolioData struct {
	Config      map[string]string // Konfigurasi situs (key-value)
	Experiences []Experience      // Daftar pengalaman kerja
	Projects    []Project         // Daftar proyek
	TechStacks  []TechStack       // Daftar tech stack per kategori
}

// ContactForm adalah struct untuk validasi input form kontak
type ContactForm struct {
	Name    string `json:"name" form:"name" binding:"required,min=2,max=100"`
	Email   string `json:"email" form:"email" binding:"required,email"`
	Message string `json:"message" form:"message" binding:"required,min=10,max=2000"`
}

// AdminLoginForm adalah struct untuk validasi input login admin
type AdminLoginForm struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
