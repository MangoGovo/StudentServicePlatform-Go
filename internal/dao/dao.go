package dao

import (
	"JH_2024_MJJ/internal/model"
	"context"
	"gorm.io/gorm"
)

type Dao struct {
	orm *gorm.DB
}

func New(orm *gorm.DB) *Dao {
	return &Dao{orm: orm}
}

type Daos interface {
	// User
	GetUserByUserName(ctx context.Context, username string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id int64) (*model.User, error)

	// User - Token
	GetToken(ctx context.Context, token string) (*model.TokenTable, error)
	GetTokenByID(ctx context.Context, id int64) (*model.TokenTable, error)
	UpdateToken(ctx context.Context, token *model.TokenTable) error
	CreateToken(ctx context.Context, token *model.TokenTable) error
	//GetUserByToken(ctx context.Context, token string) (*model.User, error)
	
	// Student
	CreateArticle(ctx context.Context, article *model.Article) error
	GetArticleList(ctx context.Context) ([]*model.Article, error)
	GetArticleByID(ctx context.Context, articleID int64) (*model.Article, error)
	DelArticle(ctx context.Context, articleID int64) error
	UpdatePost(ctx context.Context, article *model.Article) error
	QueryReport(ctx context.Context, userID int64) (error, []*model.Article)

	// Admin
	QueryUnhandledReport(ctx context.Context, userID int64) (error, []*model.Article)
	DeleteReport(ctx context.Context, article *model.Article) error
}
