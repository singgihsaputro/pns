package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/example/pns/internal/model"
)

type PromotionRepository struct {
	pool *pgxpool.Pool
}

func NewPromotionRepository(pool *pgxpool.Pool) *PromotionRepository {
	return &PromotionRepository{pool: pool}
}

func (r *PromotionRepository) Create(ctx context.Context, p model.Promotion) (model.Promotion, error) {
	sql := `insert into promotions (title, image_url, description, start_date, end_date, status)
	        values ($1,$2,$3,$4,$5,$6)
	        returning id, created_at`
	if err := r.pool.QueryRow(ctx, sql, p.Title, p.ImageURL, p.Description, p.StartDate, p.EndDate, p.Status).
		Scan(&p.ID, &p.CreatedAt); err != nil {
		return model.Promotion{}, err
	}
	return p, nil
}

func (r *PromotionRepository) Update(ctx context.Context, p model.Promotion) error {
	_, err := r.pool.Exec(ctx, `update promotions set title=$1, image_url=$2, description=$3, start_date=$4, end_date=$5, status=$6 where id=$7`,
		p.Title, p.ImageURL, p.Description, p.StartDate, p.EndDate, p.Status, p.ID)
	return err
}

func (r *PromotionRepository) SetStatus(ctx context.Context, id, status string) error {
	_, err := r.pool.Exec(ctx, `update promotions set status=$1 where id=$2`, status, id)
	return err
}

func (r *PromotionRepository) Get(ctx context.Context, id string) (model.Promotion, error) {
	var p model.Promotion
	err := r.pool.QueryRow(ctx, `select id, title, image_url, description, start_date, end_date, status, created_at from promotions where id=$1`, id).
		Scan(&p.ID, &p.Title, &p.ImageURL, &p.Description, &p.StartDate, &p.EndDate, &p.Status, &p.CreatedAt)
	return p, err
}

func (r *PromotionRepository) List(ctx context.Context) ([]model.Promotion, error) {
	rows, err := r.pool.Query(ctx, `select id, title, image_url, description, start_date, end_date, status, created_at from promotions order by created_at desc`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var promos []model.Promotion
	for rows.Next() {
		var p model.Promotion
		if err := rows.Scan(&p.ID, &p.Title, &p.ImageURL, &p.Description, &p.StartDate, &p.EndDate, &p.Status, &p.CreatedAt); err != nil {
			return nil, err
		}
		promos = append(promos, p)
	}
	return promos, rows.Err()
}

