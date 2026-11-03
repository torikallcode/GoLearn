package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint64 `json:"uid"`
	Type   string `json:"typ"`
	jwt.RegisteredClaims
}

type Tokens struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	TokenType    string    `json:"token_type"`
	ExpiresAt    time.Time `json:"expires_at"`
	ExpiresIn    int       `json:"expires_in"`
}

type TokenService struct {
	secret     []byte
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewTokenService(secret string, accessTTL, refreshTTL time.Duration) *TokenService {
	return &TokenService{secret: []byte(secret), accessTTL: accessTTL, refreshTTL: refreshTTL}
}

func (t *TokenService) Issue(userID uint64, now time.Time) (Tokens, string, error) {
	access, err := t.sign(userID, "access", now, t.accessTTL)
	if err != nil {
		return Tokens{}, "", err
	}
	refresh, err := t.sign(userID, "refresh", now, t.refreshTTL)
	if err != nil {
		return Tokens{}, "", err
	}
	return Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
		TokenType:    "Bearer",
		ExpiresAt:    now.Add(t.accessTTL),
		ExpiresIn:    int(t.accessTTL.Seconds()),
	}, HashToken(refresh), nil
}

func (t *TokenService) sign(userID uint64, typ string, now time.Time, ttl time.Duration) (string, error) {
	claims := Claims{
		UserID: userID,
		Type:   typ,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "user",
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
			ID:        randomID(),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(t.secret)
}

func (t *TokenService) Parse(token, typ string) (*Claims, error) {
	var claims Claims
	parsed, err := jwt.ParseWithClaims(token, &claims, func(tok *jwt.Token) (any, error) {
		if tok.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("unexpected signing method")
		}
		return t.secret, nil
	})
	if err != nil || !parsed.Valid || claims.Type != typ {
		return nil, ErrUnauthorized
	}
	return &claims, nil
}

func (t *TokenService) RefreshTTL() time.Duration { return t.refreshTTL }

func HashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

func randomID() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return hex.EncodeToString([]byte(time.Now().String()))
	}
	return hex.EncodeToString(b)
}
