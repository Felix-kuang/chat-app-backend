package models

import "time"

type Messages struct {
	ID        int
	ChatID    int
	Content   string
	CreatedAt time.Time
}