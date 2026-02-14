# 📓 Portofolio — Notebook Flip-Book

Website portofolio fullstack developer dengan tema buku catatan yang bisa dibolak-balik halamannya. Dibangun dengan Go (Gin), SQLite, dan vanilla HTML/CSS/JS.

![Go](https://img.shields.io/badge/Go-1.25-blue) ![SQLite](https://img.shields.io/badge/SQLite-3-green) ![License](https://img.shields.io/badge/License-MIT-yellow)

## ✨ Fitur

- **Flip-book animation** — CSS 3D transforms tanpa library eksternal
- **7 halaman buku** — Cover, About, Experience, Projects, Tech Stack, Contact, Back Cover
- **Dark mode** — Toggle mode gelap (seperti baca buku di malam hari)
- **Responsive** — Desktop: flip-book, Mobile: scroll vertikal
- **Admin panel** — CRUD konten tanpa edit kode
- **Form kontak** — Validasi frontend & backend
- **Database SQLite** — Simple, single-file, no setup
- **Docker ready** — Deploy dalam hitungan menit

## 🏗 Arsitektur

```
cmd/server/main.go          → Entry point
internal/
├── config/config.go        → Environment config
├── database/database.go    → SQLite init & migration
├── handler/                → HTTP handlers (page, contact, admin)
├── middleware/auth.go      → Session auth
├── model/models.go         → Data structs
├── repository/repository.go → Database queries
└── service/service.go      → Business logic
web/
├── templates/              → HTML templates (Go template)
└── static/                 → CSS, JS, images
```

## 🚀 Quick Start

### Prasyarat

- Go 1.21+ (dengan CGO enabled)
- GCC (untuk compile go-sqlite3)

### Development

```bash
# 1. Clone repo
git clone <repo-url>
cd portofolio-go

# 2. Salin file konfigurasi
cp .env.example .env

# 3. Edit .env sesuai kebutuhan
nano .env

# 4. Install dependencies
go mod download

# 5. Jalankan server
go run ./cmd/server
```

Buka `http://localhost:8080` di browser.

### Docker

```bash
# Build & jalankan dengan Docker Compose
docker compose up --build

# Atau build manual
docker build -t portfolio .
docker run -p 8080:8080 -v ./data:/app/data portfolio
```

## ⚙ Konfigurasi

Salin `.env.example` ke `.env` dan sesuaikan:

| Variable | Default | Deskripsi |
|---|---|---|
| `PORT` | `8080` | Port server HTTP |
| `DB_PATH` | `./data/portfolio.db` | Path file database SQLite |
| `ADMIN_USERNAME` | `admin` | Username admin panel |
| `ADMIN_PASSWORD` | `changeme` | Password admin panel |
| `SESSION_SECRET` | `...` | Secret key untuk session |
| `APP_MODE` | `development` | `development` / `production` |

## 📝 Admin Panel

Akses admin panel di `http://localhost:8080/admin/login`

Fitur:
- Update profil (nama, tagline, about, social links)
- CRUD pengalaman kerja
- CRUD proyek portofolio
- CRUD tech stack
- Baca & hapus pesan kontak

## 📂 Database

Menggunakan SQLite dengan migration otomatis. Schema ada di `migrations/001_init.sql`:

- `site_config` — Konfigurasi situs (key-value)
- `experiences` — Pengalaman kerja
- `projects` — Proyek portofolio
- `tech_stacks` — Teknologi yang dikuasai
- `contact_messages` — Pesan dari pengunjung

## 🎨 Desain

- **Tema**: Buku catatan handmade dengan tekstur kertas
- **Font**: Caveat (handwriting) + Merriweather (body) + Fira Code (code)
- **Animasi**: CSS 3D transforms flip-book
- **Dark mode**: Warna kertas → gelap, tinta → terang
- **Mobile**: Scroll vertikal dengan bottom navigation

## 📄 Lisensi

MIT License — gunakan sesuka hati.
