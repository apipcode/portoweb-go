package repository

import (
	"database/sql"
	"fmt"
	"portofolio-go/internal/model"
	"time"
)

// Repository menyediakan akses ke database untuk semua operasi CRUD
// Ini adalah lapisan paling bawah yang berinteraksi langsung dengan SQL
type Repository struct {
	db *sql.DB
}

// NewRepository membuat instance Repository baru dengan koneksi database
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// ============================================
// SITE CONFIG — Konfigurasi Situs
// ============================================

// GetAllConfig mengambil semua konfigurasi situs dari tabel site_config
// Mengembalikan map key-value untuk kemudahan akses
func (r *Repository) GetAllConfig() (map[string]string, error) {
	rows, err := r.db.Query("SELECT key, value FROM site_config")
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil konfigurasi: %w", err)
	}
	defer rows.Close()

	config := make(map[string]string)
	for rows.Next() {
		var key, value string
		if err := rows.Scan(&key, &value); err != nil {
			return nil, fmt.Errorf("gagal scan baris konfigurasi: %w", err)
		}
		config[key] = value
	}
	return config, nil
}

// UpdateConfig memperbarui nilai konfigurasi situs berdasarkan key
func (r *Repository) UpdateConfig(key, value string) error {
	_, err := r.db.Exec(
		"INSERT OR REPLACE INTO site_config (key, value, updated_at) VALUES (?, ?, ?)",
		key, value, time.Now(),
	)
	if err != nil {
		return fmt.Errorf("gagal update konfigurasi %s: %w", key, err)
	}
	return nil
}

// ============================================
// EXPERIENCES — Pengalaman Kerja
// ============================================

// GetAllExperiences mengambil semua pengalaman kerja, diurutkan berdasarkan sort_order
func (r *Repository) GetAllExperiences() ([]model.Experience, error) {
	rows, err := r.db.Query(
		"SELECT id, company, role, period, description, sort_order, created_at, updated_at FROM experiences ORDER BY sort_order ASC",
	)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil experiences: %w", err)
	}
	defer rows.Close()

	var experiences []model.Experience
	for rows.Next() {
		var exp model.Experience
		if err := rows.Scan(&exp.ID, &exp.Company, &exp.Role, &exp.Period, &exp.Description, &exp.SortOrder, &exp.CreatedAt, &exp.UpdatedAt); err != nil {
			return nil, fmt.Errorf("gagal scan experience: %w", err)
		}
		experiences = append(experiences, exp)
	}
	return experiences, nil
}

// GetExperienceByID mengambil satu pengalaman kerja berdasarkan ID
func (r *Repository) GetExperienceByID(id int) (*model.Experience, error) {
	var exp model.Experience
	err := r.db.QueryRow(
		"SELECT id, company, role, period, description, sort_order, created_at, updated_at FROM experiences WHERE id = ?", id,
	).Scan(&exp.ID, &exp.Company, &exp.Role, &exp.Period, &exp.Description, &exp.SortOrder, &exp.CreatedAt, &exp.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil experience ID %d: %w", id, err)
	}
	return &exp, nil
}

// CreateExperience menambahkan pengalaman kerja baru ke database
func (r *Repository) CreateExperience(exp *model.Experience) error {
	result, err := r.db.Exec(
		"INSERT INTO experiences (company, role, period, description, sort_order) VALUES (?, ?, ?, ?, ?)",
		exp.Company, exp.Role, exp.Period, exp.Description, exp.SortOrder,
	)
	if err != nil {
		return fmt.Errorf("gagal membuat experience: %w", err)
	}

	// Ambil ID yang baru dibuat
	id, _ := result.LastInsertId()
	exp.ID = int(id)
	return nil
}

// UpdateExperience memperbarui data pengalaman kerja yang sudah ada
func (r *Repository) UpdateExperience(exp *model.Experience) error {
	_, err := r.db.Exec(
		"UPDATE experiences SET company=?, role=?, period=?, description=?, sort_order=?, updated_at=? WHERE id=?",
		exp.Company, exp.Role, exp.Period, exp.Description, exp.SortOrder, time.Now(), exp.ID,
	)
	if err != nil {
		return fmt.Errorf("gagal update experience ID %d: %w", exp.ID, err)
	}
	return nil
}

