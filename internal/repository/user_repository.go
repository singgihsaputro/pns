package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/example/pns/internal/model"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) Create(ctx context.Context, email, passwordHash string) (model.User, error) {
	sql := `insert into users (email, password_hash) values ($1, $2) returning id, created_at`
	var u model.User
	u.Email = email
	u.PasswordHash = passwordHash
	if err := r.pool.QueryRow(ctx, sql, email, passwordHash).Scan(&u.ID, &u.CreatedAt); err != nil {
		return model.User{}, err
	}
	return u, nil
}

func (r *UserRepository) Update(ctx context.Context, id, email, passwordHash string) error {
	if passwordHash == "" {
		_, err := r.pool.Exec(ctx, `update users set email=$1 where id=$2`, email, id)
		return err
	}
	_, err := r.pool.Exec(ctx, `update users set email=$1, password_hash=$2 where id=$3`, email, passwordHash, id)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx, `delete from users where id=$1`, id)
	return err
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (model.User, error) {
	var u model.User
	err := r.pool.QueryRow(ctx, `select id, email, password_hash, created_at from users where email=$1`, email).Scan(
		&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return model.User{}, err
	}
	return u, err
}

func (r *UserRepository) Get(ctx context.Context, id string) (model.User, error) {
	var u model.User
	err := r.pool.QueryRow(ctx, `select id, email, password_hash, created_at from users where id=$1`, id).Scan(
		&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt,
	)
	return u, err
}

func (r *UserRepository) List(ctx context.Context) ([]model.User, error) {
	rows, err := r.pool.Query(ctx, `select id, email, password_hash, created_at from users order by created_at desc`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

