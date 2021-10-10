package pages

import "time"

type Page struct {
	ID          int
	Name       string
	Slug        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