// DeleteExperience menghapus pengalaman kerja berdasarkan ID
func (r *Repository) DeleteExperience(id int) error {
	_, err := r.db.Exec("DELETE FROM experiences WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("gagal hapus experience ID %d: %w", id, err)
	}
	return nil
}

// ============================================
// PROJECTS — Proyek Portofolio
// ============================================

// GetAllProjects mengambil semua proyek, diurutkan berdasarkan sort_order
func (r *Repository) GetAllProjects() ([]model.Project, error) {
	rows, err := r.db.Query(
		"SELECT id, title, description, tech_used, link, image_url, sort_order, created_at, updated_at FROM projects ORDER BY sort_order ASC",
	)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil projects: %w", err)
	}
	defer rows.Close()

	var projects []model.Project
	for rows.Next() {
		var proj model.Project
		if err := rows.Scan(&proj.ID, &proj.Title, &proj.Description, &proj.TechUsed, &proj.Link, &proj.ImageURL, &proj.SortOrder, &proj.CreatedAt, &proj.UpdatedAt); err != nil {
			return nil, fmt.Errorf("gagal scan project: %w", err)
		}
		projects = append(projects, proj)
	}
	return projects, nil
}

// GetProjectByID mengambil satu proyek berdasarkan ID
func (r *Repository) GetProjectByID(id int) (*model.Project, error) {
	var proj model.Project
	err := r.db.QueryRow(
		"SELECT id, title, description, tech_used, link, image_url, sort_order, created_at, updated_at FROM projects WHERE id = ?", id,
	).Scan(&proj.ID, &proj.Title, &proj.Description, &proj.TechUsed, &proj.Link, &proj.ImageURL, &proj.SortOrder, &proj.CreatedAt, &proj.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil project ID %d: %w", id, err)
	}
	return &proj, nil
}

// CreateProject menambahkan proyek baru ke database
func (r *Repository) CreateProject(proj *model.Project) error {
	result, err := r.db.Exec(
		"INSERT INTO projects (title, description, tech_used, link, image_url, sort_order) VALUES (?, ?, ?, ?, ?, ?)",
		proj.Title, proj.Description, proj.TechUsed, proj.Link, proj.ImageURL, proj.SortOrder,
	)
	if err != nil {
		return fmt.Errorf("gagal membuat project: %w", err)
	}
	id, _ := result.LastInsertId()
	proj.ID = int(id)
	return nil
}

// UpdateProject memperbarui data proyek yang sudah ada
func (r *Repository) UpdateProject(proj *model.Project) error {
	_, err := r.db.Exec(
		"UPDATE projects SET title=?, description=?, tech_used=?, link=?, image_url=?, sort_order=?, updated_at=? WHERE id=?",
		proj.Title, proj.Description, proj.TechUsed, proj.Link, proj.ImageURL, proj.SortOrder, time.Now(), proj.ID,
	)
	if err != nil {
		return fmt.Errorf("gagal update project ID %d: %w", proj.ID, err)
	}
	return nil
}

// DeleteProject menghapus proyek berdasarkan ID
func (r *Repository) DeleteProject(id int) error {
	_, err := r.db.Exec("DELETE FROM projects WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("gagal hapus project ID %d: %w", id, err)
	}
	return nil
}

// ============================================
// TECH STACKS — Teknologi yang Dikuasai
// ============================================

// GetAllTechStacks mengambil semua tech stack, diurutkan berdasarkan sort_order
func (r *Repository) GetAllTechStacks() ([]model.TechStack, error) {
	rows, err := r.db.Query(
		"SELECT id, category, name, description, sort_order, created_at, updated_at FROM tech_stacks ORDER BY sort_order ASC",
	)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil tech stacks: %w", err)
	}
	defer rows.Close()

	var stacks []model.TechStack
	for rows.Next() {
		var ts model.TechStack
		if err := rows.Scan(&ts.ID, &ts.Category, &ts.Name, &ts.Description, &ts.SortOrder, &ts.CreatedAt, &ts.UpdatedAt); err != nil {
			return nil, fmt.Errorf("gagal scan tech stack: %w", err)
		}
		stacks = append(stacks, ts)
	}
	return stacks, nil
}

