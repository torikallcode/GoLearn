package service

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"golearn/internal/model"
	"golearn/internal/store"
)

type AuthService struct {
	store  *store.Store
	tokens *TokenService
	now    func() time.Time
}

func NewAuthService(s *store.Store, tokens *TokenService) *AuthService {
	return &AuthService{store: s, tokens: tokens, now: time.Now}
}

type AuthResult struct {
	User   *model.User `json:"user"`
	Stats  model.Stats `json:"stats"`
	Tokens Tokens      `json:"tokens"`
}

func (a *AuthService) Register(name, email, password string) (*AuthResult, error) {
	name = strings.TrimSpace(name)
	email = strings.ToLower(strings.TrimSpace(email))
	if len(name) < 2 {
		return nil, NewError(http.StatusUnprocessableEntity, "validation", "Nama minimal 2 karakter.")
	}
	if !strings.Contains(email, "@") || len(email) > 190 {
		return nil, NewError(http.StatusUnprocessableEntity, "validation", "Format email tidak valid.")
	}
	if len(password) < 8 {
		return nil, NewError(http.StatusUnprocessableEntity, "validation", "Password minimal 8 karakter.")
	}
	if _, err := a.store.UserByEmail(email); err == nil {
		return nil, ErrEmailTaken
	} else if !errors.Is(err, store.ErrNotFound) {
		return nil, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user, err := a.store.CreateUser(name, email, string(hash))
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return nil, ErrEmailTaken
		}
		return nil, err
	}
	if err = a.store.EnsureStats(user.ID); err != nil {
		return nil, err
	}
	return a.issue(user)
}

func (a *AuthService) Login(email, password string) (*AuthResult, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	user, err := a.store.UserByEmail(email)
	if errors.Is(err, store.ErrNotFound) {
		return nil, ErrInvalidCredentials
	}
	if err != nil {
		return nil, err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
		return nil, ErrInvalidCredentials
	}
	return a.issue(user)
}

func (a *AuthService) issue(user *model.User) (*AuthResult, error) {
	now := a.now()
	tokens, refreshHash, err := a.tokens.Issue(user.ID, now)
	if err != nil {
		return nil, err
	}
	if err = a.store.SaveRefreshToken(user.ID, refreshHash, now.Add(a.tokens.RefreshTTL())); err != nil {
		return nil, err
	}
	stats, err := a.store.GetStats(user.ID)
	if err != nil {
		return nil, err
	}
	return &AuthResult{User: user, Stats: stats, Tokens: tokens}, nil
}

func (a *AuthService) Refresh(refreshToken string) (*AuthResult, error) {
	claims, err := a.tokens.Parse(refreshToken, "refresh")
	if err != nil {
		return nil, err
	}
	now := a.now()
	hash := HashToken(refreshToken)
	valid, err := a.store.RefreshTokenValid(claims.UserID, hash, now)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, ErrUnauthorized
	}
	user, err := a.store.UserByID(claims.UserID)
	if err != nil {
		return nil, ErrUnauthorized
	}
	if err = a.store.RevokeRefreshToken(hash, now); err != nil {
		return nil, err
	}
	return a.issue(user)
}

func (a *AuthService) Logout(refreshToken string) error {
	if strings.TrimSpace(refreshToken) == "" {
		return nil
	}
	return a.store.RevokeRefreshToken(HashToken(refreshToken), a.now())
}

func (a *AuthService) Me(userID uint64) (*model.User, model.Stats, error) {
	user, err := a.store.UserByID(userID)
	if err != nil {
		return nil, model.Stats{}, ErrUnauthorized
	}
	stats, err := a.store.GetStats(userID)
	return user, stats, err
}

func (a *AuthService) DeleteAccount(userID uint64) error {
	return a.store.DeleteUser(userID)
}
