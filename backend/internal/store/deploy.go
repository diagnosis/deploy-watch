package store

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DeployEvent struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	RepoName      string    `json:"repo_name"`
	CommitSHA     string    `json:"commit_sha"`
	CommitMessage string    `json:"commit_message"`
	Author        string    `json:"author"`
	Branch        string    `json:"branch"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

type DeployStore interface {
	Create(ctx context.Context, userID uuid.UUID, repoName, commitSHA, commitMessage, author, branch, status string) (*DeployEvent, error)
	GetByUserID(ctx context.Context, userID uuid.UUID, limit int) ([]*DeployEvent, error)
}

type PGDeployStore struct {
	pool *pgxpool.Pool
}

func NewPGDeployStore(pool *pgxpool.Pool) *PGDeployStore {
	return &PGDeployStore{pool: pool}
}

func (s *PGDeployStore) Create(ctx context.Context, userID uuid.UUID, repoName, commitSHA, commitMessage, author, branch, status string) (*DeployEvent, error) {
	now := time.Now().UTC()
	const q = `
    INSERT INTO deploy_events (user_id, repo_name, commit_sha, commit_message, author, branch, status, created_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING id, user_id, repo_name, commit_sha, commit_message, author, branch, status, created_at
`

	var event DeployEvent
	err := s.pool.QueryRow(ctx, q, userID, repoName, commitSHA, commitMessage, author, branch, status, now).Scan(
		&event.ID,
		&event.UserID,
		&event.RepoName,
		&event.CommitSHA,
		&event.CommitMessage,
		&event.Author,
		&event.Branch,
		&event.Status,
		&event.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &event, nil

}
func (s *PGDeployStore) GetByUserID(ctx context.Context, userID uuid.UUID, limit int) ([]*DeployEvent, error) {
	if limit < 5 {
		limit = 5
	}
	const q = `
		SELECT id, user_id, repo_name, commit_sha, commit_message, author, branch, status, created_at 
		FROM deploy_events WHERE user_id = $1 ORDER BY created_at DESC LIMIT $2;
`
	deployEvents := make([]*DeployEvent, 0, limit)
	rows, err := s.pool.Query(ctx, q, userID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var deployEvent DeployEvent
		err = rows.Scan(&deployEvent.ID, &deployEvent.UserID, &deployEvent.RepoName,
			&deployEvent.CommitSHA, &deployEvent.CommitMessage, &deployEvent.Author,
			&deployEvent.Branch, &deployEvent.Status, &deployEvent.CreatedAt)
		if err != nil {
			return nil, err
		}
		deployEvents = append(deployEvents, &deployEvent)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return deployEvents, nil
}
