package handler

import (
	"net/http"
	"portofolio-go/internal/service"

	"github.com/gin-gonic/gin"
)

// PageHandler menangani request untuk halaman-halaman utama portofolio
type PageHandler struct {
	svc *service.Service
}

// NewPageHandler membuat instance PageHandler baru
func NewPageHandler(svc *service.Service) *PageHandler {
	return &PageHandler{svc: svc}
}

// Index menampilkan halaman utama portofolio (flip-book)
// Mengambil semua data dari database dan merender template index.html
func (h *PageHandler) Index(c *gin.Context) {
	// Ambil semua data portofolio dari service
	data, err := h.svc.GetPortfolioData()
	if err != nil {
		c.String(http.StatusInternalServerError, "Gagal memuat data portofolio")
		return
	}

	// Kelompokkan tech stacks berdasarkan kategori untuk template
	techByCategory := make(map[string][]map[string]string)
	for _, ts := range data.TechStacks {
		techByCategory[ts.Category] = append(techByCategory[ts.Category], map[string]string{
			"name":        ts.Name,
			"description": ts.Description,
		})
	}

	// Render halaman utama dengan semua data
	c.HTML(http.StatusOK, "index.html", gin.H{
		"config":         data.Config,
		"experiences":    data.Experiences,
		"projects":       data.Projects,
		"techByCategory": techByCategory,
	})
}
