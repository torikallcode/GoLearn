---
slug: pengenalan-go
title: Pengenalan Go & Struktur Proyek
summary: Kenali Go, pasang toolchain, dan jalankan program pertama dengan go mod.
minutes: 12
exercise:
  title: Program Go pertama
  difficulty: mudah
  xp: 20
  instructions: |
    Lengkapi program supaya mencetak tepat satu baris:

    Halo, Go!

    Gunakan fmt.Println. Perhatikan huruf besar, tanda koma, dan tanda seru.
  starter: |
    package main

    import "fmt"

    func main() {
        // TODO: cetak "Halo, Go!"
    }
  solution: |
    package main

    import "fmt"

    func main() {
        fmt.Println("Halo, Go!")
    }
  hints:
    - Setiap program Go yang bisa dijalankan punya package main dan func main.
    - fmt.Println mencetak teks lalu menambahkan newline.
  tests:
    - expected: |
        Halo, Go!
quizzes:
  - question: Apa nama fungsi yang menjadi titik masuk program Go?
    explanation: Program yang dapat dieksekusi memakai package main dan fungsi main sebagai entry point.
    options:
      - text: func main()
        correct: true
      - text: func start()
      - text: func init()
      - text: func run()
  - question: Perintah apa yang membuat file go.mod untuk proyek baru?
    explanation: go mod init <nama-modul> membuat go.mod dan menandai folder sebagai module Go.
    options:
      - text: go mod init contoh
        correct: true
      - text: go build init
      - text: go get init
      - text: go run mod
  - question: Apa fungsi perintah go run main.go?
    explanation: go run mengompilasi lalu langsung menjalankan program tanpa meninggalkan binary.
    options:
      - text: Mengompilasi dan menjalankan program
        correct: true
      - text: Hanya memformat kode
      - text: Mengunduh dependency
      - text: Membuat dokumentasi
flashcards:
  - front: Apa dua syarat program Go yang bisa dijalankan?
    back: Package main dan fungsi main().
  - front: Apa isi file go.mod?
    back: Nama module, versi Go, dan daftar dependency yang dipakai proyek.
---

## 🎯 Tujuan
Setelah ini kamu bisa memasang Go, membuat module, menjalankan program pertama, dan mengenali struktur proyek yang rapi.

## ⚡ Inti dalam 1 Menit
- Go adalah bahasa compiled, statically typed, dan terkenal karena sederhana serta concurrency-nya yang mudah.
- Program yang bisa dijalankan selalu punya `package main` dan `func main()`.
- `go mod init` membuat module; `go run` menjalankan; `go build` menghasilkan binary.
- Kode Go diformat otomatis dengan `gofmt` (atau `go fmt`).

## 📖 Konsep
Go dibuat di Google untuk menggabungkan kemudahan bahasa dinamis dengan performa bahasa compiled. Hasil kompilasinya berupa satu binary mandiri, cocok untuk backend dan container.

Anggap module seperti "folder proyek resmi". Di dalamnya ada `go.mod` yang mencatat nama module, versi Go, dan dependency. Semua package di dalam module bisa saling di-import memakai nama module sebagai awalan.

Struktur kecil yang umum:

    hello/
    ├── go.mod
    ├── main.go
    └── internal/
        └── greeting/greeting.go

`main.go` adalah pintu masuk, `internal/` berisi kode yang tidak boleh di-import proyek lain.

## 💻 Contoh
```go
package main

import "fmt"

func main() {
	name := "Backend Engineer"
	fmt.Printf("Halo, %s!\n", name)
}
```

Jalankan:

```bash
go mod init hello
go run .
```

`go run .` berarti "kompilasi dan jalankan package di folder ini".

## ✍️ Coba Sendiri
Lengkapi latihan di panel kanan sampai mencetak `Halo, Go!` lalu tekan Submit. Setelah lolos, tekan Tandai Selesai.

## ⚠️ Kesalahan Umum
- Menulis `package Main` — Go case-sensitive, harus `main`.
- Lupa import `fmt` saat memakai `fmt.Println`.
- Mendeklarasikan variabel lalu tidak memakainya; Go menolak compile.
- Menjalankan `go run main.go` dari folder yang salah.

## ✅ Kuis
Jawab kuis di bawah. Setiap jawaban langsung diberi penjelasan.

## 📝 Ringkasan
- `package main` + `func main()` adalah titik masuk.
- `go mod init` menyiapkan module, `go run .` menjalankan.
- `gofmt` menjaga format konsisten.
- Gunakan `internal/` untuk kode privat proyek.
