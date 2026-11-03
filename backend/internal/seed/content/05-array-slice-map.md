---
slug: array-slice-map
title: Array, Slice & Map
summary: Pahami panjang vs kapasitas, append, dan cara aman memakai map.
minutes: 15
exercise:
  title: Hitung frekuensi kata
  difficulty: sedang
  xp: 25
  instructions: |
    Lengkapi fungsi frekuensi untuk menghitung kemunculan setiap kata.

    Output harus:
    go: 2
    belajar: 2
    itu: 1
    menyenangkan: 1

    Urutan output mengikuti urutan kemunculan pertama.
  starter: |
    package main

    import "fmt"

    func frekuensi(kata []string) map[string]int {
        // TODO: hitung frekuensi
        return nil
    }

    func main() {
        data := []string{"go", "belajar", "go", "itu", "belajar", "menyenangkan"}
        for _, k := range []string{"go", "belajar", "itu", "menyenangkan"} {
            fmt.Printf("%s: %d\n", k, frekuensi(data)[k])
        }
    }
  solution: |
    package main

    import "fmt"

    func frekuensi(kata []string) map[string]int {
        hasil := map[string]int{}
        for _, k := range kata {
            hasil[k]++
        }
        return hasil
    }

    func main() {
        data := []string{"go", "belajar", "go", "itu", "belajar", "menyenangkan"}
        for _, k := range []string{"go", "belajar", "itu", "menyenangkan"} {
            fmt.Printf("%s: %d\n", k, frekuensi(data)[k])
        }
    }
  hints:
    - Map harus dibuat dulu dengan make atau literal map[string]int{}.
    - hasil[k]++ otomatis memakai zero value 0 untuk key baru.
  tests:
    - expected: |
        go: 2
        belajar: 2
        itu: 1
        menyenangkan: 1
quizzes:
  - question: Apa perbedaan len dan cap pada slice?
    explanation: len adalah jumlah elemen saat ini; cap adalah kapasitas sebelum slice perlu realokasi.
    options:
      - text: len jumlah elemen, cap kapasitas sebelum realloc
        correct: true
      - text: Keduanya sama saja
      - text: len kapasitas, cap jumlah elemen
      - text: cap hanya ada pada array
  - question: Apa yang terjadi jika append melebihi kapasitas slice?
    explanation: Runtime mengalokasikan array baru yang lebih besar lalu menyalin elemen lama.
    options:
      - text: Slice dialokasikan ulang dengan kapasitas lebih besar
        correct: true
      - text: Program panic
      - text: Elemen lama hilang
      - text: Terjadi error compile
  - question: Bagaimana mengecek apakah key ada di map?
    explanation: Gunakan comma-ok, misalnya v, ok := m["k"]. ok bernilai false jika key tidak ada.
    options:
      - text: v, ok := m["k"]
        correct: true
      - text: m.exists("k")
      - text: "m[\"k\"] != nil"
      - text: len(m["k"]) > 0
flashcards:
  - front: Apa zero value slice dan map?
    back: Keduanya nil. Slice nil bisa di-append; map nil tidak bisa ditulis tanpa make.
  - front: Apa arti urutan map di Go?
    back: Tidak dijamin. Urutkan key secara eksplisit jika butuh output stabil.
---

## 🎯 Tujuan
Setelah ini kamu bisa memilih array, slice, atau map dan menghindari bug klasik seperti map nil dan urutan acak.

## ⚡ Inti dalam 1 Menit
- Array berukuran tetap; slice adalah view dinamis di atas array.
- `append` bisa mengalokasikan ulang; selalu pakai hasil kembaliannya.
- Map menyimpan key-value dan tidak menjamin urutan iterasi.
- Map nil tidak bisa ditulis; buat dengan `make` atau literal.

## 📖 Konsep
Slice terdiri dari pointer ke array, `len`, dan `cap`.

```go
angka := []int{1, 2, 3}      // len=3 cap=3
angka = append(angka, 4)      // bisa realloc
bagian := angka[1:3]          // [2 3]
```

Map:

```go
stok := map[string]int{"go": 3}
stok["gin"] = 2
v, ok := stok["echo"]
if !ok {
	fmt.Println("tidak ada")
}
delete(stok, "go")
```

Iterasi map acak secara desain. Untuk output stabil, kumpulkan key lalu `sort.Strings`.

## 💻 Contoh
```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	stok := map[string]int{"mangga": 4, "apel": 2, "jeruk": 7}
	keys := make([]string, 0, len(stok))
	for k := range stok {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s=%d\n", k, stok[k])
	}
}
```

## ✍️ Coba Sendiri
Selesaikan latihan frekuensi. Perhatikan bahwa `hasil[k]++` aman untuk key baru.

## ⚠️ Kesalahan Umum
- Lupa menulis `hasil = append(hasil, x)`.
- Menulis ke map nil sehingga panic.
- Mengandalkan urutan iterasi map.
- Menyimpan slice besar di map padahal cukup ID.

## ✅ Kuis
Tiga soal di bawah.

## 📝 Ringkasan
- Array jarang dipakai langsung, slice lebih umum.
- `append` mengembalikan slice baru.
- Map cepat untuk lookup, tetapi tidak berurutan.
- Selalu buat map sebelum menulis.
