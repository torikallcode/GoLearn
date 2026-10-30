# PRD: GoLearn — Platform Belajar Go untuk Menjadi Backend Engineer Profesional

| Item | Detail |
|---|---|
| Versi | 1.0 (Draft) |
| Tanggal | 30 Oktober 2026 |
| Status | Perencanaan |
| Tech Stack | React.js, Vite, Tailwind CSS, Go, Gin, MySQL |

---

## 1. Ringkasan Produk

**GoLearn** adalah website belajar bahasa pemrograman Go yang dirancang agar pengguna **cepat paham**, **belajar dengan cara paling efektif**, dan **siap bekerja sebagai backend engineer profesional**.

Berbeda dengan sekadar kumpulan artikel, GoLearn menggabungkan:
- Materi singkat dan terstruktur (bite-sized)
- Editor kode interaktif langsung di browser
- Latihan dan kuis dengan umpan balik instan
- Pengulangan terjadwal (spaced repetition) agar tidak mudah lupa
- Proyek nyata bertahap yang menyerupai pekerjaan backend engineer

### Visi
Menjadi tempat paling cepat dan efektif bagi seseorang untuk berubah dari nol menjadi backend engineer Go yang siap kerja.

### Masalah yang Diselesaikan
| Masalah | Solusi GoLearn |
|---|---|
| Materi online tersebar dan tidak berurutan | Learning path terstruktur dari dasar sampai profesional |
| Hanya membaca/menonton, tidak praktik | Editor kode interaktif + latihan di setiap materi |
| Cepat lupa setelah belajar | Spaced repetition dan review harian |
| Bingung kapan "sudah siap kerja" | Milestone, skill checklist, dan proyek portofolio |
| Tidak tahu praktik backend nyata | Modul REST API, database, auth, testing, deployment |

---

## 2. Target Pengguna

### Persona Utama
1. **Pemula total** — mahasiswa/fresh graduate yang ingin masuk dunia backend.
2. **Developer berpindah bahasa** — sudah bisa PHP/Node/Python/Java, ingin pindah ke Go.
3. **Career switcher** — dari bidang lain, belajar mandiri menuju karier backend.

### Kebutuhan Utama
- Penjelasan ringkas, jelas, bahasa Indonesia
- Tahu persis "belajar apa selanjutnya"
- Bisa langsung praktik tanpa instalasi rumit
- Merasa progres nyata dan termotivasi

---

## 3. Tujuan & Metrik Keberhasilan

### Tujuan Produk
1. Pengguna memahami konsep inti Go dalam waktu singkat.
2. Pengguna menyelesaikan minimal 1 proyek backend lengkap.
3. Pengguna konsisten belajar (retensi tinggi).

### Metrik (KPI)
| Metrik | Target Awal |
|---|---|
| Completion rate pelajaran pertama | ≥ 70% |
| Retensi D7 | ≥ 30% |
| Rata-rata latihan benar pada percobaan ke-2 | ≥ 80% |
| Pengguna menyelesaikan 1 Learning Path | ≥ 20% |
| Waktu rata-rata menyelesaikan satu pelajaran | 10–15 menit |
| Skor kepuasan (rating 1–5) | ≥ 4.3 |

---

## 4. Prinsip Pembelajaran Efektif (Fondasi Desain)

Setiap fitur harus mengacu pada prinsip berikut:

1. **Learning by doing** — setiap konsep diikuti praktik langsung.
2. **Bite-sized** — satu pelajaran = satu konsep, 5–15 menit.
3. **Active recall** — kuis dan latihan menarik jawaban dari ingatan, bukan hanya membaca.
4. **Spaced repetition** — konsep lama diulang di interval yang makin panjang.
5. **Instant feedback** — hasil kode dan penjelasan error langsung muncul.
6. **Progressive complexity** — dari sederhana ke kompleks, tiap langkah membangun yang sebelumnya.
7. **Project-based** — konsep diterapkan ke proyek nyata, bukan contoh abstrak.
8. **Visual & analogi** — diagram, analogi sehari-hari, perbandingan dengan bahasa lain.
9. **Gamifikasi ringan** — streak, XP, badge untuk menjaga motivasi, tanpa mengganggu fokus.

---

## 5. Kurikulum (Learning Path)

