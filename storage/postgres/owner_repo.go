package postgres

import (
	"context"
	"odilbekqazaqov4657/my_blog_backend/models"
	log "odilbekqazaqov4657/my_blog_backend/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ownerRepo struct {
	db  *pgxpool.Pool
	log log.Log
}

func NewOwnerRepo(db *pgxpool.Pool, log log.Log) OwnerRepoI {
	return &ownerRepo{db, log}
}

func (o *ownerRepo) Login(ctx context.Context, login *models.LoginOwn) (*models.Owner, error) {
	return nil, nil
}
