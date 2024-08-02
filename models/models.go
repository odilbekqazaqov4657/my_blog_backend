package models

import (
	"time"

	"github.com/google/uuid"
)

// Owner represents the owner table in the database
type Owner struct {
	Fullname    string `json:"fullname"`
	Password    string `json:"password"`
	Role        string `json:"role" default:"owner"`
	PhoneNumber string `json:"phone_number"`
	Gmail       string `json:"gmail"`
	Telegram    string `json:"telegram"`
	Github      string `json:"github"`
	LinkedIn    string `json:"linked_in"`
	Leetcode    string `json:"leetcode"`
	AboutMe     string `json:"about_me"`
}

type LoginOwn struct {
	PhoneNumber string `json:"phone_number"`
	Gmail       string `json:"gmail"`
	Password    string `json:"password"`
}

// Category represents the categories table in the database
type Category struct {
	CategoryID uuid.UUID `json:"category_id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at" default:"current_timestamp"`
}

type CreateCategoryReq struct {
	Name string `json:"name"`
}

type GetCategoryListResp struct {
	Categories []*Category
	Count      int32
}

// SubCategory represents the sub_categories table in the database
type SubCategory struct {
	SubCategoryID uuid.UUID `json:"sub_category_id"`
	Name          string    `json:"name"`
	CreatedAt     time.Time `json:"created_at" default:"current_timestamp"`
	CategoryID    uuid.UUID `json:"category_id"`
}

// Article represents the articles table in the database
type Article struct {
	ArticleID     uuid.UUID  `json:"article_id"`
	Title         string     `json:"title"`
	Content       string     `json:"content"`
	CreatedAt     time.Time  `json:"created_at" default:"current_timestamp"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
	CategoryID    uuid.UUID  `json:"category_id"`
	SubCategoryID uuid.UUID  `json:"subCategory_id"`
}

// Viewer represents the viewers table in the database
type Viewer struct {
	ViewerID uuid.UUID `json:"viewer_id"`
	Fullname string    `json:"fullname"`
	Username string    `json:"username"`
	Gmail    string    `json:"gmail"`
	Password string    `json:"password"`
}

var CheckViewer struct {
	Gmail string `json:"gmail"`
}

// Comment represents the comments table in the database
type Comment struct {
	CommentID uuid.UUID `json:"comment_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at" default:"current_timestamp"`
	ArticleID uuid.UUID `json:"article_id"`
	ViewerID  uuid.UUID `json:"viewer_id"`
}

type GetListReq struct {
	Page  int32 `json:"page"`
	Limit int32 `json:"limit"`
}

type Common struct {
	TableName  string `jsonn:"table_name"`
	ColumnName string `jsonn:"column_name"`
	ExpValue   any    `jsonn:"exp_value"`
}

type CheckExistsResp struct {
	IsExists bool
	Status   string
}

type OtpData struct {
	Otp   string `json:"otp"`
	Gmail string `json:"gmail"`
}
type CheckOTP struct {
	Gmail string `json:"gmail"`
	Otp   string `json:"otp"`
}