### Path 1 — Fondasi Go (Pemula)
1. Pengenalan Go & instalasi, `go mod`, struktur proyek
2. Variabel, tipe data, konstanta, zero value
3. Operator & control flow (`if`, `switch`, `for`)
4. Fungsi (multiple return, variadic, closure, `defer`)
5. Array, slice, map
6. String, rune, dan package `strings`/`strconv`
7. Struct & method
8. Pointer
9. Interface & polimorfisme
10. Error handling (`error`, wrapping, custom error)
11. Package & modul, visibility, organisasi kode

### Path 2 — Go Menengah
1. Generics
2. Goroutine
3. Channel (buffered/unbuffered, `select`)
4. `sync` (WaitGroup, Mutex), race condition
5. `context` (cancel, timeout, deadline)
6. File I/O, JSON encoding/decoding
7. Testing (`testing`, table-driven test, mocking, benchmark)
8. Standard library penting (`net/http`, `time`, `log/slog`, `os`)

### Path 3 — Backend Engineering dengan Go
1. Dasar HTTP & REST API
2. Membangun API dengan **Gin** (routing, binding, validation, middleware)
3. Arsitektur berlapis (handler → service → repository)
4. **MySQL**: desain skema, relasi, index, transaksi
5. Akses database (`database/sql` / GORM / sqlx), migrasi
6. Autentikasi & otorisasi (JWT, bcrypt, RBAC)
7. Konfigurasi, environment, logging terstruktur
8. Error handling & response standar API
9. Dokumentasi API (Swagger/OpenAPI)
10. Caching (Redis) dan rate limiting
11. Docker & Docker Compose
12. CI/CD dasar & deployment
13. Observability dasar (logging, metrics, health check)

### Path 4 — Profesional & Siap Kerja
1. Clean architecture / hexagonal
2. Design pattern idiomatik Go
3. Concurrency patterns (worker pool, pipeline, fan-in/out)
4. gRPC & Protobuf
5. Message queue (RabbitMQ/Kafka) — pengenalan
6. Microservices vs monolith
7. Keamanan backend (OWASP, SQL injection, validasi input)
8. Performance tuning & profiling (`pprof`)
9. Code review, Git workflow, kolaborasi tim
10. Persiapan wawancara: soal algoritma, system design, studi kasus

### Proyek Portofolio (Bertahap)
| Proyek | Level | Fokus |
|---|---|---|
| CLI Todo App | Pemula | Dasar Go, file I/O |
| URL Shortener | Menengah | REST API, MySQL |
| Blog / Task Manager API | Menengah | Auth JWT, CRUD, relasi |
| E-commerce API | Lanjut | Transaksi, caching, Docker |
| Realtime Chat / Notification Service | Lanjut | Concurrency, WebSocket |

---

## 6. Fitur Produk

### 6.1 MVP (Prioritas Tinggi — P0)

| ID | Fitur | Deskripsi |
|---|---|---|
| F1 | Autentikasi | Registrasi, login, logout (JWT) |
| F2 | Learning Path & Roadmap | Tampilan peta belajar visual, status tiap materi (terkunci/aktif/selesai) |
| F3 | Halaman Pelajaran | Konten Markdown dengan syntax highlight, diagram, ringkasan "inti dalam 1 menit" |
| F4 | Code Playground | Editor kode (Monaco) + tombol Run, output & error ditampilkan |
| F5 | Latihan Koding | Soal dengan test case otomatis, hint bertingkat, solusi setelah mencoba |
| F6 | Kuis | Pilihan ganda/isian singkat dengan penjelasan jawaban |
| F7 | Progress Tracking | Persentase per path, pelajaran selesai, waktu belajar |
| F8 | Dashboard | Lanjutkan belajar, target harian, ringkasan progres |

### 6.2 Fase Berikutnya (P1)

| ID | Fitur | Deskripsi |
|---|---|---|
| F9 | Spaced Repetition | Flashcard & review harian dengan algoritma SM-2 sederhana |
| F10 | Gamifikasi | XP, level, streak harian, badge |
| F11 | Cheatsheet | Ringkasan cepat per topik, bisa dicari dan di-bookmark |
| F12 | Pencarian | Cari materi, snippet, dan istilah |
| F13 | Catatan Pribadi | Catatan per pelajaran dan bookmark |
| F14 | Dark/Light Mode | Preferensi tampilan tersimpan |
| F15 | Proyek Terpandu | Langkah-langkah membangun proyek dengan checklist dan kode starter |

### 6.3 Masa Depan (P2)

- Placement test (tes awal untuk menentukan titik mulai)
- Rekomendasi materi adaptif berdasarkan kelemahan
- Komunitas/diskusi per pelajaran
- Sertifikat penyelesaian path
- Mock interview & latihan system design
- Review kode otomatis atas proyek pengguna
- Panel admin untuk mengelola konten

