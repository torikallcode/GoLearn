---
slug: pointer
title: Pointer
summary: Alamat memori, dereference, dan kapan pointer benar-benar dibutuhkan.
minutes: 12
exercise:
  title: Naikkan skor lewat pointer
  difficulty: mudah
  xp: 20
  instructions: |
    Lengkapi fungsi tambahSkor supaya mengubah nilai asli.

    Output harus:
    skor=15
  starter: |
    package main

    import "fmt"

    func tambahSkor(skor *int, tambah int) {
        // TODO: tambahkan nilai ke variabel yang ditunjuk
    }

    func main() {
        skor := 10
        tambahSkor(&skor, 5)
        fmt.Printf("skor=%d\n", skor)
    }
  solution: |
    package main

    import "fmt"

    func tambahSkor(skor *int, tambah int) {
        *skor += tambah
    }

    func main() {
        skor := 10
        tambahSkor(&skor, 5)
        fmt.Printf("skor=%d\n", skor)
    }
  hints:
    - Gunakan & untuk mengambil alamat variabel.
    - Gunakan *skor untuk mengakses nilai yang ditunjuk.
  tests:
    - expected: |
        skor=15
quizzes:
  - question: Apa yang dilakukan operator &?
    explanation: "& mengambil alamat memori dari sebuah variabel, menghasilkan pointer."
    options:
      - text: Mengambil alamat variabel
        correct: true
      - text: Mengalikan dua nilai
      - text: Mengambil nilai yang ditunjuk
      - text: Membuat variabel baru
  - question: Apa zero value dari pointer?
    explanation: Pointer yang tidak diinisialisasi bernilai nil; dereference pointer nil menyebabkan panic.
    options:
      - text: nil
        correct: true
      - text: "0"
      - text: alamat nol
      - text: string kosong
  - question: Apakah Go mendukung pointer arithmetic seperti C?
    explanation: Tidak. Go sengaja membatasi pointer agar lebih aman; gunakan slice untuk akses berurutan.
    options:
      - text: Tidak
        correct: true
      - text: Ya, dengan ++
      - text: Ya, tetapi hanya di unsafe
      - text: Ya, untuk array saja
flashcards:
  - front: Apa arti * pada tipe dan pada ekspresi?
    back: Pada tipe berarti pointer; pada ekspresi berarti dereference (ambil nilai yang ditunjuk).
  - front: Kapan sebaiknya mengembalikan pointer?
    back: Saat perlu menandai "tidak ada" dengan nil, atau agar tidak menyalin struct besar.
---

## 🎯 Tujuan
Setelah ini kamu bisa memakai pointer tanpa takut, tahu kapan perlu, dan mengenali panic akibat pointer nil.

## ⚡ Inti dalam 1 Menit
- `&x` mengambil alamat; `*p` mengambil nilai di alamat.
- Pointer memungkinkan fungsi mengubah nilai pemanggil.
- Zero value pointer adalah `nil`; dereference nil = panic.
- Go tidak punya pointer arithmetic.

## 📖 Konsep
Saat memanggil fungsi, argumen disalin. Pointer membuat salinan berisi alamat, sehingga data asli bisa diubah.

```go
func reset(p *int) {
	*p = 0
}

n := 42
reset(&n)
fmt.Println(n) // 0
```

Pointer juga dipakai untuk membedakan "tidak diisi" dari zero value, misalnya `*string` pada update parsial API.

```go
type UpdateUser struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
}
```

Jika field nil, berarti klien tidak mengirim field itu.

## 💻 Contoh
```go
package main

import "fmt"

type Counter struct{ n int }

func (c *Counter) Inc() { c.n++ }

func main() {
	c := &Counter{}
	c.Inc()
	c.Inc()
	fmt.Println(c.n)
}
```

## ✍️ Coba Sendiri
Selesaikan latihan lalu ubah `tambahSkor` agar mengembalikan nilai baru sebagai alternatif tanpa pointer.

## ⚠️ Kesalahan Umum
- Lupa `&` saat memanggil fungsi berparameter pointer.
- Melakukan `*p` pada pointer nil.
- Mengira `*` di tipe sama dengan dereference; konteksnya berbeda.
- Memakai pointer untuk tipe kecil padahal value lebih sederhana.

## ✅ Kuis
Tiga soal di bawah.

## 📝 Ringkasan
- `&` alamat, `*` nilai.
- Pointer untuk mutasi dan "optional".
- Pointer nil harus dicek.
- Tanpa pointer arithmetic di Go.
