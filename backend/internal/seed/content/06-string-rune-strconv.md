---
slug: string-rune-strconv
title: String, Rune & Package strconv
summary: Bedakan byte dan rune, iterasi teks Unicode, dan konversi angka dengan aman.
minutes: 13
exercise:
  title: Hitung huruf vokal
  difficulty: sedang
  xp: 25
  instructions: |
    Hitung jumlah huruf vokal (a, i, u, e, o) pada teks, termasuk huruf kapital.

    Output harus:
    vokal=5

    Gunakan range agar aman untuk Unicode.
  starter: |
    package main

    import "fmt"

    func hitungVokal(teks string) int {
        // TODO: iterasi rune dan hitung vokal
        return 0
    }

    func main() {
        fmt.Printf("vokal=%d\n", hitungVokal("Belajar Go Itu Asyik"))
    }
  solution: |
    package main

    import "fmt"

    func hitungVokal(teks string) int {
        total := 0
        for _, r := range teks {
            switch r {
            case 'a', 'i', 'u', 'e', 'o', 'A', 'I', 'U', 'E', 'O':
                total++
            }
        }
        return total
    }

    func main() {
        fmt.Printf("vokal=%d\n", hitungVokal("Belajar Go Itu Asyik"))
    }
  hints:
    - range pada string menghasilkan indeks byte dan rune.
    - Gunakan tanda kutip tunggal untuk rune, misalnya 'a'.
  tests:
    - expected: |
        vokal=5
quizzes:
  - question: Apa tipe data untuk satu karakter Unicode di Go?
    explanation: Rune adalah alias int32 dan merepresentasikan satu code point Unicode.
    options:
      - text: rune
        correct: true
      - text: byte
      - text: char
      - text: string
  - question: Apa hasil len("é") jika é memakai dua byte UTF-8?
    explanation: len string menghitung byte, bukan karakter. Panjang byte-nya 2.
    options:
      - text: "2"
        correct: true
      - text: "1"
      - text: "4"
      - text: "0"
  - question: Fungsi apa untuk mengubah string angka menjadi int?
    explanation: strconv.Atoi mengembalikan int dan error. Selalu cek error-nya.
    options:
      - text: strconv.Atoi
        correct: true
      - text: strings.ToInt
      - text: fmt.ParseInt
      - text: int.Parse
flashcards:
  - front: Apa perbedaan byte dan rune?
    back: byte = uint8, satu byte; rune = int32, satu code point Unicode.
  - front: Mengapa len("é") bisa 2?
    back: len menghitung byte. Karakter non-ASCII memakai lebih dari satu byte di UTF-8.
---

## 🎯 Tujuan
Setelah ini kamu bisa mengolah teks dengan aman, membedakan byte dan rune, serta mengonversi angka memakai `strconv`.

## ⚡ Inti dalam 1 Menit
- `string` adalah urutan byte yang immutable.
- `byte` = `uint8`; `rune` = `int32` untuk satu karakter Unicode.
- `range` pada string menghasilkan rune, aman untuk teks non-ASCII.
- `strconv.Atoi` / `Itoa` untuk konversi angka; `fmt.Sprintf` untuk format.

## 📖 Konsep
Indeks string memberi byte, bukan karakter:

```go
s := "Go"
fmt.Println(s[0])        // 71 (byte)
fmt.Printf("%c\n", s[0]) // G
```

Untuk karakter Unicode:

```go
for i, r := range "Halo é" {
	fmt.Println(i, r, string(r))
}
```

Package `strings` berisi `Contains`, `Split`, `Join`, `ReplaceAll`, `ToLower`, `TrimSpace`. Package `strconv` fokus konversi:

```go
n, err := strconv.Atoi("42")
if err != nil { ... }
teks := strconv.Itoa(n)
b, _ := strconv.ParseBool("true")
```

## 💻 Contoh
```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	raw := "  10,20,30  "
	bagian := strings.Split(strings.TrimSpace(raw), ",")
	total := 0
	for _, b := range bagian {
		n, err := strconv.Atoi(b)
		if err != nil {
			fmt.Println("bukan angka:", b)
			continue
		}
		total += n
	}
	fmt.Println("total:", total)
}
```

## ✍️ Coba Sendiri
Selesaikan latihan hitung vokal. Tambahkan huruf besar pada daftar case.

## ⚠️ Kesalahan Umum
- Memakai `s[i]` untuk karakter non-ASCII.
- Memakai `strconv.Atoi` tanpa menangani error.
- Memakai `+` berulang di loop untuk membangun string besar; gunakan `strings.Builder`.
- Membandingkan rune dengan string `'a' == "a"` (tipe berbeda).

## ✅ Kuis
Tiga soal di bawah.

## 📝 Ringkasan
- String immutable; setiap operasi menghasilkan string baru.
- Byte untuk data biner, rune untuk teks.
- `range` untuk Unicode, indeks untuk byte.
- `strconv` untuk konversi dengan error handling.
