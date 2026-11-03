---
slug: variabel-dan-tipe-data
title: Variabel, Tipe Data & Zero Value
summary: Deklarasi variabel dengan var dan :=, pahami tipe dasar dan zero value.
minutes: 12
exercise:
  title: Kartu profil
  difficulty: mudah
  xp: 20
  instructions: |
    Cetak profil dengan format tepat:

    Nama: Rina
    Umur: 21
    Aktif: true
    Skor: 87.5

    Gunakan variabel bertipe string, int, bool, dan float64.
  starter: |
    package main

    import "fmt"

    func main() {
        nama := "Rina"
        umur := 21
        // TODO: buat variabel aktif (bool) dan skor (float64), lalu cetak.
        _ = nama
        _ = umur
    }
  solution: |
    package main

    import "fmt"

    func main() {
        nama := "Rina"
        umur := 21
        aktif := true
        skor := 87.5
        fmt.Println("Nama:", nama)
        fmt.Println("Umur:", umur)
        fmt.Println("Aktif:", aktif)
        fmt.Println("Skor:", skor)
    }
  hints:
    - fmt.Println memisahkan argumen dengan spasi otomatis.
    - true dan false ditulis tanpa tanda kutip.
  tests:
    - expected: |
        Nama: Rina
        Umur: 21
        Aktif: true
        Skor: 87.5
quizzes:
  - question: Berapa zero value dari tipe int?
    explanation: Variabel int yang tidak diberi nilai otomatis bernilai 0.
    options:
      - text: "0"
        correct: true
      - text: "1"
      - text: "nil"
      - text: '""'
  - question: Kapan operator := boleh dipakai?
    explanation: := melakukan deklarasi sekaligus assignment, hanya di dalam fungsi. Untuk level package gunakan var.
    options:
      - text: Di dalam fungsi untuk deklarasi singkat
        correct: true
      - text: Di level package
      - text: Untuk mengganti tipe variabel
      - text: Hanya untuk konstanta
  - question: Apa tipe data yang cocok untuk harga 19.99?
    explanation: float64 menyimpan bilangan pecahan. Untuk uang produksi pertimbangkan integer terkecil (sen) untuk menghindari pembulatan.
    options:
      - text: float64
        correct: true
      - text: int
      - text: bool
      - text: rune
flashcards:
  - front: Sebutkan zero value string, bool, dan int.
    back: string = "", bool = false, int = 0.
  - front: Apa perbedaan var dan := ?
    back: var bisa di level package maupun fungsi dan boleh tanpa nilai; := deklarasi singkat di dalam fungsi yang wajib punya nilai.
---

## 🎯 Tujuan
Setelah ini kamu bisa memilih tipe data yang tepat, mendeklarasikan variabel dengan benar, dan menjelaskan zero value.

## ⚡ Inti dalam 1 Menit
- Tipe dasar: `string`, `bool`, `int`, `int64`, `float64`, `rune`, `byte`.
- `var nama tipe` untuk deklarasi; `nama := nilai` untuk deklarasi singkat di dalam fungsi.
- Variabel tanpa nilai mendapat **zero value**.
- Gunakan `const` untuk nilai yang tidak berubah.

## 📖 Konsep
Go statically typed: tipe ditentukan saat compile dan tidak berubah. Ini mencegah banyak bug sebelum program jalan.

Zero value berguna karena variabel selalu punya nilai yang aman:

| Tipe | Zero value |
|---|---|
| string | `""` |
| int | `0` |
| float64 | `0` |
| bool | `false` |
| pointer, slice, map | `nil` |

Konversi tipe harus eksplisit: `float64(umur)`. Go tidak mengonversi otomatis walaupun tipe mirip.

## 💻 Contoh
```go
package main

import "fmt"

func main() {
	var nama string        // ""
	var umur int           // 0
	aktif := true          // bool
	skor := 87.5           // float64
	const maksimal = 100

	fmt.Println(nama, umur, aktif, skor, maksimal)
	fmt.Printf("%s berumur %d tahun\n", "Rina", 21)
}
```

## ✍️ Coba Sendiri
Selesaikan latihan kartu profil. Pastikan urutan dan spasi output sama persis.

## ⚠️ Kesalahan Umum
- Menulis `:=` di luar fungsi.
- Mencampur `int` dan `float64` dalam operasi tanpa konversi.
- Mengira variabel yang dideklarasikan tapi tidak dipakai boleh; Go akan error.
- Memakai `float64` untuk uang tanpa strategi pembulatan.

## ✅ Kuis
Tiga soal singkat menunggu di bawah.

## 📝 Ringkasan
- Pilih tipe sesuai kebutuhan, bukan asal.
- Zero value membuat kode lebih aman.
- Konversi selalu eksplisit.
- `const` untuk nilai tetap.
