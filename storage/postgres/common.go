package postgres

import (
	"context"
	"fmt"
	"odilbekqazaqov4657/my_blog_backend/models"
	log "odilbekqazaqov4657/my_blog_backend/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/saidamir98/udevs_pkg/logger"
)

type CommontRepo struct {
	db  *pgxpool.Pool
	log log.Log
}

func NewCommontRepo(db *pgxpool.Pool, log log.Log) CommonRepoI {
	return &CommontRepo{db, log}
}

func (c *CommontRepo) CheckIsExists(ctx context.Context, req *models.Common) (bool, error) {
	var isExists bool

	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s='%s')", req.TableName, req.ColumnName, req.ExpValue)

	err := c.db.QueryRow(ctx, query).Scan(&isExists)

	if err != nil {
		c.log.Error("error on checking is exists", logger.Error(err))
		return false, err
	}
	return isExists, nil
}
