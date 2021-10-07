package settings

import "time"

type Setting struct {
	ID          int
	Title       string
	Keyword     string
	Description string
	Facebook    string
	Instagram   string
	Maps        string
	Address     string
	Phone       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
