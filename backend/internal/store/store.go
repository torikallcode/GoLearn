package store

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"golearn/internal/model"
)

var ErrNotFound = errors.New("not found")

type Store struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Store { return &Store{db: db} }

// ---- users ----

func (s *Store) CreateUser(name, email, hash string) (*model.User, error) {
	res, err := s.db.Exec(`INSERT INTO users (name, email, password_hash) VALUES (?, ?, ?)`, name, email, hash)
	if err != nil {
		return nil, err
	}
	id, _ := res.LastInsertId()
	return s.UserByID(uint64(id))
}

func (s *Store) UserByEmail(email string) (*model.User, error) {
	var u model.User
	err := s.db.Get(&u, `SELECT * FROM users WHERE email = ?`, email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	return &u, err
}

func (s *Store) UserByID(id uint64) (*model.User, error) {
	var u model.User
	err := s.db.Get(&u, `SELECT * FROM users WHERE id = ?`, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	return &u, err
}

func (s *Store) DeleteUser(id uint64) error {
	_, err := s.db.Exec(`DELETE FROM users WHERE id = ?`, id)
	return err
}

// ---- refresh tokens ----

func (s *Store) SaveRefreshToken(userID uint64, hash string, expires time.Time) error {
	_, err := s.db.Exec(`INSERT INTO refresh_tokens (user_id, token_hash, expires_at) VALUES (?, ?, ?)`, userID, hash, expires)
	return err
}

func (s *Store) RefreshTokenValid(userID uint64, hash string, now time.Time) (bool, error) {
	var n int
	err := s.db.Get(&n, `SELECT COUNT(*) FROM refresh_tokens
		WHERE user_id = ? AND token_hash = ? AND revoked_at IS NULL AND expires_at > ?`, userID, hash, now)
	return n > 0, err
}

func (s *Store) RevokeRefreshToken(hash string, now time.Time) error {
	_, err := s.db.Exec(`UPDATE refresh_tokens SET revoked_at = ? WHERE token_hash = ? AND revoked_at IS NULL`, now, hash)
	return err
}

func (s *Store) RevokeAllRefreshTokens(userID uint64, now time.Time) error {
	_, err := s.db.Exec(`UPDATE refresh_tokens SET revoked_at = ? WHERE user_id = ? AND revoked_at IS NULL`, now, userID)
	return err
}

// ---- stats ----

func (s *Store) EnsureStats(userID uint64) error {
	_, err := s.db.Exec(`INSERT IGNORE INTO user_stats (user_id) VALUES (?)`, userID)
	return err
}

func (s *Store) GetStats(userID uint64) (model.Stats, error) {
	var st model.Stats
	err := s.db.Get(&st, `SELECT * FROM user_stats WHERE user_id = ?`, userID)
	if errors.Is(err, sql.ErrNoRows) {
		if err = s.EnsureStats(userID); err != nil {
			return st, err
		}
		err = s.db.Get(&st, `SELECT * FROM user_stats WHERE user_id = ?`, userID)
	}
	return st, err
}

// Award adds XP and updates the daily streak atomically.
func (s *Store) Award(userID uint64, xp int, now time.Time) (model.Stats, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return model.Stats{}, err
	}
	defer tx.Rollback()

	if _, err = tx.Exec(`INSERT IGNORE INTO user_stats (user_id) VALUES (?)`, userID); err != nil {
		return model.Stats{}, err
	}
	var st model.Stats
	if err = tx.Get(&st, `SELECT * FROM user_stats WHERE user_id = ? FOR UPDATE`, userID); err != nil {
		return model.Stats{}, err
	}
	today := now.UTC().Truncate(24 * time.Hour)
	if st.LastActiveDate == nil {
		st.CurrentStreak = 1
	} else {
		last := st.LastActiveDate.UTC().Truncate(24 * time.Hour)
		diff := int(today.Sub(last).Hours() / 24)
		switch {
		case diff == 0:
			// already active today
		case diff == 1:
			st.CurrentStreak++
		default:
			st.CurrentStreak = 1
		}
	}
	if st.CurrentStreak > st.LongestStreak {
		st.LongestStreak = st.CurrentStreak
	}
	st.XP += xp
	st.Level = st.XP/100 + 1
	st.LastActiveDate = &today
	_, err = tx.Exec(`UPDATE user_stats SET xp=?, level=?, current_streak=?, longest_streak=?, last_active_date=? WHERE user_id=?`,
		st.XP, st.Level, st.CurrentStreak, st.LongestStreak, today, userID)
	if err != nil {
		return model.Stats{}, err
	}
	if err = tx.Commit(); err != nil {
		return model.Stats{}, err
	}
	return st, nil
}

func normalizeDate(t time.Time) string { return t.UTC().Format("2006-01-02") }

func (s *Store) AwardBadge(userID uint64, code string) error {
	_, err := s.db.Exec(`INSERT IGNORE INTO user_badges (user_id, badge_id)
		SELECT ?, id FROM badges WHERE code = ?`, userID, code)
	return err
}

func (s *Store) Badges(userID uint64) ([]map[string]any, error) {
	rows, err := s.db.Queryx(`SELECT b.code, b.name, b.description, ub.earned_at
		FROM user_badges ub JOIN badges b ON b.id = ub.badge_id
		WHERE ub.user_id = ? ORDER BY ub.earned_at DESC`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []map[string]any{}
	for rows.Next() {
		row := map[string]any{}
		if err := rows.MapScan(row); err != nil {
			return nil, err
		}
		for k, v := range row {
			switch b := v.(type) {
			case []byte:
				row[k] = string(b)
			}
		}
		out = append(out, row)
	}
	return out, rows.Err()
}

func nullTimePtr(t *time.Time) any {
	if t == nil {
		return nil
	}
	return *t
}

func (s *Store) Ping() error {
	var one int
	if err := s.db.Get(&one, `SELECT 1`); err != nil {
		return fmt.Errorf("db ping: %w", err)
	}
	return nil
}