// GetTechStackByID mengambil satu tech stack berdasarkan ID
func (r *Repository) GetTechStackByID(id int) (*model.TechStack, error) {
	var ts model.TechStack
	err := r.db.QueryRow(
		"SELECT id, category, name, description, sort_order, created_at, updated_at FROM tech_stacks WHERE id = ?", id,
	).Scan(&ts.ID, &ts.Category, &ts.Name, &ts.Description, &ts.SortOrder, &ts.CreatedAt, &ts.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil tech stack ID %d: %w", id, err)
	}
	return &ts, nil
}

// CreateTechStack menambahkan tech stack baru ke database
func (r *Repository) CreateTechStack(ts *model.TechStack) error {
	result, err := r.db.Exec(
		"INSERT INTO tech_stacks (category, name, description, sort_order) VALUES (?, ?, ?, ?)",
		ts.Category, ts.Name, ts.Description, ts.SortOrder,
	)
	if err != nil {
		return fmt.Errorf("gagal membuat tech stack: %w", err)
	}
	id, _ := result.LastInsertId()
	ts.ID = int(id)
	return nil
}

// UpdateTechStack memperbarui data tech stack yang sudah ada
func (r *Repository) UpdateTechStack(ts *model.TechStack) error {
	_, err := r.db.Exec(
		"UPDATE tech_stacks SET category=?, name=?, description=?, sort_order=?, updated_at=? WHERE id=?",
		ts.Category, ts.Name, ts.Description, ts.SortOrder, time.Now(), ts.ID,
	)
	if err != nil {
		return fmt.Errorf("gagal update tech stack ID %d: %w", ts.ID, err)
	}
	return nil
}

// DeleteTechStack menghapus tech stack berdasarkan ID
func (r *Repository) DeleteTechStack(id int) error {
	_, err := r.db.Exec("DELETE FROM tech_stacks WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("gagal hapus tech stack ID %d: %w", id, err)
	}
	return nil
}

// ============================================
// CONTACT MESSAGES — Pesan Kontak
// ============================================

// GetAllContactMessages mengambil semua pesan kontak, terbaru duluan
func (r *Repository) GetAllContactMessages() ([]model.ContactMessage, error) {
	rows, err := r.db.Query(
		"SELECT id, name, email, message, is_read, created_at FROM contact_messages ORDER BY created_at DESC",
	)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil contact messages: %w", err)
	}
	defer rows.Close()

	var messages []model.ContactMessage
	for rows.Next() {
		var msg model.ContactMessage
		if err := rows.Scan(&msg.ID, &msg.Name, &msg.Email, &msg.Message, &msg.IsRead, &msg.CreatedAt); err != nil {
			return nil, fmt.Errorf("gagal scan contact message: %w", err)
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

// CreateContactMessage menyimpan pesan kontak baru dari pengunjung
func (r *Repository) CreateContactMessage(msg *model.ContactMessage) error {
	result, err := r.db.Exec(
		"INSERT INTO contact_messages (name, email, message) VALUES (?, ?, ?)",
		msg.Name, msg.Email, msg.Message,
	)
	if err != nil {
		return fmt.Errorf("gagal menyimpan pesan kontak: %w", err)
	}
	id, _ := result.LastInsertId()
	msg.ID = int(id)
	return nil
}

// MarkMessageAsRead menandai pesan kontak sebagai sudah dibaca
func (r *Repository) MarkMessageAsRead(id int) error {
	_, err := r.db.Exec("UPDATE contact_messages SET is_read = 1 WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("gagal menandai pesan ID %d sebagai dibaca: %w", id, err)
	}
	return nil
}

// DeleteContactMessage menghapus pesan kontak berdasarkan ID
func (r *Repository) DeleteContactMessage(id int) error {
	_, err := r.db.Exec("DELETE FROM contact_messages WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("gagal hapus pesan kontak ID %d: %w", id, err)
	}
	return nil
}
