package postgres

import (
	"context"
	"odilbekqazaqov4657/my_blog_backend/models"
	log "odilbekqazaqov4657/my_blog_backend/pkg/logger"

	"github.com/google/uuid"
	"github.com/saidamir98/udevs_pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

type contentRepo struct {
	db  *pgxpool.Pool
	log log.Log
}

func NewContentRepo(db *pgxpool.Pool, log log.Log) ContentRepoI {
	return &contentRepo{db, log}
}

// create category
func (c *contentRepo) CreateCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	c.log.Debug("request in Create category")

	category.CategoryID = uuid.New() // uuid generatsiya qiladi

	query := `
		INSERT INTO
			categories(
				CategoryID
				Name      
			) VALUES (
			 	$1, $2
			)`

	_, err := c.db.Exec(ctx, query, category.CategoryID, category.Name)

	if err != nil {
		c.log.Error("error on creating category", logger.Error(err))
		return nil, err
	}

	cat, err := c.GetCategory(ctx, category.CategoryID.String())

	if err != nil {
		c.log.Error("error on getting new category", logger.Error(err))
		return nil, err
	}

	return cat, nil
}

// Get categories
func (c contentRepo) GetCategories(ctx context.Context, page, limit int32) (*models.GetCategoryListResp, error) {

	c.log.Debug("request in Create category")
	query := `
		SELECT
			*
		FROM
			categories
		LIMIT 
			$1
		OFFSET
		 $2`

	offset := (page - 1) * limit

	rows, err := c.db.Query(ctx, query, limit, offset)

	if err != nil {
		c.log.Error("error on getting all category ", logger.Error(err))
		return nil, err
	}

	defer rows.Close()

	var categories []*models.Category

	for rows.Next() {
		var category models.Category

		err := rows.Scan(
			&category.CategoryID,
			&category.Name,
			&category.CreatedAt,
		)

		if err != nil {
			c.log.Error("error on scanning category ", logger.Error(err))

			return nil, err
		}
		categories = append(categories, &category)
	}
	var count int32

	err = c.db.QueryRow(ctx, `SELECT count(*) FROM categories`).Scan(&count)

	if err != nil {
		c.log.Error("error on scanning category count", logger.Error(err))
		return nil, err
	}

	return &models.GetCategoryListResp{
		Categories: categories,
		Count:      count,
	}, nil
}

// Get category
func (c contentRepo) GetCategory(ctx context.Context, id string) (*models.Category, error) {

	c.log.Debug("request in Create category")

	var category models.Category

	query := `
		SELECT 
			* 
		FROM 
			categories 
		WHERE 
			category_id=$1`

	err := c.db.QueryRow(ctx, query, id).Scan(
		&category.CategoryID,
		&category.Name,
		&category.CreatedAt,
	)

	if err != nil {
		c.log.Error("error on getting category by id ", logger.Error(err))
		return nil, err
	}

	return &category, nil
}

func (c contentRepo) UpdateCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	return nil, nil
}
func (c contentRepo) DeleteCategory(ctx context.Context, id string) error {
	return nil
}

// sub-category
func (c contentRepo) CreateSubCategory(ctx context.Context, subcategory *models.SubCategory) (*models.SubCategory, error) {
	return nil, nil
}
func (c contentRepo) GetSubCategories(ctx context.Context, page, limit int32) ([]*models.SubCategory, error) {
	return nil, nil
}
func (c contentRepo) GetSubCategory(ctx context.Context, id string) (*models.SubCategory, error) {
	return nil, nil
}
func (c contentRepo) UpdateSubCategory(ctx context.Context, subcategory *models.SubCategory) (*models.SubCategory, error) {
	return nil, nil
}
func (c contentRepo) DeleteSubCategory(ctx context.Context, id string) error {
	return nil
}

// article
func (c contentRepo) CreateArticle(ctx context.Context, category *models.Category) (*models.Category, error) {
	return nil, nil
}
func (c contentRepo) GetArticles(ctx context.Context, page, limit int32) ([]*models.Category, error) {
	return nil, nil
}
func (c contentRepo) GetArticle(ctx context.Context, id string) (*models.Category, error) {
	return nil, nil
}
func (c contentRepo) UpdateArticle(ctx context.Context, category *models.Category) (*models.Category, error) {
	return nil, nil
}
func (c contentRepo) DeleteArticle(ctx context.Context, id string) error {
	return nil
}
