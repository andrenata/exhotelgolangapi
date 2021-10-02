package pages

import "time"

type Page struct {
	ID          int
	Title       string
	Slug        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
