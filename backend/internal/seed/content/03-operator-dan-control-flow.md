---
slug: operator-dan-control-flow
title: Operator & Control Flow
summary: if, switch, for, dan operator logika untuk mengatur alur program.
minutes: 14
exercise:
  title: Klasifikasi nilai
  difficulty: sedang
  xp: 25
  instructions: |
    Lengkapi fungsi klasifikasi lalu cetak untuk tiga nilai: 92, 75, dan 50.

    Aturan:
    - nilai >= 85: "A"
    - nilai >= 70: "B"
    - selain itu: "C"

    Output tepat:
    92: A
    75: B
    50: C
  starter: |
    package main

    import "fmt"

    func klasifikasi(nilai int) string {
        // TODO: gunakan if/else if/else
        return ""
    }

    func main() {
        for _, n := range []int{92, 75, 50} {
            fmt.Printf("%d: %s\n", n, klasifikasi(n))
        }
    }
  solution: |
    package main

    import "fmt"

    func klasifikasi(nilai int) string {
        if nilai >= 85 {
            return "A"
        } else if nilai >= 70 {
            return "B"
        }
        return "C"
    }

    func main() {
        for _, n := range []int{92, 75, 50} {
            fmt.Printf("%d: %s\n", n, klasifikasi(n))
        }
    }
  hints:
    - return menghentikan fungsi, jadi else setelah return bisa dihilangkan.
    - for dengan range dipakai untuk mengiterasi slice.
  tests:
    - expected: |
        92: A
        75: B
        50: C
quizzes:
  - question: Apa bentuk perulangan di Go?
    explanation: Go hanya punya kata kunci for; varian while ditulis sebagai for kondisi.
    options:
      - text: Hanya for
        correct: true
      - text: for dan while
      - text: for, while, dan do-while
      - text: loop
  - question: Apa yang dilakukan fallthrough di dalam switch?
    explanation: fallthrough memaksa eksekusi berlanjut ke case berikutnya walaupun kondisinya tidak cocok. Jarang dibutuhkan.
    options:
      - text: Lanjut ke case berikutnya tanpa cek kondisi
        correct: true
      - text: Keluar dari switch
      - text: Mengulang case saat ini
      - text: Melempar error
  - question: Di Go, apakah if wajib memakai tanda kurung?
    explanation: Go tidak memakai kurung kurawal biasa untuk kondisi, tetapi tetap wajib memakai kurung kurawal untuk blok.
    options:
      - text: Tidak, tetapi blok wajib memakai kurung kurawal
        correct: true
      - text: Ya, selalu
      - text: Hanya jika ada else
      - text: Hanya untuk switch
flashcards:
  - front: Bagaimana menulis while di Go?
    back: Gunakan for kondisi { ... } atau for { ... } dengan break.
  - front: Apa itu short declaration di if?
    back: "if err := kerja(); err != nil { ... } — variabel hanya hidup di blok if/else."
---

## 🎯 Tujuan
Setelah ini kamu bisa mengatur alur program memakai `if`, `switch`, dan `for` tanpa membawa kebiasaan bahasa lain yang tidak perlu.

## ⚡ Inti dalam 1 Menit
- Kondisi `if` tidak memakai tanda kurung; blok tetap memakai `{}`.
- `switch` di Go tidak butuh `break` di akhir case.
- Hanya ada `for`, tetapi bisa berperan sebagai while dan loop tak terbatas.
- Operator umum: `+ - * / %`, `== != < <= > >=`, `&& || !`.

## 📖 Konsep
Go sengaja kecil. Sedikit bentuk kontrol membuat kode mudah dibaca.

```go
if skor >= 85 {
	fmt.Println("A")
} else if skor >= 70 {
	fmt.Println("B")
} else {
	fmt.Println("C")
}

switch hari {
case "Sabtu", "Minggu":
	fmt.Println("Libur")
default:
	fmt.Println("Kerja")
}

for i := 0; i < 3; i++ {
	fmt.Println(i)
}

for _, v := range []string{"a", "b"} {
	fmt.Println(v)
}
```

`if` juga bisa mendeklarasikan variabel sementara:

```go
if err := simpan(); err != nil {
	fmt.Println("gagal:", err)
}
```

## 💻 Contoh
```go
package main

import "fmt"

func main() {
	for i := 1; i <= 15; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
```

## ✍️ Coba Sendiri
Selesaikan latihan klasifikasi nilai. Ubah dulu nilai uji untuk memastikan cabang lain juga benar.

## ⚠️ Kesalahan Umum
- Memakai `while` — tidak ada di Go.
- Menulis `if (x > 0)` — valid tapi tidak idiomatik; gofmt tidak menghapusnya, jadi biasakan tanpa kurung.
- Lupa `break` di loop tak terbatas.
- Membandingkan string dan angka tanpa konversi.

## ✅ Kuis
Tiga soal tersedia di bawah.

## 📝 Ringkasan
- `if` tanpa kurung kondisi, blok wajib.
- `switch` lebih bersih untuk banyak cabang.
- `for` menutup semua kebutuhan perulangan.
- Gunakan `if` dengan inisialisasi untuk error handling ringkas.
