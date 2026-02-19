package handler

import (
	"net/http"
	"portofolio-go/internal/config"
	"portofolio-go/internal/middleware"
	"portofolio-go/internal/model"
	"portofolio-go/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AdminHandler menangani semua request untuk admin panel
// Termasuk login, logout, dan CRUD untuk konten portofolio
type AdminHandler struct {
	svc *service.Service
	cfg *config.AppConfig
}

// NewAdminHandler membuat instance AdminHandler baru
func NewAdminHandler(svc *service.Service, cfg *config.AppConfig) *AdminHandler {
	return &AdminHandler{svc: svc, cfg: cfg}
}

// ============================================
// AUTH — Login & Logout
// ============================================

// ShowLogin menampilkan halaman login admin
func (h *AdminHandler) ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

// Login memproses form login admin
// Mengecek username dan password dari environment variable
func (h *AdminHandler) Login(c *gin.Context) {
	var form model.AdminLoginForm

	// Validasi input form login
	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": "Username dan password harus diisi.",
		})
		return
	}

	// Cek credential — bandingkan dengan config dari environment
	if form.Username != h.cfg.AdminUsername || form.Password != h.cfg.AdminPassword {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "Username atau password salah.",
		})
		return
	}

	// Login berhasil — buat session dan redirect ke dashboard
	middleware.CreateSession(c, form.Username)
	c.Redirect(http.StatusFound, "/admin")
}

// Logout menghapus session admin dan redirect ke login
func (h *AdminHandler) Logout(c *gin.Context) {
	middleware.DestroySession(c)
	c.Redirect(http.StatusFound, "/admin/login")
}

// ============================================
// DASHBOARD — Halaman Utama Admin
// ============================================

// Dashboard menampilkan halaman utama admin panel
// Memuat semua data untuk ditampilkan di tabel CRUD
func (h *AdminHandler) Dashboard(c *gin.Context) {
	// Ambil semua data untuk ditampilkan
	experiences, _ := h.svc.GetAllExperiences()
	projects, _ := h.svc.GetAllProjects()
	techStacks, _ := h.svc.GetAllTechStacks()
	messages, _ := h.svc.GetAllContactMessages()
	siteConfig, _ := h.svc.GetAllConfig()

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"experiences": experiences,
		"projects":    projects,
		"techStacks":  techStacks,
		"messages":    messages,
		"siteConfig":  siteConfig,
		"username":    c.GetString("admin_username"),
	})
}

// ============================================
// EXPERIENCE — CRUD Pengalaman Kerja
// ============================================

// CreateExperience menambahkan pengalaman kerja baru via POST
func (h *AdminHandler) CreateExperience(c *gin.Context) {
	sortOrder, _ := strconv.Atoi(c.PostForm("sort_order"))
	exp := &model.Experience{
		Company:     c.PostForm("company"),
		Role:        c.PostForm("role"),
		Period:      c.PostForm("period"),
		Description: c.PostForm("description"),
		SortOrder:   sortOrder,
	}

	if err := h.svc.CreateExperience(exp); err != nil {
		c.Redirect(http.StatusFound, "/admin?error=Gagal+menambah+experience")
		return
	}
	c.Redirect(http.StatusFound, "/admin?success=Experience+berhasil+ditambahkan")
}

// UpdateExperience memperbarui pengalaman kerja via POST
func (h *AdminHandler) UpdateExperience(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sortOrder, _ := strconv.Atoi(c.PostForm("sort_order"))
	exp := &model.Experience{
		ID:          id,
		Company:     c.PostForm("company"),
		Role:        c.PostForm("role"),
		Period:      c.PostForm("period"),
		Description: c.PostForm("description"),
		SortOrder:   sortOrder,
	}

	if err := h.svc.UpdateExperience(exp); err != nil {
		c.Redirect(http.StatusFound, "/admin?error=Gagal+update+experience")
		return
	}
	c.Redirect(http.StatusFound, "/admin?success=Experience+berhasil+diupdate")
}

// DeleteExperience menghapus pengalaman kerja via POST
func (h *AdminHandler) DeleteExperience(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.svc.DeleteExperience(id); err != nil {
		c.Redirect(http.StatusFound, "/admin?error=Gagal+hapus+experience")
		return
	}
	c.Redirect(http.StatusFound, "/admin?success=Experience+berhasil+dihapus")
}

// ============================================
// PROJECTS — CRUD Proyek
// ============================================

