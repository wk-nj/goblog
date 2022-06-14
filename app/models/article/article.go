package article

import (
	"goblog/pkg/route"
	"goblog/pkg/types"
	"time"
)

type Article struct {
	ID          uint64
	Title, Body string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", types.Uint64ToString(a.ID))
}


