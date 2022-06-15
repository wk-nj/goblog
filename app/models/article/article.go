package article

import (
	"goblog/app/models"
	"goblog/pkg/route"
	"goblog/pkg/types"
	"gorm.io/gorm"
	"time"
)

type Article struct {
	models.BaseModel
	Title, Body string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", types.Uint64ToString(a.ID))
}


