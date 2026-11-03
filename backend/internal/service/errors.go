package service

import "net/http"

type Error struct {
	Status  int
	Code    string
	Message string
}

func (e *Error) Error() string { return e.Message }

func NewError(status int, code, message string) *Error {
	return &Error{Status: status, Code: code, Message: message}
}

var (
	ErrInvalidCredentials = NewError(http.StatusUnauthorized, "invalid_credentials", "Email atau password salah.")
	ErrEmailTaken         = NewError(http.StatusConflict, "email_taken", "Email sudah terdaftar.")
	ErrUnauthorized       = NewError(http.StatusUnauthorized, "unauthorized", "Sesi tidak valid atau sudah berakhir.")
	ErrNotFound           = NewError(http.StatusNotFound, "not_found", "Data tidak ditemukan.")
	ErrBadRequest         = NewError(http.StatusBadRequest, "bad_request", "Permintaan tidak valid.")
	ErrLocked             = NewError(http.StatusForbidden, "lesson_locked", "Selesaikan pelajaran sebelumnya terlebih dahulu.")
	ErrRunner             = NewError(http.StatusBadGateway, "runner_error", "Gagal menjalankan kode. Coba lagi.")
)
