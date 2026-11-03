---
slug: struct-dan-method
title: Struct & Method
summary: Kelompokkan data dengan struct dan beri perilaku lewat method.
minutes: 14
exercise:
  title: Persegi panjang
  difficulty: mudah
  xp: 20
  instructions: |
    Lengkapi struct dan method supaya output menjadi:

    luas=24 keliling=20

    Gunakan method Luas dan Keliling pada struct PersegiPanjang.
  starter: |
    package main

    import "fmt"

    type PersegiPanjang struct {
        Panjang float64
        Lebar   float64
    }

    // TODO: method Luas dan Keliling

    func main() {
        p := PersegiPanjang{Panjang: 6, Lebar: 4}
        fmt.Printf("luas=%g keliling=%g\n", p.Luas(), p.Keliling())
    }
  solution: |
    package main

    import "fmt"

    type PersegiPanjang struct {
        Panjang float64
        Lebar   float64
    }

    func (p PersegiPanjang) Luas() float64 {
        return p.Panjang * p.Lebar
    }

    func (p PersegiPanjang) Keliling() float64 {
        return 2 * (p.Panjang + p.Lebar)
    }

    func main() {
        p := PersegiPanjang{Panjang: 6, Lebar: 4}
        fmt.Printf("luas=%g keliling=%g\n", p.Luas(), p.Keliling())
    }
  hints:
    - Method adalah fungsi dengan receiver, misalnya func (p PersegiPanjang) Luas() float64.
    - "%g mencetak angka tanpa nol di belakang desimal."
  tests:
    - expected: |
        luas=24 keliling=20
quizzes:
  - question: Apa perbedaan method dengan receiver value dan pointer?
    explanation: Receiver pointer bisa mengubah struct asli dan menghindari salinan; receiver value bekerja pada salinan.
    options:
      - text: Pointer bisa mengubah nilai asli, value bekerja pada salinan
        correct: true
      - text: Keduanya identik
      - text: Value selalu lebih cepat
      - text: Pointer tidak bisa punya method
  - question: Apa arti field dengan huruf kapital di awal?
    explanation: Identifier yang diawali huruf kapital diekspor dan bisa diakses package lain.
    options:
      - text: Field tersebut exported
        correct: true
      - text: Field tersebut readonly
      - text: Field tersebut private
      - text: Field tersebut static
  - question: Bagaimana membuat struct tanpa memberi nama field?
    explanation: Literal seperti Titik{1, 2} memakai urutan field. Rentan rusak jika urutan field berubah, jadi lebih aman memakai nama field.
    options:
      - text: Titik{1, 2}
        correct: true
      - text: "Titik{X: 1, Y: 2} saja"
      - text: new(Titik){1, 2}
      - text: struct Titik{1, 2}
flashcards:
  - front: Kapan memakai receiver pointer?
    back: Saat method perlu mengubah data atau struct cukup besar sehingga menyalin mahal.
  - front: Apa arti struct tag?
    back: Metadata di belakang field, misalnya `json:"nama"`, dibaca package seperti encoding/json.
---

## 🎯 Tujuan
Setelah ini kamu bisa memodelkan data dengan struct, memberi method, dan memilih receiver value atau pointer.

## ⚡ Inti dalam 1 Menit
- Struct mengelompokkan field bertipe.
- Method adalah fungsi dengan receiver.
- Receiver pointer (`*T`) bisa mengubah data; receiver value (`T`) bekerja pada salinan.
- Field/method berhuruf kapital diekspor ke package lain.

## 📖 Konsep
```go
type User struct {
	ID    int
	Name  string
	Email string
}

func (u User) Display() string {
	return fmt.Sprintf("%s <%s>", u.Name, u.Email)
}

func (u *User) Rename(nama string) {
	u.Name = nama
}
```

Struct bisa disematkan (embedding):

```go
type Tim struct {
	Nama  string
	Admin User
}
```

Embedding anonymous memberi akses field secara langsung, tetapi komposisi eksplisit sering lebih jelas untuk pemula.

## 💻 Contoh
```go
package main

import "fmt"

type Rekening struct {
	Saldo float64
}

func (r *Rekening) Setor(nominal float64) {
	r.Saldo += nominal
}

func main() {
	r := &Rekening{}
	r.Setor(50000)
	fmt.Println(r.Saldo)
}
```

Karena `Setor` memakai pointer receiver, perubahan `Saldo` terlihat di luar method.

## ✍️ Coba Sendiri
Selesaikan latihan persegi panjang, lalu ubah Lebar menjadi 2 dan lihat outputnya.

## ⚠️ Kesalahan Umum
- Lupa `*` pada receiver sehingga perubahan tidak tersimpan.
- Membuat struct dengan urutan field lalu menambah field baru.
- Memakai struct besar sebagai value di banyak fungsi tanpa alasan.
- Menyimpan password di struct yang ikut di-JSON tanpa tag `json:"-"`.

## ✅ Kuis
Tiga soal di bawah.

## 📝 Ringkasan
- Struct untuk data, method untuk perilaku.
- Pointer receiver untuk mutasi.
- Nama kapital = exported.
- Gunakan named fields saat membuat literal.
