# GoLearn

Platform belajar Go berbahasa Indonesia yang dirancang agar kamu cepat paham dan siap bekerja sebagai backend engineer. Menggabungkan materi bite-sized, editor kode interaktif, latihan dengan test otomatis, kuis, spaced repetition, dan proyek portofolio.

Repo ini mengimplementasikan MVP pada [PRD-GoLearn.md](./PRD-GoLearn.md).

## Tech Stack

| Layer | Teknologi |
|---|---|
| Frontend | React 18, Vite, Tailwind CSS, TanStack Query, Zustand, Monaco Editor, react-markdown |
| Backend | Go, Gin, sqlx, golang-migrate, JWT (access + refresh), bcrypt |
| Database | MySQL 8 |
| Runner kode | Go Playground API (default) atau Docker sandbox (`RUNNER_MODE=docker`) |

## Fitur yang Sudah Ada

- **Autentikasi** — registrasi, login, refresh token, logout, hapus akun (JWT + bcrypt).
- **Learning path & roadmap** — status pelajaran `terkunci / aktif / selesai`, progres per path.
- **Halaman pelajaran** — Markdown + syntax highlight, panel materi & praktik dua kolom (tab di mobile).
- **Code playground** — Monaco Editor + tombol Run, menampilkan output & error.
- **Latihan koding** — test case otomatis, hint bertingkat, solusi tampil setelah lulus, XP.
- **Kuis** — pilihan ganda & jawaban singkat, penjelasan instan.
- **Progress tracking** — persentase per path, pelajaran selesai, waktu belajar.
- **Dashboard** — lanjutkan belajar, streak, XP/level, review hari ini.
- **Spaced repetition** — flashcard dengan algoritma SM-2 dan review harian.
- **Gamifikasi ringan** — XP, level, streak harian, badge.
- **Proyek terpandu** — langkah proyek dengan checklist progres.
- **Cheatsheet & pencarian** — ringkasan sintaks dan pencarian materi.
- **Catatan & bookmark** per pelajaran.
- **Dark/Light mode**.

## Struktur Repo

```
GoLearn/
├── backend/
│   ├── cmd/api/main.go          # entry point
│   ├── internal/
│   │   ├── config/              # env config
│   │   ├── database/            # koneksi + migrasi otomatis
│   │   ├── handler/             # HTTP handler (Gin)
│   │   ├── middleware/          # auth, CORS, rate limit
│   │   ├── model/               # struct data
│   │   ├── repository → store/  # akses SQL (sqlx)
│   │   ├── runner/              # eksekusi kode (playground/docker)
│   │   ├── seed/                # seed konten dari Markdown/YAML
│   │   ├── service/             # logika bisnis
│   │   └── router/              # definisi route
│   ├── migrations/              # SQL migration (embedded)
│   ├── Dockerfile
│   └── .env.example
├── frontend/
│   ├── src/
│   │   ├── api/                 # fetch client + endpoint
│   │   ├── components/          # UI reusable
│   │   ├── pages/               # halaman
│   │   ├── store/               # Zustand (auth, theme)
│   │   └── ...
│   ├── Dockerfile
│   ├── nginx.conf
│   └── .env.example
├── docker-compose.yml
└── PRD-GoLearn.md
```

## Menjalankan Secara Lokal

### Prasyarat
- Go 1.26+
- Node.js 20+
- MySQL 8 (atau Docker)

### 1. Database

Dengan Docker:

```bash
docker run --name golearn-mysql -e MYSQL_ROOT_PASSWORD=password \
  -e MYSQL_DATABASE=golearn -p 3306:3306 -d mysql:8.4
```

Atau pakai MySQL lokal dan buat database `golearn`.

### 2. Backend

```bash
cd backend
cp .env.example .env       # sesuaikan DB_DSN & JWT_SECRET
go run ./cmd/api
```

Saat start, backend otomatis:
- menjalankan migrasi database (`golang-migrate`),
- melakukan seeding path, pelajaran, latihan, kuis, flashcard, proyek, dan badge.

API berjalan di `http://localhost:8080` (health check: `GET /healthz`).

### 3. Frontend

```bash
cd frontend
npm install
npm run dev
```

Buka `http://localhost:5173`. Vite mem-proxy `/api` ke backend, jadi tidak ada masalah CORS saat dev.

