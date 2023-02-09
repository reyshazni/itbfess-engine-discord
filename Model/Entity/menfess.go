package Entity

import "time"

type Menfess struct {
	ID              uint      `gorm:"primarykey"`
	CreatedAt       time.Time `gorm:"type:timestamp"`
	Message         string
	AuthorID        string
	AuthorChannelID string
}
