package store

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	GitHubId    int64     `json:"git_hub_id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	AvatarURL   string    `json:"avatar_url"`
	AccessToken string    `json:"-"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserStore interface {
	Create(ctx context.Context, gitHubID int64, username, email, avatarURL, accessToken string) (*User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetByGitHubID(ctx context.Context, gitHubID int64) (*User, error)
	UpdateAccessToken(ctx context.Context, id uuid.UUID, accessToken string) error
}

type PGUserStore struct {
	pool *pgxpool.Pool
}

var (
	ErrDuplicatedGitHubID error = errors.New("github id has been used")
)

func (s *PGUserStore) Create(ctx context.Context, gitHubID int64, username, email, avatarURL, accessToken string) (*User, error) {
	const q = `
		INSERT into users (github_id, username, email, avatar_url, access_token, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $6)
		RETURNING id, github_id, username, email, avatar_url, access_token, created_at, updated_at
`

	now := time.Now().UTC()
	var user User

	err := s.pool.QueryRow(ctx, q, gitHubID, username, email, avatarURL, accessToken, now).Scan(
		&user.ID,
		&user.GitHubId,
		&user.Username,
		&user.Email,
		&user.AvatarURL,
		&user.AccessToken,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, ErrDuplicatedGitHubID
		}
		return nil, err
	}
	return &user, nil

}

func (s *PGUserStore) GetByID(ctx context.Context, id uuid.UUID) (*User, error) {
	const q = `
	SELECT id, github_id, username, email, avatar_url, access_token, created_at, updated_at 
	FROM users WHERE id = $1;
`
	var user User
	err := s.pool.QueryRow(ctx, q, id).Scan(
		&user.ID,
		&user.GitHubId,
		&user.Username,
		&user.Email,
		&user.AvatarURL,
		&user.AccessToken,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("not found")
		}
		return nil, err
	}
	return &user, nil
}

func (s *PGUserStore) GetByGitHubID(ctx context.Context, gitHubID int64) (*User, error) {
	const q = `
        SELECT id, github_id, username, email, avatar_url, access_token, created_at, updated_at 
        FROM users WHERE github_id = $1
    `

	var user User
	err := s.pool.QueryRow(ctx, q, gitHubID).Scan(
		&user.ID,
		&user.GitHubId,
		&user.Username,
		&user.Email,
		&user.AvatarURL,
		&user.AccessToken,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (s *PGUserStore) UpdateAccessToken(ctx context.Context, id uuid.UUID, accessToken string) error {
	const q = `
        UPDATE users 
        SET access_token = $2, updated_at = $3
        WHERE id = $1
    `

	_, err := s.pool.Exec(ctx, q, id, accessToken, time.Now().UTC())
	return err
}

func NewPGUserStore(pool *pgxpool.Pool) *PGUserStore {
	return &PGUserStore{pool: pool}
}

var _ UserStore = (*PGUserStore)(nil)
