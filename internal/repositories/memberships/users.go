package memberships

import (
	"context"
	"database/sql"
	"github.com/Richieid23/simple-forum/internal/models/memberships"
)

func (r *repository) GetUser(ctx context.Context, username, email string) (*memberships.UserModel, error) {
	query := `SELECT id, username, email, password, created_at, updated_at, created_by, updated_by FROM users WHERE username = ? OR email = ?`
	row := r.db.QueryRowContext(ctx, query, username, email)

	var user memberships.UserModel
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.CreatedBy, &user.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *repository) CreateUser(ctx context.Context, user memberships.UserModel) error {
	query := `INSERT INTO users (username, email, password, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.CreatedBy, user.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}
