-- =============================================
-- Migration: Inisialisasi schema database
-- Deskripsi: Membuat tabel-tabel utama untuk portofolio
-- =============================================

-- Tabel pengalaman kerja
CREATE TABLE IF NOT EXISTS experiences (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    company TEXT NOT NULL,           -- Nama perusahaan
    role TEXT NOT NULL,              -- Posisi/jabatan
    period TEXT NOT NULL,            -- Periode kerja (misal: "Jan 2022 - Mar 2024")
    description TEXT NOT NULL,       -- Deskripsi narasi pengalaman
    sort_order INTEGER DEFAULT 0,    -- Urutan tampil (kecil = paling atas)
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Tabel proyek portofolio
CREATE TABLE IF NOT EXISTS projects (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,              -- Judul proyek
    description TEXT NOT NULL,        -- Deskripsi proyek (konteks bisnis & impact)
    tech_used TEXT NOT NULL,          -- Teknologi yang dipakai (comma-separated)
    link TEXT DEFAULT '',             -- Link ke demo/repo (opsional)
    github_url TEXT DEFAULT '',       -- Link ke repository GitHub (opsional)
    image_url TEXT DEFAULT '',        -- URL gambar proyek (opsional)
    sort_order INTEGER DEFAULT 0,     -- Urutan tampil
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Tabel tech stack
CREATE TABLE IF NOT EXISTS tech_stacks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    category TEXT NOT NULL,           -- Kategori (misal: "Backend", "Frontend", "DevOps")
    name TEXT NOT NULL,               -- Nama teknologi
    description TEXT NOT NULL,        -- Deskripsi konteks penggunaan (bukan level skill)
    sort_order INTEGER DEFAULT 0,     -- Urutan tampil
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Tabel pesan kontak dari pengunjung
CREATE TABLE IF NOT EXISTS contact_messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,               -- Nama pengirim
    email TEXT NOT NULL,              -- Email pengirim
    message TEXT NOT NULL,            -- Isi pesan
    is_read INTEGER DEFAULT 0,       -- Status sudah dibaca (0=belum, 1=sudah)
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Tabel konfigurasi situs (key-value)
CREATE TABLE IF NOT EXISTS site_config (
    key TEXT PRIMARY KEY,             -- Kunci konfigurasi (misal: "name", "tagline")
    value TEXT NOT NULL,              -- Nilai konfigurasi
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- =============================================
-- Seed data default untuk konfigurasi situs
-- =============================================
INSERT OR IGNORE INTO site_config (key, value) VALUES
    ('name', 'Habiburramdhan Lesmana'),
    ('tagline', 'Fullstack Developer — Menulis kode, merangkai solusi'),
    ('about', 'Halo! Saya adalah seorang developer yang percaya bahwa kode yang baik adalah kode yang bercerita. Setiap baris yang saya tulis punya tujuan, setiap fungsi punya makna. Saya suka membangun hal-hal yang berguna dan membuat teknologi jadi lebih mudah dipahami.'),
    ('email', 'abiplesmana@gmail.com'),
    ('github', 'https://github.com/apipcode'),
    ('linkedin', 'https://www.linkedin.com/in/ramdhan-lesmana/'),
    ('photo_url', '');

-- Seed data contoh pengalaman kerja
INSERT OR IGNORE INTO experiences (id, company, role, period, description, sort_order) VALUES
    (1, 'PT Teknologi Nusantara', 'Senior Backend Developer', 'Jan 2023 - Sekarang',
     'Memimpin pengembangan arsitektur microservices untuk platform e-commerce yang melayani jutaan pengguna. Berhasil menurunkan response time API hingga 40% melalui optimasi query dan implementasi caching layer. Sehari-hari bekerja dengan Go, PostgreSQL, dan Redis dalam ekosistem Kubernetes.',
     1),
    (2, 'Startup Kreatif Indonesia', 'Fullstack Developer', 'Mar 2021 - Des 2022',
     'Membangun dashboard analytics dari nol untuk internal tim marketing. Mengintegrasikan data dari berbagai sumber (Google Analytics, database internal, third-party API) menjadi satu tampilan yang mudah dibaca. Stack yang digunakan: React untuk frontend, Node.js untuk backend, dan MongoDB untuk penyimpanan data.',
     2),
    (3, 'Freelance', 'Web Developer', 'Jun 2019 - Feb 2021',
     'Mengerjakan berbagai proyek klien mulai dari company profile, sistem inventory, hingga aplikasi reservasi restoran. Pengalaman ini mengajarkan saya cara berkomunikasi langsung dengan klien, memahami kebutuhan bisnis, dan mendeliver tepat waktu.',
     3);

-- Seed data contoh proyek
INSERT OR IGNORE INTO projects (id, title, description, tech_used, link, sort_order) VALUES
    (1, 'Platform E-Commerce Microservices',
     'Merancang dan membangun ulang arsitektur monolith menjadi microservices untuk platform e-commerce. Hasilnya: deployment jadi 5x lebih cepat, system downtime turun 90%, dan tim bisa develop fitur secara paralel tanpa saling mengganggu.',
     'Go, gRPC, PostgreSQL, Redis, Kubernetes, Docker',
     'https://github.com/username/ecommerce-ms', 1),
    (2, 'Dashboard Analytics Internal',
     'Dashboard real-time yang menyatukan data dari 5+ sumber berbeda. Tim marketing bisa melihat performa campaign dalam satu layar, tanpa harus buka 5 tab berbeda. Mengurangi waktu reporting dari 2 hari menjadi otomatis.',
     'React, Node.js, MongoDB, D3.js, WebSocket',
     'https://github.com/username/analytics-dash', 2),
    (3, 'Sistem Reservasi Restoran',
     'Aplikasi booking meja restoran dengan fitur estimasi waktu tunggu, notifikasi via WhatsApp, dan integrasi POS. Digunakan oleh 3 cabang restoran dengan total 500+ reservasi per minggu.',
     'Laravel, Vue.js, MySQL, WhatsApp API',
     '', 3);

-- Seed data contoh tech stack
INSERT OR IGNORE INTO tech_stacks (id, category, name, description, sort_order) VALUES
    (1, 'Backend', 'Go (Golang)',
     'Bahasa utama untuk membangun layanan backend yang perlu performa tinggi dan concurrency. Dipakai sehari-hari untuk REST API, microservices, dan CLI tools.',
     1),
    (2, 'Backend', 'Node.js',
     'Digunakan untuk prototyping cepat dan project yang butuh ekosistem npm yang luas. Cocok untuk real-time applications dengan WebSocket.',
     2),
    (3, 'Backend', 'PHP (Laravel)',
     'Framework andalan untuk proyek web tradisional. Eloquent ORM-nya membuat kerja dengan database jadi sangat produktif.',
     3),
    (4, 'Frontend', 'React',
     'Library pilihan untuk membangun UI yang kompleks. Digunakan terutama untuk dashboard dan single-page applications.',
     4),
    (5, 'Frontend', 'Vue.js',
     'Alternatif yang lebih ringan dari React. Saya pilih untuk proyek yang butuh setup cepat dan learning curve rendah bagi tim.',
     5),
    (6, 'Frontend', 'HTML/CSS/JavaScript',
     'Fondasi dari semua yang saya bangun di web. Saya percaya memahami vanilla JS dengan baik lebih penting dari menguasai framework apapun.',
     6),
    (7, 'Database', 'PostgreSQL',
     'Database relasional utama untuk production. Fitur JSONB dan full-text search-nya sangat powerful untuk kebutuhan kompleks.',
     7),
    (8, 'Database', 'MongoDB',
     'Digunakan ketika schema data belum pasti atau butuh fleksibilitas tinggi. Terutama untuk logging dan analytics.',
     8),
    (9, 'DevOps', 'Docker & Kubernetes',
     'Containerisasi adalah keharusan di workflow saya. Docker untuk development environment yang konsisten, Kubernetes untuk orchestration di production.',
     9),
    (10, 'DevOps', 'Git & CI/CD',
     'Version control dan automated deployment pipeline. GitHub Actions adalah go-to saya untuk CI/CD yang straightforward.',
     10);