---

## 7. Pengalaman Belajar (UX Flow)

### Alur Satu Pelajaran (Pola Konsisten)
1. **Tujuan** — "Setelah ini kamu bisa ...".
2. **Konsep inti** — penjelasan singkat + analogi.
3. **Contoh kode** — bisa langsung dijalankan dan diubah.
4. **Coba sendiri** — latihan kecil dengan umpan balik instan.
5. **Kesalahan umum** — jebakan yang sering dialami pemula.
6. **Kuis cepat** — 3–5 soal.
7. **Ringkasan** — poin kunci + masuk ke antrean review.

### Prinsip UI
- Tampilan dua panel: materi di kiri, editor/output di kanan (responsif untuk mobile: tab).
- Bersih, fokus, minim distraksi.
- Navigasi "Sebelumnya / Berikutnya" yang jelas.
- Pesan error kompiler diterjemahkan menjadi penjelasan ramah.
- Aksesibilitas: kontras warna baik, navigasi keyboard.

### Halaman Utama
| Halaman | Isi |
|---|---|
| Landing | Value proposition, preview roadmap, CTA daftar |
| Login/Register | Form autentikasi |
| Dashboard | Lanjutkan belajar, streak, review hari ini |
| Roadmap | Peta path dan status |
| Lesson | Materi + playground |
| Exercise | Soal koding + test result |
| Review | Flashcard harian |
| Project | Panduan proyek bertahap |
| Cheatsheet | Ringkasan cepat |
| Profile & Settings | Data akun, preferensi, statistik |

---

## 8. Arsitektur Teknis

### 8.1 Gambaran Umum

```
[ React + Vite + Tailwind ]  --HTTPS/JSON-->  [ Go + Gin API ]  -->  [ MySQL ]
                                                    |
                                                    +--> [ Code Runner Sandbox ]
```

### 8.2 Frontend
- **React.js** (Vite) + **React Router**
- **Tailwind CSS** untuk styling
- **TanStack Query** untuk data fetching/caching
- **Zustand** (atau Context) untuk state ringan
- **Monaco Editor** untuk editor kode
- **react-markdown** + **rehype/remark** + highlight kode untuk render materi
- **Axios/fetch** dengan interceptor token

Struktur folder (usulan):
```
frontend/
├── src/
│   ├── api/            # klien API
│   ├── components/     # komponen UI reusable
│   ├── features/       # auth, lessons, playground, review, dashboard
│   ├── hooks/
│   ├── pages/
│   ├── routes/
│   ├── store/
│   └── utils/
├── index.html
├── tailwind.config.js
└── vite.config.js
```

### 8.3 Backend
- **Go** + **Gin**
- Arsitektur berlapis: `handler → service → repository`
- Akses DB: GORM atau sqlx (pilih satu; disarankan sqlx + migrasi eksplisit untuk pembelajaran yang transparan)
- Migrasi: `golang-migrate`
- Autentikasi: JWT (access + refresh token), hash password dengan bcrypt
- Validasi: `go-playground/validator`
- Konfigurasi: environment variable (`.env`)
- Logging: `log/slog`
- Dokumentasi API: Swagger (`swaggo`)

Struktur folder (usulan):
```
backend/
├── cmd/api/main.go
├── internal/
│   ├── config/
│   ├── handler/
│   ├── middleware/
│   ├── model/
│   ├── repository/
│   ├── service/
│   ├── runner/          # eksekusi kode
│   └── router/
├── migrations/
├── pkg/
├── Dockerfile
└── go.mod
```

### 8.4 Eksekusi Kode (Komponen Kritis)

Menjalankan kode pengguna adalah risiko keamanan terbesar. Opsi:

| Opsi | Kelebihan | Kekurangan |
|---|---|---|
| A. Go Playground API (pihak ketiga) | Cepat dibuat, aman | Bergantung layanan luar, terbatas fitur |
| B. Sandbox sendiri (Docker, tanpa jaringan) | Kontrol penuh, bisa test case | Lebih kompleks, perlu hardening |
| C. WebAssembly di browser | Tidak membebani server | Dukungan Go di wasm terbatas |

**Rekomendasi:** MVP memakai **Opsi A** untuk playground bebas; fase berikutnya beralih ke **Opsi B** untuk latihan dengan test case otomatis.

