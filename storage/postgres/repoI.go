package postgres

import (
	"context"
	"odilbekqazaqov4657/my_blog_backend/models"
)

type ContentRepoI interface {

	//  category
	CreateCategory(ctx context.Context, category *models.Category) (*models.Category, error)
	GetCategories(ctx context.Context, page, limit int32) (*models.GetCategoryListResp, error)
	GetCategory(ctx context.Context, id string) (*models.Category, error)
	UpdateCategory(ctx context.Context, category *models.Category) (*models.Category, error)
	DeleteCategory(ctx context.Context, id string) error

	//  sub-category
	CreateSubCategory(ctx context.Context, subcategory *models.SubCategory) (*models.SubCategory, error)
	GetSubCategories(ctx context.Context, page, limit int32) ([]*models.SubCategory, error)
	GetSubCategory(ctx context.Context, id string) (*models.SubCategory, error)
	UpdateSubCategory(ctx context.Context, subcategory *models.SubCategory) (*models.SubCategory, error)
	DeleteSubCategory(ctx context.Context, id string) error

	//  article
	CreateArticle(ctx context.Context, category *models.Category) (*models.Category, error)
	GetArticles(ctx context.Context, page, limit int32) ([]*models.Category, error)
	GetArticle(ctx context.Context, id string) (*models.Category, error)
	UpdateArticle(ctx context.Context, category *models.Category) (*models.Category, error)
	DeleteArticle(ctx context.Context, id string) error
}

type OwnerRepoI interface {
	Login(ctx context.Context, login *models.LoginOwn) (*models.Owner, error)
}

type CommonRepoI interface {
	CheckIsExists(ctx context.Context, req *models.Common) (bool, error)
}
