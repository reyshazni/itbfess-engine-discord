package Entity

import (
	"time"
)

type User struct {
	AuthorID   string `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	IsVerified bool
}
