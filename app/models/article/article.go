package article

import "time"

type Article struct {
	ID          uint64
	Title, Body string
	CreatedAt time.Time
	UpdatedAt time.Time
}


