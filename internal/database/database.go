package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	// Driver SQLite — menggunakan CGO
	_ "github.com/mattn/go-sqlite3"
)

// InitDB menginisialisasi koneksi database SQLite
// Fungsi ini membuat file database jika belum ada,
// lalu menjalankan migration script untuk membuat tabel-tabel
func InitDB(dbPath string) (*sql.DB, error) {
	// Pastikan direktori untuk database ada
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("gagal membuat direktori database: %w", err)
	}

	// Buka koneksi ke SQLite
	db, err := sql.Open("sqlite3", dbPath+"?_journal_mode=WAL&_busy_timeout=5000")
	if err != nil {
		return nil, fmt.Errorf("gagal membuka database: %w", err)
	}

	// Tes koneksi
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("gagal ping database: %w", err)
	}

	// Aktifkan foreign keys di SQLite
	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return nil, fmt.Errorf("gagal mengaktifkan foreign keys: %w", err)
	}

	// Jalankan migration
	if err := runMigrations(db); err != nil {
		return nil, fmt.Errorf("gagal menjalankan migration: %w", err)
	}

	return db, nil
}

// runMigrations membaca dan mengeksekusi file migration SQL
// File migration disimpan di folder migrations/
func runMigrations(db *sql.DB) error {
	// Baca file migration
	migrationFile := "migrations/001_init.sql"
	sqlBytes, err := os.ReadFile(migrationFile)
	if err != nil {
		return fmt.Errorf("gagal membaca file migration %s: %w", migrationFile, err)
	}

	// Eksekusi SQL migration
	if _, err := db.Exec(string(sqlBytes)); err != nil {
		return fmt.Errorf("gagal mengeksekusi migration: %w", err)
	}

	return nil
}
