---
slug: fungsi
title: Fungsi, Multiple Return, Variadic & defer
summary: Bikin fungsi kecil yang jelas, kembalikan beberapa nilai, dan pahami defer.
minutes: 14
exercise:
  title: Min dan Maks
  difficulty: sedang
  xp: 25
  instructions: |
    Lengkapi fungsi minMax yang mengembalikan nilai terkecil dan terbesar.

    Output harus:
    min=2 maks=9
  starter: |
    package main

    import "fmt"

    func minMax(angka ...int) (int, int) {
        // TODO: kembalikan min dan max
        return 0, 0
    }

    func main() {
        min, maks := minMax(4, 9, 2, 7)
        fmt.Printf("min=%d maks=%d\n", min, maks)
    }
  solution: |
    package main

    import "fmt"

    func minMax(angka ...int) (int, int) {
        min, maks := angka[0], angka[0]
        for _, n := range angka[1:] {
            if n < min {
                min = n
            }
            if n > maks {
                maks = n
            }
        }
        return min, maks
    }

    func main() {
        min, maks := minMax(4, 9, 2, 7)
        fmt.Printf("min=%d maks=%d\n", min, maks)
    }
  hints:
    - Parameter variadic ...int membuat angka bertipe []int.
    - Inisialisasi min dan maks dengan elemen pertama.
  tests:
    - expected: |
        min=2 maks=9
quizzes:
  - question: Apa yang terjadi pada argumen fungsi Go?
    explanation: Go meneruskan argumen dengan value. Struct besar sebaiknya lewat pointer agar tidak menyalin data.
    options:
      - text: Disalin (pass by value)
        correct: true
      - text: Selalu pass by reference
      - text: Disalin hanya untuk string
      - text: Tidak disalin sama sekali
  - question: Kapan defer dijalankan?
    explanation: defer dijadwalkan dan dijalankan saat fungsi yang membungkusnya selesai, termasuk saat panic.
    options:
      - text: Saat fungsi selesai, dalam urutan LIFO
        correct: true
      - text: Segera setelah ditulis
      - text: Sebelum fungsi dimulai
      - text: Hanya jika terjadi error
  - question: Apa kegunaan multiple return value?
    explanation: Pola paling umum adalah mengembalikan hasil dan error, misalnya value, err := kerja().
    options:
      - text: Mengembalikan hasil dan error sekaligus
        correct: true
      - text: Mengembalikan dua tipe sekaligus
      - text: Membuat fungsi lebih cepat
      - text: Menggantikan struct
flashcards:
  - front: Apa arti ... pada parameter fungsi?
    back: Parameter variadic; nilainya dikumpulkan menjadi slice.
  - front: Urutan eksekusi beberapa defer?
    back: LIFO — defer terakhir dijalankan paling awal saat fungsi selesai.
---

## 🎯 Tujuan
Setelah ini kamu bisa menulis fungsi dengan signature jelas, mengembalikan lebih dari satu nilai, memakai variadic, dan memakai `defer` untuk pembersihan.

## ⚡ Inti dalam 1 Menit
- Fungsi dideklarasikan dengan `func nama(param tipe) tipeHasil`.
- Go mendukung multiple return, biasanya `(hasil, error)`.
- `...int` membuat parameter variadic dan di dalam fungsi menjadi `[]int`.
- `defer` menjalankan pembersihan saat fungsi selesai, urutan LIFO.

## 📖 Konsep
Fungsi adalah unit terkecil yang bisa dites. Jaga satu fungsi mengerjakan satu hal.

```go
func bagi(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("pembagi tidak boleh nol")
	}
	return a / b, nil
}

hitung, err := bagi(10, 2)
```

Closure adalah fungsi yang menangkap variabel di sekitarnya. Berguna untuk membuat fungsi dengan konfigurasi.

```go
func pengali(faktor int) func(int) int {
	return func(n int) int { return n * faktor }
}
kaliDua := pengali(2)
```

`defer` biasa dipakai untuk menutup file atau melepas lock:

```go
f, err := os.Open("data.txt")
if err != nil { return err }
defer f.Close()
```

## 💻 Contoh
```go
package main

import "fmt"

func jumlah(angka ...int) int {
	total := 0
	for _, n := range angka {
		total += n
	}
	return total
}

func main() {
	defer fmt.Println("selesai")
	fmt.Println(jumlah(1, 2, 3))
}
```

## ✍️ Coba Sendiri
Selesaikan `minMax`. Uji juga dengan angka negatif setelah latihan lolos.

## ⚠️ Kesalahan Umum
- Mengira slice dikirim by reference. Header slice disalin, array di belakangnya tetap sama.
- Lupa menangani error kedua dari multiple return.
- Memakai defer di dalam loop besar tanpa fungsi pembungkus.
- Nama hasil yang tidak jelas seperti `a, b int`.

## ✅ Kuis
Tiga soal di bawah.

## 📝 Ringkasan
- Signature yang jelas adalah dokumentasi pertama.
- Kembalikan `(nilai, error)` untuk operasi yang bisa gagal.
- Variadic memudahkan API.
- `defer` menjaga cleanup tidak terlupa.