// CreateProject menambahkan proyek baru via POST
func (h *AdminHandler) CreateProject(c *gin.Context) {
	sortOrder, _ := strconv.Atoi(c.PostForm("sort_order"))
	proj := &model.Project{
		Title:       c.PostForm("title"),
		Description: c.PostForm("description"),
		TechUsed:    c.PostForm("tech_used"),
		Link:        c.PostForm("link"),
		GithubURL:   c.PostForm("github_url"),
		ImageURL:    c.PostForm("image_url"),
		SortOrder:   sortOrder,
	}

	if err := h.svc.CreateProject(proj); err != nil {
		c.Redirect(http.StatusFound, "/admin?error=Gagal+menambah+project")
		return
	}
	c.Redirect(http.StatusFound, "/admin?success=Project+berhasil+ditambahkan")
}

// UpdateProject memperbarui proyek via POST
func (h *AdminHandler) UpdateProject(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sortOrder, _ := strconv.Atoi(c.PostForm("sort_order"))
	proj := &model.Project{
		ID:          id,
		Title:       c.PostForm("title"),
		Description: c.PostForm("description"),
		TechUsed:    c.PostForm("tech_used"),
		Link:        c.PostForm("link"),
		GithubURL:   c.PostForm("github_url"),
		ImageURL:    c.PostForm("image_url"),
		SortOrder:   sortOrder,
	}

	if err := h.svc.UpdateProject(proj); err != nil {
		c.Redirect(http.StatusFound, "/admin?error=Gagal+update+project")
		return
	}
	c.Redirect(http.StatusFound, "/admin?success=Project+berhasil+diupdate")
}

// DeleteProject menghapus proyek via POST
func (h *AdminHandler) DeleteProject(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.svc.DeleteProject(id); err != nil {
		c.Redirect(http.StatusFound, "/admin?error=Gagal+hapus+project")
		return
	}
	c.Redirect(http.StatusFound, "/admin?success=Project+berhasil+dihapus")
}

// ============================================
// TECH STACKS — CRUD Tech Stack
// ============================================

// CreateTechStack menambahkan tech stack baru via POST
func (h *AdminHandler) CreateTechStack(c *gin.Context) {
	sortOrder, _ := strconv.Atoi(c.PostForm("sort_order"))
	ts := &model.TechStack{
		Category:    c.PostForm("category"),
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
		SortOrder:   sortOrder,
	}

	if err := h.svc.CreateTechStack(ts); err != nil {
		c.Redirect(http.StatusFound, "/admin?error=Gagal+menambah+tech+stack")
		return
	}
	c.Redirect(http.StatusFound, "/admin?success=Tech+stack+berhasil+ditambahkan")
}

// UpdateTechStack memperbarui tech stack via POST
func (h *AdminHandler) UpdateTechStack(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sortOrder, _ := strconv.Atoi(c.PostForm("sort_order"))
	ts := &model.TechStack{
		ID:          id,
		Category:    c.PostForm("category"),
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
		SortOrder:   sortOrder,
	}

	if err := h.svc.UpdateTechStack(ts); err != nil {
		c.Redirect(http.StatusFound, "/admin?error=Gagal+update+tech+stack")
		return
	}
	c.Redirect(http.StatusFound, "/admin?success=Tech+stack+berhasil+diupdate")
}

// DeleteTechStack menghapus tech stack via POST
func (h *AdminHandler) DeleteTechStack(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.svc.DeleteTechStack(id); err != nil {
		c.Redirect(http.StatusFound, "/admin?error=Gagal+hapus+tech+stack")
		return
	}
	c.Redirect(http.StatusFound, "/admin?success=Tech+stack+berhasil+dihapus")
}

// ============================================
// SITE CONFIG — Update Konfigurasi Situs
// ============================================

// UpdateSiteConfig memperbarui konfigurasi situs via POST
func (h *AdminHandler) UpdateSiteConfig(c *gin.Context) {
	// Daftar key yang bisa diupdate
	keys := []string{"name", "tagline", "about", "email", "github", "linkedin", "photo_url"}
	for _, key := range keys {
		value := c.PostForm(key)
		if value != "" {
			if err := h.svc.UpdateConfig(key, value); err != nil {
				c.Redirect(http.StatusFound, "/admin?error=Gagal+update+konfigurasi")
				return
			}
		}
	}
	c.Redirect(http.StatusFound, "/admin?success=Konfigurasi+berhasil+diupdate")
}

// ============================================
// MESSAGES — Pesan Kontak
// ============================================

// MarkMessageRead menandai pesan sebagai sudah dibaca
func (h *AdminHandler) MarkMessageRead(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.svc.MarkMessageAsRead(id)
	c.Redirect(http.StatusFound, "/admin?success=Pesan+ditandai+dibaca")
}

// DeleteMessage menghapus pesan kontak
func (h *AdminHandler) DeleteMessage(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.svc.DeleteContactMessage(id)
	c.Redirect(http.StatusFound, "/admin?success=Pesan+berhasil+dihapus")
}