**Wajib untuk Opsi B:**
- Container terisolasi, tanpa akses jaringan
- Batas CPU, memori, dan waktu eksekusi (mis. 5 detik)
- Filesystem read-only, user non-root
- Batas ukuran kode dan output
- Rate limiting per pengguna
- Pembersihan container setelah selesai

### 8.5 Skema Database MySQL (Usulan)

```sql
users (id, name, email UNIQUE, password_hash, role, created_at, updated_at)

paths (id, slug UNIQUE, title, description, level, order_no)

modules (id, path_id FK, title, order_no)

lessons (id, module_id FK, slug UNIQUE, title, summary,
         content_md, estimated_minutes, order_no, is_published)

exercises (id, lesson_id FK, title, instructions_md, starter_code,
           solution_code, difficulty, xp_reward)

exercise_tests (id, exercise_id FK, input, expected_output, is_hidden)

quizzes (id, lesson_id FK, question, type, explanation)
quiz_options (id, quiz_id FK, text, is_correct)

user_lesson_progress (id, user_id FK, lesson_id FK, status,
                      completed_at, time_spent_sec, UNIQUE(user_id, lesson_id))

user_exercise_submissions (id, user_id FK, exercise_id FK, code,
                           is_passed, output, created_at)

user_quiz_attempts (id, user_id FK, quiz_id FK, is_correct, created_at)

flashcards (id, lesson_id FK, front, back)
user_flashcard_reviews (id, user_id FK, flashcard_id FK, ease_factor,
                        interval_days, repetitions, next_review_at,
                        UNIQUE(user_id, flashcard_id))

user_stats (user_id PK/FK, xp, level, current_streak, longest_streak, last_active_date)

badges (id, code UNIQUE, name, description)
user_badges (user_id FK, badge_id FK, earned_at)

notes (id, user_id FK, lesson_id FK, content, updated_at)
bookmarks (user_id FK, lesson_id FK, created_at)

projects (id, slug, title, description, level)
project_steps (id, project_id FK, title, content_md, order_no)
user_project_progress (user_id FK, step_id FK, is_done, completed_at)
```

Index penting: `users.email`, `lessons.slug`, `user_lesson_progress(user_id, lesson_id)`, `user_flashcard_reviews(user_id, next_review_at)`.

### 8.6 Daftar Endpoint API (Usulan)

Base URL: `/api/v1`

**Auth**
- `POST /auth/register`
- `POST /auth/login`
- `POST /auth/refresh`
- `POST /auth/logout`
- `GET  /auth/me`

**Konten**
- `GET  /paths`
- `GET  /paths/:slug`
- `GET  /lessons/:slug`
- `GET  /lessons/:slug/exercises`
- `GET  /lessons/:slug/quizzes`
- `GET  /search?q=`

**Interaksi**
- `POST /playground/run`
- `POST /exercises/:id/submit`
- `POST /quizzes/:id/answer`
- `POST /lessons/:slug/complete`

**Progres & Review**
- `GET  /me/progress`
- `GET  /me/dashboard`
- `GET  /me/review/today`
- `POST /me/review/:flashcardId` (rating kualitas ingatan)
- `GET  /me/stats`

**Catatan & Bookmark**
- `GET/POST/PUT/DELETE /me/notes`
- `GET/POST/DELETE /me/bookmarks`

**Proyek**
- `GET  /projects`
- `GET  /projects/:slug`
- `POST /projects/steps/:id/complete`

**Admin (P2)**
- CRUD `paths`, `lessons`, `exercises`, `quizzes`

**Format Response Standar**
```json
{
  "success": true,
  "data": {},
  "error": null,
  "meta": { "page": 1, "limit": 20, "total": 100 }
}
```

---

## 9. Persyaratan Non-Fungsional

| Aspek | Persyaratan |
|---|---|
| Performa | Halaman utama < 2 dtk (LCP); API p95 < 300 ms; run kode < 5 dtk |
| Keamanan | Hash bcrypt, JWT, CORS ketat, validasi input, rate limit, proteksi SQL injection & XSS, sandbox terisolasi |
| Skalabilitas | Backend stateless, siap horizontal scaling |
| Ketersediaan | Target uptime 99% (tahap awal) |
| Responsif | Desktop, tablet, dan mobile |
| Aksesibilitas | Mengacu WCAG AA dasar |
| Maintainability | Test unit (service/handler), lint, format (`gofmt`, ESLint, Prettier) |
| Lokalisasi | Bahasa Indonesia sebagai default; istilah teknis Inggris dipertahankan |
| Privasi | Data minimal, kebijakan privasi, hak hapus akun |

