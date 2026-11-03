---
slug: package-dan-modul
title: Package, Modul & Visibility
summary: "Organisasi kode Go: package, module, exported identifier, dan dependency."
minutes: 13
exercise:
  title: Simulasi helper package
  difficulty: sedang
  xp: 25
  instructions: |
    Di satu file, buat fungsi exported FormatNama dan helper unexported rapikan.

    Output harus:
    Rina Wijaya
    rina-wijaya

    FormatNama mengembalikan nama dengan spasi tunggal.
    rapikan digunakan oleh Slug.
  starter: |
    package main

    import (
        "fmt"
        "strings"
    )

    // TODO: rapikan, FormatNama, dan Slug

    func main() {
        fmt.Println(FormatNama("  Rina   Wijaya "))
        fmt.Println(Slug("Rina Wijaya"))
    }
  solution: |
    package main

    import (
        "fmt"
        "strings"
    )

    func rapikan(s string) string {
        return strings.Join(strings.Fields(s), " ")
    }

    func FormatNama(s string) string {
        return rapikan(s)
    }

    func Slug(s string) string {
        return strings.ReplaceAll(strings.ToLower(rapikan(s)), " ", "-")
    }

    func main() {
        fmt.Println(FormatNama("  Rina   Wijaya "))
        fmt.Println(Slug("Rina Wijaya"))
    }
  hints:
    - strings.Fields memecah string berdasarkan whitespace dan membuang spasi berlebih.
    - strings.Join menggabungkan slice string dengan pemisah.
  tests:
    - expected: |
        Rina Wijaya
        rina-wijaya
quizzes:
  - question: Apa yang menentukan sebuah identifier bisa diakses package lain?
    explanation: Huruf kapital di awal identifier membuatnya exported.
    options:
      - text: Huruf kapital di awal nama
        correct: true
      - text: Kata kunci public
      - text: Nama file
      - text: Tag struct
  - question: Apa fungsi folder internal pada module Go?
    explanation: Package di bawah internal hanya bisa di-import oleh kode yang berada dalam parent folder internal tersebut.
    options:
      - text: Membatasi import ke dalam module induknya
        correct: true
      - text: Menyembunyikan file dari git
      - text: Membuat package otomatis teruji
      - text: Menghapus dependency
  - question: Perintah apa untuk merapikan dependency go.mod?
    explanation: go mod tidy menambahkan dependency yang dipakai dan menghapus yang tidak dipakai.
    options:
      - text: go mod tidy
        correct: true
      - text: go clean mod
      - text: go fmt mod
      - text: go get tidy
flashcards:
  - front: Apa perbedaan package dan module?
    back: Package adalah kumpulan file Go dalam satu folder; module adalah kumpulan package dengan go.mod.
  - front: Apa fungsi go mod tidy?
    back: Menyelaraskan go.mod/go.sum dengan import yang benar-benar dipakai.
---

## 🎯 Tujuan
Setelah ini kamu bisa memecah kode menjadi package yang rapi, memilih nama exported, dan mengelola dependency module.

## ⚡ Inti dalam 1 Menit
- Satu folder = satu package (kecuali `_test`).
- `package main` menghasilkan program; package lain adalah library.
- Huruf kapital di awal = exported, huruf kecil = privat package.
- `go mod tidy` merapikan dependency; `internal/` membatasi import.

## 📖 Konsep
Struktur yang tumbuh sehat:

    toko/
    ├── go.mod
    ├── cmd/api/main.go
    ├── internal/
    │   ├── handler/
    │   ├── service/
    │   └── repository/
    └── pkg/
        └── validator/

- `cmd/` berisi entry point aplikasi.
- `internal/` berisi kode privat aplikasi.
- `pkg/` untuk library yang boleh dipakai proyek lain (opsional).

Import memakai path module:

```go
import "contoh/toko/internal/service"
```

Package tidak boleh saling import melingkar. Jika terjadi, biasanya batas tanggung jawab perlu dirapikan.

## 💻 Contoh
```go
package main

import (
	"fmt"
	"strings"
)

func rapikan(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func Slug(s string) string {
	return strings.ReplaceAll(strings.ToLower(rapikan(s)), " ", "-")
}

func main() {
	fmt.Println(Slug("  Belajar Go  "))
}
```

`rapikan` hanya bisa dipakai di package yang sama. `Slug` diekspor karena dibutuhkan package lain.

## ✍️ Coba Sendiri
Selesaikan latihan. Setelah lolos, pindahkan fungsi ke package terpisah `internal/text` dan import dari `main`.

## ⚠️ Kesalahan Umum
- Nama package dan folder berbeda tanpa alasan.
- Membuat package `util` yang berisi segalanya.
- Import melingkar antar package.
- Lupa `go mod tidy` setelah menambah dependency.

## ✅ Kuis
Tiga soal di bawah.

## 📝 Ringkasan
- Folder = package.
- Kapital = exported.
- Module mengelola dependency lewat go.mod.
- `internal/` membatasi akses.
