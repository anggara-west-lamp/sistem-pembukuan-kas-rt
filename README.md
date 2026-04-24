# Sistem Pembukuan Kas RT

Backend API (Golang + Gin + GORM + PostgreSQL) untuk mengelola kas RT: autentikasi JWT, pengguna/role, kas, transaksi, dan laporan bulanan. Penyimpanan bukti pembayaran disiapkan untuk MinIO/S3.

## Fitur
- JWT Auth (login)
- CRUD User & Role (sederhana)
- CRUD Kas (master data kas/kategori)
- Transaksi (kas masuk/keluar) dengan dukungan URL bukti pembayaran
- Laporan bulanan (total pemasukan, pengeluaran, saldo)

## Quick Start (Docker Compose)
1. Salin .env.example menjadi .env dan sesuaikan bila perlu
2. Jalankan: `docker compose up -d --build`
3. API di `http://localhost:8080`

## Tanpa Docker
1. Siapkan PostgreSQL dan buat DB sesuai .env
2. `go run ./cmd/server`

## Endpoints (awal)
- POST /api/v1/auth/login { email, password }
- GET /api/v1/healthz
- CRUD /api/v1/users (Bearer token)
- CRUD /api/v1/kas (Bearer token)
- POST /api/v1/transaksi (Bearer token)
- GET /api/v1/laporan?month=2026-04 (Bearer token)

Catatan: Skema diperluas dari model C4 yang disediakan. Struktur repo mengikuti komponen: auth, user, kas, transaksi, laporan dengan layer repository (GORM).

