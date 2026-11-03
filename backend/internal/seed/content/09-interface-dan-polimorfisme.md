---
slug: interface-dan-polimorfisme
title: Interface & Polimorfisme
summary: Kontrak perilaku dengan interface dan cara memakainya secara idiomatik.
minutes: 14
exercise:
  title: Hitung total luas
  difficulty: sedang
  xp: 25
  instructions: |
    Lengkapi interface Bentuk dan dua implementasinya.

    Output harus:
    total=43

    Gunakan nilai π = 3 untuk menyederhanakan. Persegi sisi 5, lingkaran jari-jari 3.
  starter: |
    package main

    import "fmt"

    type Bentuk interface {
        Luas() int
    }

    // TODO: buat Persegi dan Lingkaran yang memenuhi interface Bentuk.

    func totalLuas(bentuk ...Bentuk) int {
        total := 0
        for _, b := range bentuk {
            total += b.Luas()
        }
        return total
    }

    func main() {
        fmt.Printf("total=%d\n", totalLuas(Persegi{Sisi: 5}, Lingkaran{JariJari: 3}))
    }
  solution: |
    package main

    import "fmt"

    type Bentuk interface {
        Luas() int
    }

    type Persegi struct{ Sisi int }

    func (p Persegi) Luas() int { return p.Sisi * p.Sisi }

    type Lingkaran struct{ JariJari int }

    func (l Lingkaran) Luas() int { return 3 * l.JariJari * l.JariJari }

    func totalLuas(bentuk ...Bentuk) int {
        total := 0
        for _, b := range bentuk {
            total += b.Luas()
        }
        return total
    }

    func main() {
        fmt.Printf("total=%d\n", totalLuas(Persegi{Sisi: 5}, Lingkaran{JariJari: 3}))
    }
  hints:
    - Interface dipenuhi otomatis jika method-nya cocok.
    - Persegi dan Lingkaran tidak perlu menulis "implements".
  tests:
    - expected: |
        total=43
quizzes:
  - question: Bagaimana sebuah tipe memenuhi interface di Go?
    explanation: Secara implisit, cukup punya semua method dengan signature yang sesuai.
    options:
      - text: Otomatis jika punya semua method yang diminta
        correct: true
      - text: Dengan kata kunci implements
      - text: Dengan mendaftar ke interface
      - text: Dengan embedding interface
  - question: Di mana interface sebaiknya dideklarasikan?
    explanation: Interface yang kecil paling baik dideklarasikan di package yang memakainya, bukan di package implementasi.
    options:
      - text: Di package yang memakainya
        correct: true
      - text: Selalu di package implementasi
      - text: Di file main saja
      - text: Di package model
  - question: Apa yang terjadi jika method interface dipanggil pada interface nil?
    explanation: Memanggil method pada interface nil menyebabkan panic karena tidak ada nilai konkret.
    options:
      - text: Panic
        correct: true
      - text: Mengembalikan zero value
      - text: Error compile
      - text: Tidak terjadi apa-apa
flashcards:
  - front: Berapa jumlah method yang ideal untuk interface Go?
    back: Kecil, sering satu. Interface kecil lebih mudah dipenuhi dan dites.
  - front: Apa itu empty interface any?
    back: Interface tanpa method; semua tipe memenuhinya. Gunakan hanya saat benar-benar perlu.
---

## 🎯 Tujuan
Setelah ini kamu bisa mendefinisikan kontrak perilaku, membuat implementasi tanpa deklarasi khusus, dan menghindari interface yang terlalu besar.

## ⚡ Inti dalam 1 Menit
- Interface adalah kumpulan signature method.
- Implementasi bersifat implisit; tidak ada kata kunci `implements`.
- Interface kecil lebih fleksibel dan mudah dites.
- Interface nil bisa panic jika method dipanggil.

## 📖 Konsep
```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

Tipe apa pun yang punya `Read` dengan signature tersebut otomatis memenuhi `Reader`. Ini membuat `os.File`, `bytes.Buffer`, dan `strings.Reader` bisa dipakai di fungsi yang sama.

Gunakan interface saat:

- butuh polimorfisme (banyak tipe, satu perilaku),
- ingin mengganti dependency saat test (mock),
- mendesain batas antar-layer.

Jangan membuat interface sebelum ada dua implementasi atau kebutuhan nyata. "Accept interfaces, return structs" adalah panduan umum.

## 💻 Contoh
```go
package main

import "fmt"

type Notifier interface {
	Send(pesan string) error
}

type Email struct{ Alamat string }

func (e Email) Send(pesan string) error {
	fmt.Println("email ke", e.Alamat, ":", pesan)
	return nil
}

type SMS struct{ Nomor string }

func (s SMS) Send(pesan string) error {
	fmt.Println("sms ke", s.Nomor, ":", pesan)
	return nil
}

func kirimSemua(n Notifier, pesan string) {
	if err := n.Send(pesan); err != nil {
		fmt.Println("gagal:", err)
	}
}

func main() {
	kirimSemua(Email{Alamat: "a@example.com"}, "Halo")
	kirimSemua(SMS{Nomor: "0812"}, "Halo")
}
```

## ✍️ Coba Sendiri
Selesaikan latihan. Setelah lolos, tambahkan bentuk Segitiga yang juga memenuhi `Bentuk`.

## ⚠️ Kesalahan Umum
- Membuat interface raksasa yang sulit diimplementasikan.
- Menyimpan nil pointer ke dalam interface lalu memanggil method.
- Mendeklarasikan interface di package implementasi.
- Memakai `any` untuk menghindari tipe yang sebenarnya jelas.

## ✅ Kuis
Tiga soal di bawah.

## 📝 Ringkasan
- Interface = kontrak method.
- Implementasi implisit.
- Kecil lebih baik.
- Interface nil bisa panic.
