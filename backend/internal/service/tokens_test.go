package service

import (
	"testing"
	"time"
)

func TestTokenIssueAndParse(t *testing.T) {
	svc := NewTokenService("test-secret", time.Minute, time.Hour)
	now := time.Now()

	toks, hash, err := svc.Issue(42, now)
	if err != nil {
		t.Fatalf("issue: %v", err)
	}
	if toks.AccessToken == "" || toks.RefreshToken == "" {
		t.Fatal("expected both tokens")
	}
	if hash != HashToken(toks.RefreshToken) {
		t.Fatal("refresh hash mismatch")
	}

	claims, err := svc.Parse(toks.AccessToken, "access")
	if err != nil {
		t.Fatalf("parse access: %v", err)
	}
	if claims.UserID != 42 {
		t.Fatalf("expected user 42, got %d", claims.UserID)
	}

	if _, err := svc.Parse(toks.AccessToken, "refresh"); err == nil {
		t.Fatal("access token must not parse as refresh")
	}
}

func TestTokenRejectsWrongSecret(t *testing.T) {
	issuer := NewTokenService("secret-a", time.Minute, time.Hour)
	other := NewTokenService("secret-b", time.Minute, time.Hour)
	toks, _, err := issuer.Issue(1, time.Now())
	if err != nil {
		t.Fatalf("issue: %v", err)
	}
	if _, err := other.Parse(toks.AccessToken, "access"); err == nil {
		t.Fatal("token signed with another secret must be rejected")
	}
}
