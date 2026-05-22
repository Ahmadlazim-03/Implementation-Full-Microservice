package persistence

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"backend/internal/auth/domain/entity"
	"backend/internal/auth/domain/valueobject"
	shareddomain "backend/internal/shared/domain"
)

// PostgresUserRepository implements the domain UserRepository port.
// The domain doesn't know Postgres exists — switching to MySQL/Mongo
// means replacing this file only.
type PostgresUserRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresUserRepository(pool *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{pool: pool}
}

func (r *PostgresUserRepository) Save(ctx context.Context, u *entity.User) error {
	const q = `
		INSERT INTO users (id, email, password_hash, name, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.pool.Exec(ctx, q,
		u.ID(),
		u.Email().String(),
		u.PasswordHash(),
		u.Name(),
		string(u.Role()),
		u.CreatedAt(),
		u.UpdatedAt(),
	)
	return err
}

func (r *PostgresUserRepository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	const q = `SELECT id, email, password_hash, name, role, created_at, updated_at FROM users WHERE id = $1`
	return r.queryOne(ctx, q, id)
}

func (r *PostgresUserRepository) FindByEmail(ctx context.Context, email valueobject.Email) (*entity.User, error) {
	const q = `SELECT id, email, password_hash, name, role, created_at, updated_at FROM users WHERE email = $1`
	return r.queryOne(ctx, q, email.String())
}

func (r *PostgresUserRepository) ExistsByEmail(ctx context.Context, email valueobject.Email) (bool, error) {
	const q = `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	var exists bool
	err := r.pool.QueryRow(ctx, q, email.String()).Scan(&exists)
	return exists, err
}

func (r *PostgresUserRepository) queryOne(ctx context.Context, q string, args ...any) (*entity.User, error) {
	var (
		id, emailStr, hash, name, role string
		createdAt, updatedAt           time.Time
	)
	err := r.pool.QueryRow(ctx, q, args...).Scan(&id, &emailStr, &hash, &name, &role, &createdAt, &updatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, shareddomain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	email, err := valueobject.NewEmail(emailStr)
	if err != nil {
		return nil, err
	}
	return entity.Hydrate(id, email, hash, name, entity.Role(role), createdAt, updatedAt), nil
}
