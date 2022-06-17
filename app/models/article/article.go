package article

import (
	"goblog/app/models"
	"goblog/app/models/user"
	"goblog/pkg/route"
	"goblog/pkg/types"
	"gorm.io/gorm"
)

type Article struct {
	models.BaseModel
	Title string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body string `gorm:"type:longtext;not null;" valid:"body"`
	UserID uint64 `gorm:"not null;index"`
	User   user.User
	DeletedAt gorm.DeletedAt
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", types.Uint64ToString(a.ID))
}

func (a Article) CreatedAtDate()  string{
	return a.CreatedAt.Format("2006-01-02")
}