## Menjalankan dengan Docker Compose

```bash
JWT_SECRET="ganti-dengan-secret-panjang" docker compose up --build
```

- Frontend: `http://localhost:5173`
- API: `http://localhost:8080`
- MySQL: `localhost:3306`

## Variabel Environment (Backend)

| Variabel | Default | Keterangan |
|---|---|---|
| `PORT` | `8080` | Port API |
| `DB_DSN` | `root:password@tcp(127.0.0.1:3306)/golearn?...` | DSN MySQL |
| `JWT_SECRET` | dev secret | **Wajib diubah di production** |
| `ACCESS_TTL` | `15m` | Masa berlaku access token |
| `REFRESH_TTL` | `168h` | Masa berlaku refresh token |
| `CORS_ORIGINS` | `http://localhost:5173` | Origin yang diizinkan (koma) |
| `RUNNER_MODE` | `playground` | `playground` atau `docker` |
| `RUN_TIMEOUT` | `8s` | Batas waktu eksekusi kode |
| `RATE_LIMIT_PER_MIN` | `120` | Rate limit umum |
| `PLAYGROUND_RATE_PER_MIN` | `20` | Rate limit run/submit |

### Mode runner

- **`playground`** (default): meneruskan kode ke Go Playground API. Cepat dan aman, tetapi **tidak mendukung stdin**, sehingga latihan dengan input akan menampilkan pesan untuk memakai Docker.
- **`docker`**: mengeksekusi kode di container terisolasi (`--network none`, batas CPU/memori, read-only). Produksi sebaiknya memakai mode ini. Butuh Docker di host API.

## Endpoint API

Base URL: `/api/v1`. Semua respons memakai format:

```json
{ "success": true, "data": {}, "error": null, "meta": null }
```

**Auth** `POST /auth/register`, `POST /auth/login`, `POST /auth/refresh`, `POST /auth/logout`, `GET /auth/me`, `DELETE /auth/me`

**Konten** `GET /paths`, `GET /paths/:slug`, `GET /lessons/:slug`, `GET /lessons/:slug/exercises`, `GET /lessons/:slug/quizzes`, `GET /search?q=`

**Interaksi** `POST /playground/run`, `POST /exercises/:id/submit`, `POST /quizzes/:id/answer`, `POST /lessons/:slug/complete`

**Progres & review** `GET /me/progress`, `GET /me/dashboard`, `GET /me/review/today`, `POST /me/review/:flashcardId`, `GET /me/stats`

**Catatan & bookmark** `GET/POST /me/notes`, `PUT/DELETE /me/notes/:id`, `GET/POST /me/bookmarks`, `DELETE /me/bookmarks/:lessonId`

**Proyek** `GET /projects`, `GET /projects/:slug`, `POST /projects/steps/:id/complete`

## Menambah Konten

Konten pelajaran ditulis sebagai Markdown dengan frontmatter YAML di `backend/internal/seed/content/`. Format mengikuti lampiran PRD:

```markdown
---
slug: variabel-dan-tipe-data
title: Variabel dan Tipe Data
summary: ...
minutes: 10
exercise:
  title: ...
  starter: |
    package main
  solution: |
    ...
  hints: [...]
  tests:
    - expected: |
        output
quizzes:
  - question: ...
    explanation: ...
    options:
      - text: ...
        correct: true
flashcards:
  - front: ...
    back: ...
---
# Isi markdown pelajaran
```

Konten di-seed ulang otomatis setiap API start. Daftar path/modul didefinisikan di `backend/internal/seed/path1.go`, proyek di `content/projects.yaml`.

## Testing

```bash
cd backend
go test ./...
```

Frontend:

```bash
cd frontend
npm run build
```

## Catatan Keamanan

- Password di-hash dengan bcrypt; refresh token disimpan sebagai hash SHA-256 dan bisa dicabut.
- Validasi input di layer service; query SQL memakai parameter terikat.
- Rate limit in-memory per user/IP (cukup untuk satu instance; ganti ke Redis bila horizontal scaling).
- Untuk `RUNNER_MODE=docker`, pastikan batas resource dan `--network none` tetap aktif.

## Lisensi

Proyek pembelajaran. Silakan gunakan dan modifikasi sesuai kebutuhan.
