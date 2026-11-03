---
slug: error-handling
title: Error Handling, Wrapping & Custom Error
summary: Tangani error seperti data, bungkus konteks, dan buat error yang bisa diperiksa.
minutes: 15
exercise:
  title: Pembagian aman
  difficulty: sedang
  xp: 25
  instructions: |
    Lengkapi fungsi bagi supaya mengembalikan error saat pembagi nol.

    Output harus:
    5 <nil>
    pembagi tidak boleh nol

    Gunakan errors.New untuk error.
  starter: |
    package main

    import (
        "errors"
        "fmt"
    )

    func bagi(a, b int) (int, error) {
        // TODO: tangani b == 0
        return 0, nil
    }

    func main() {
        hasil, err := bagi(10, 2)
        fmt.Println(hasil, err)
        _, err = bagi(1, 0)
        fmt.Println(err)
    }
  solution: |
    package main

    import (
        "errors"
        "fmt"
    )

    func bagi(a, b int) (int, error) {
        if b == 0 {
            return 0, errors.New("pembagi tidak boleh nol")
        }
        return a / b, nil
    }

    func main() {
        hasil, err := bagi(10, 2)
        fmt.Println(hasil, err)
        _, err = bagi(1, 0)
        fmt.Println(err)
    }
  hints:
    - Error adalah nilai; return nil jika sukses.
    - fmt.Println pada error memanggil method Error().
  tests:
    - expected: |
        5 <nil>
        pembagi tidak boleh nol
quizzes:
  - question: Apa cara idiomatik menangani error di Go?
    explanation: Error dikembalikan sebagai nilai dan diperiksa pemanggil; Go tidak memakai exception.
    options:
      - text: Kembalikan error dan periksa di pemanggil
        correct: true
      - text: Gunakan try/catch
      - text: Abaikan error agar kode ringkas
      - text: Selalu panic
  - question: Apa fungsi fmt.Errorf dengan %w?
    explanation: "%w membungkus error asli sehingga errors.Is dan errors.As tetap bisa memeriksanya."
    options:
      - text: Membungkus error agar bisa diperiksa dengan errors.Is
        correct: true
      - text: Mencetak error ke layar
      - text: Menghapus error
      - text: Mengubah error menjadi string permanen
  - question: Kapan panic sebaiknya dipakai?
    explanation: Panic untuk bug programmer yang tidak bisa dipulihkan, bukan untuk kegagalan normal seperti input salah.
    options:
      - text: Hanya untuk kondisi yang benar-benar tidak bisa dipulihkan
        correct: true
      - text: Setiap kali fungsi gagal
      - text: Untuk validasi input pengguna
      - text: Saat koneksi database putus
flashcards:
  - front: Apa perbedaan errors.Is dan errors.As?
    back: Is membandingkan dengan error sentinel; As mengekstrak tipe error tertentu.
  - front: Mengapa error sebaiknya tidak diabaikan?
    back: Kegagalan tersembunyi menyebabkan data korup dan bug sulit dilacak.
---

## 🎯 Tujuan
Setelah ini kamu bisa mengembalikan error dengan konteks, membungkusnya, dan membuat tipe error sendiri saat perlu.

## ⚡ Inti dalam 1 Menit
- Error adalah nilai yang mengimplementasikan `Error() string`.
- Kembalikan `nil` jika sukses; periksa error segera.
- `fmt.Errorf("...: %w", err)` membungkus error tanpa kehilangan rantai.
- `errors.Is` untuk sentinel, `errors.As` untuk tipe error.

## 📖 Konsep
Pola dasar:

```go
func bacaFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("baca %s: %w", path, err)
	}
	return data, nil
}
```

Custom error:

```go
type ValidationError struct {
	Field string
}

func (e ValidationError) Error() string {
	return "field tidak valid: " + e.Field
}
```

Pemanggil bisa memeriksa:

```go
var ve ValidationError
if errors.As(err, &ve) {
	fmt.Println("validasi gagal di", ve.Field)
}
```

Jangan menambahkan konteks berulang di setiap layer; satu atau dua lapis biasanya cukup.

## 💻 Contoh
```go
package main

import (
	"errors"
	"fmt"
)

var ErrTidakDitemukan = errors.New("tidak ditemukan")

func cari(id int) (string, error) {
	if id != 1 {
		return "", fmt.Errorf("cari user %d: %w", id, ErrTidakDitemukan)
	}
	return "Rina", nil
}

func main() {
	_, err := cari(2)
	fmt.Println(err)
	fmt.Println(errors.Is(err, ErrTidakDitemukan))
}
```

## ✍️ Coba Sendiri
Selesaikan latihan pembagian aman. Setelah lolos, ubah `errors.New` menjadi error sentinel yang bisa dicek `errors.Is`.

## ⚠️ Kesalahan Umum
- Menulis `if err != nil { }` kosong.
- Memakai `%v` padahal butuh `%w` untuk membungkus.
- Panic untuk validasi input.
- Menyembunyikan error asli dengan pesan baru tanpa `%w`.

## ✅ Kuis
Tiga soal di bawah.

## 📝 Ringkasan
- Error adalah data, bukan exception.
- Bungkus dengan `%w` untuk konteks.
- Gunakan sentinel atau tipe error sesuai kebutuhan.
- Jangan panic untuk kegagalan biasa.
