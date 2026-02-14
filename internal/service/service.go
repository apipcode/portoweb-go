package service

import (
	"fmt"
	"html"
	"portofolio-go/internal/model"
	"portofolio-go/internal/repository"
	"strings"
)

// Service menyediakan business logic untuk aplikasi
// Layer ini berada di antara handler dan repository
type Service struct {
	repo *repository.Repository
}

// NewService membuat instance Service baru dengan dependency repository
func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

// ============================================
// PORTFOLIO DATA — Data Halaman Utama
// ============================================

// GetPortfolioData mengumpulkan semua data yang dibutuhkan untuk halaman utama
// Menggabungkan config, experiences, projects, dan tech stacks
func (s *Service) GetPortfolioData() (*model.PortfolioData, error) {
	// Ambil konfigurasi situs
	config, err := s.repo.GetAllConfig()
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil config: %w", err)
	}

	// Ambil daftar pengalaman kerja
	experiences, err := s.repo.GetAllExperiences()
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil experiences: %w", err)
	}

	// Ambil daftar proyek
	projects, err := s.repo.GetAllProjects()
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil projects: %w", err)
	}

	// Ambil daftar tech stack
	techStacks, err := s.repo.GetAllTechStacks()
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil tech stacks: %w", err)
	}

	return &model.PortfolioData{
		Config:      config,
		Experiences: experiences,
		Projects:    projects,
		TechStacks:  techStacks,
	}, nil
}

// ============================================
// CONTACT — Pesan Kontak
// ============================================

// SubmitContactMessage memvalidasi dan menyimpan pesan kontak dari pengunjung
// Melakukan sanitasi input untuk mencegah XSS
func (s *Service) SubmitContactMessage(form *model.ContactForm) error {
	// Sanitasi input — bersihkan HTML tags yang berbahaya
	msg := &model.ContactMessage{
		Name:    sanitizeInput(form.Name),
		Email:   sanitizeInput(form.Email),
		Message: sanitizeInput(form.Message),
	}

	// Simpan ke database
	if err := s.repo.CreateContactMessage(msg); err != nil {
		return fmt.Errorf("gagal menyimpan pesan kontak: %w", err)
	}

	return nil
}

// GetAllContactMessages mengambil semua pesan kontak untuk admin
func (s *Service) GetAllContactMessages() ([]model.ContactMessage, error) {
	return s.repo.GetAllContactMessages()
}

// MarkMessageAsRead menandai pesan sebagai sudah dibaca
func (s *Service) MarkMessageAsRead(id int) error {
	return s.repo.MarkMessageAsRead(id)
}

// DeleteContactMessage menghapus pesan kontak
func (s *Service) DeleteContactMessage(id int) error {
	return s.repo.DeleteContactMessage(id)
}

// ============================================
// EXPERIENCES — Pengalaman Kerja (CRUD Admin)
// ============================================

// GetAllExperiences mengambil semua pengalaman kerja
func (s *Service) GetAllExperiences() ([]model.Experience, error) {
	return s.repo.GetAllExperiences()
}

// GetExperienceByID mengambil pengalaman kerja berdasarkan ID
func (s *Service) GetExperienceByID(id int) (*model.Experience, error) {
	return s.repo.GetExperienceByID(id)
}

// CreateExperience membuat pengalaman kerja baru setelah sanitasi
func (s *Service) CreateExperience(exp *model.Experience) error {
	exp.Company = sanitizeInput(exp.Company)
	exp.Role = sanitizeInput(exp.Role)
	exp.Period = sanitizeInput(exp.Period)
	exp.Description = sanitizeInput(exp.Description)
	return s.repo.CreateExperience(exp)
}

// UpdateExperience memperbarui pengalaman kerja setelah sanitasi
func (s *Service) UpdateExperience(exp *model.Experience) error {
	exp.Company = sanitizeInput(exp.Company)
	exp.Role = sanitizeInput(exp.Role)
	exp.Period = sanitizeInput(exp.Period)
	exp.Description = sanitizeInput(exp.Description)
	return s.repo.UpdateExperience(exp)
}

// DeleteExperience menghapus pengalaman kerja
func (s *Service) DeleteExperience(id int) error {
	return s.repo.DeleteExperience(id)
}

// ============================================
// PROJECTS — Proyek (CRUD Admin)
// ============================================

// GetAllProjects mengambil semua proyek
func (s *Service) GetAllProjects() ([]model.Project, error) {
	return s.repo.GetAllProjects()
}

// GetProjectByID mengambil proyek berdasarkan ID
func (s *Service) GetProjectByID(id int) (*model.Project, error) {
	return s.repo.GetProjectByID(id)
}

// CreateProject membuat proyek baru setelah sanitasi
func (s *Service) CreateProject(proj *model.Project) error {
	proj.Title = sanitizeInput(proj.Title)
	proj.Description = sanitizeInput(proj.Description)
	proj.TechUsed = sanitizeInput(proj.TechUsed)
	return s.repo.CreateProject(proj)
}

// UpdateProject memperbarui proyek setelah sanitasi
func (s *Service) UpdateProject(proj *model.Project) error {
	proj.Title = sanitizeInput(proj.Title)
	proj.Description = sanitizeInput(proj.Description)
	proj.TechUsed = sanitizeInput(proj.TechUsed)
	return s.repo.UpdateProject(proj)
}

// DeleteProject menghapus proyek
func (s *Service) DeleteProject(id int) error {
	return s.repo.DeleteProject(id)
}

// ============================================
// TECH STACKS — Teknologi (CRUD Admin)
// ============================================

// GetAllTechStacks mengambil semua tech stack
func (s *Service) GetAllTechStacks() ([]model.TechStack, error) {
	return s.repo.GetAllTechStacks()
}

// GetTechStackByID mengambil tech stack berdasarkan ID
func (s *Service) GetTechStackByID(id int) (*model.TechStack, error) {
	return s.repo.GetTechStackByID(id)
}

// CreateTechStack membuat tech stack baru setelah sanitasi
func (s *Service) CreateTechStack(ts *model.TechStack) error {
	ts.Category = sanitizeInput(ts.Category)
	ts.Name = sanitizeInput(ts.Name)
	ts.Description = sanitizeInput(ts.Description)
	return s.repo.CreateTechStack(ts)
}

// UpdateTechStack memperbarui tech stack setelah sanitasi
func (s *Service) UpdateTechStack(ts *model.TechStack) error {
	ts.Category = sanitizeInput(ts.Category)
	ts.Name = sanitizeInput(ts.Name)
	ts.Description = sanitizeInput(ts.Description)
	return s.repo.UpdateTechStack(ts)
}

// DeleteTechStack menghapus tech stack
func (s *Service) DeleteTechStack(id int) error {
	return s.repo.DeleteTechStack(id)
}

// ============================================
// SITE CONFIG — Konfigurasi (CRUD Admin)
// ============================================

// GetAllConfig mengambil semua konfigurasi situs
func (s *Service) GetAllConfig() (map[string]string, error) {
	return s.repo.GetAllConfig()
}

// UpdateConfig memperbarui konfigurasi situs
func (s *Service) UpdateConfig(key, value string) error {
	return s.repo.UpdateConfig(sanitizeInput(key), sanitizeInput(value))
}

// ============================================
// HELPER FUNCTIONS
// ============================================

// sanitizeInput membersihkan input dari karakter HTML berbahaya
// untuk mencegah serangan XSS (Cross-Site Scripting)
func sanitizeInput(input string) string {
	// Trim whitespace di awal dan akhir
	input = strings.TrimSpace(input)
	// Escape karakter HTML khusus
	input = html.EscapeString(input)
	return input
}