---

## 10. Rencana Pengembangan (Roadmap)

| Fase | Durasi (perkiraan) | Hasil |
|---|---|---|
| **Fase 0 — Persiapan** | 1 minggu | Repo, struktur proyek, desain UI dasar (wireframe), setup Docker Compose (MySQL), CI dasar |
| **Fase 1 — Fondasi** | 2 minggu | Auth, skema DB + migrasi, CRUD konten (seed), layout frontend, routing |
| **Fase 2 — Inti Belajar** | 3 minggu | Halaman pelajaran (Markdown), playground, kuis, progress tracking, dashboard |
| **Fase 3 — Konten MVP** | 2–3 minggu (paralel) | Path 1 (Fondasi Go) lengkap dengan latihan dan kuis |
| **Fase 4 — Latihan & Sandbox** | 3 minggu | Sandbox Docker, test case otomatis, hint bertingkat |
| **Fase 5 — Retensi** | 2 minggu | Spaced repetition, streak, XP, badge, cheatsheet |
| **Fase 6 — Backend Track** | 4 minggu | Path 3 + proyek terpandu (REST API, MySQL, JWT, Docker) |
| **Fase 7 — Polish & Rilis Beta** | 2 minggu | Optimasi, QA, keamanan, analytics, deployment |
| **Fase 8 — Lanjutan** | Berkelanjutan | Path 2 & 4, komunitas, adaptif, panel admin |

> Catatan: estimasi untuk satu pengembang; sesuaikan dengan waktu luang. Produksi **konten** biasanya memakan waktu lebih lama daripada kode, jadi rencanakan secara paralel.

### Definisi Selesai MVP
- Pengguna bisa daftar, login, dan mengikuti seluruh Path 1.
- Setiap pelajaran punya contoh kode interaktif, latihan, dan kuis.
- Progres tersimpan dan tampil di dashboard.
- Playground berjalan aman dengan batas waktu dan rate limit.
- Dideploy dan dapat diakses publik.

---

## 11. Risiko & Mitigasi

| Risiko | Dampak | Mitigasi |
|---|---|---|
| Eksekusi kode berbahaya | Tinggi | Sandbox terisolasi, tanpa jaringan, batas resource, rate limit |
| Produksi konten lambat | Tinggi | Mulai dari Path 1, template pelajaran seragam, bantuan AI untuk draft lalu tinjau manual |
| Scope terlalu besar | Sedang | Disiplin MVP, fitur P1/P2 ditunda |
| Biaya server eksekusi kode | Sedang | Mulai dengan Playground API, caching hasil, batas penggunaan |
| Pengguna kehilangan motivasi | Sedang | Streak, target kecil, progres yang terlihat |
| Konten kedaluwarsa (versi Go berubah) | Rendah | Versi konten, tinjauan berkala tiap rilis Go |

---

## 12. Asumsi & Pertanyaan Terbuka

**Asumsi**
- Konten utama berbahasa Indonesia.
- Tahap awal gratis; model bisnis ditentukan kemudian.
- Konten ditulis dalam Markdown dan di-*seed* ke database.

**Perlu Diputuskan**
1. GORM atau sqlx/`database/sql` murni? (Disarankan sqlx untuk pemahaman SQL)
2. Playground: Go Playground API dulu atau langsung sandbox sendiri?
3. Apakah akan ada login sosial (Google/GitHub)?
4. Model monetisasi (gratis, freemium, atau berbayar)?
5. Perlu fitur komunitas/forum di rilis awal?
6. Hosting yang dipilih (VPS, Railway, Fly.io, dll.)?

---

## 13. Lampiran: Template Pelajaran (Markdown)

```markdown
---
slug: variabel-dan-tipe-data
title: Variabel dan Tipe Data
level: pemula
estimated_minutes: 10
---

## 🎯 Tujuan
Setelah ini kamu bisa mendeklarasikan variabel dan memilih tipe data yang tepat.

## ⚡ Inti dalam 1 Menit
(ringkasan 3–4 poin)

## 📖 Konsep
(penjelasan + analogi)

## 💻 Contoh
(kode yang bisa dijalankan)

## ✍️ Coba Sendiri
(latihan kecil)

## ⚠️ Kesalahan Umum
(daftar jebakan)

## ✅ Kuis
(3–5 soal)

## 📝 Ringkasan
(poin kunci + flashcard)
```

---

*Dokumen ini bersifat hidup dan dapat diperbarui seiring perkembangan proyek.*
