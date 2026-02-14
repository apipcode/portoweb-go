package handler

import (
	"net/http"
	"portofolio-go/internal/model"
	"portofolio-go/internal/service"

	"github.com/gin-gonic/gin"
)

// ContactHandler menangani request terkait form kontak
type ContactHandler struct {
	svc *service.Service
}

// NewContactHandler membuat instance ContactHandler baru
func NewContactHandler(svc *service.Service) *ContactHandler {
	return &ContactHandler{svc: svc}
}

// SubmitContact menerima dan memproses pesan dari form kontak
// Validasi dilakukan oleh binding Gin, lalu disimpan ke database
func (h *ContactHandler) SubmitContact(c *gin.Context) {
	var form model.ContactForm

	// Validasi input — Gin akan mengecek required, email format, min/max length
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Data tidak valid. Pastikan semua field terisi dengan benar.",
			"error":   err.Error(),
		})
		return
	}

	// Simpan pesan melalui service layer
	if err := h.svc.SubmitContactMessage(&form); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal mengirim pesan. Silakan coba lagi.",
		})
		return
	}

	// Berhasil — kirim response sukses
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Pesan berhasil dikirim! Terima kasih sudah menghubungi.",
	})
}
