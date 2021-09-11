package category

import (
	"time"
)

type Category struct {
	ID        int
	Name      string
	Slug      string
	ProductID int
	CreatedAt time.Time
	UpdatedAt time.Time
	// CategoryRelations []CategoryRelation
}

// type CategoryRelation struct {
// 	ProductID int
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }
