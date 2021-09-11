package category

type CategoryFormatter struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	ProductID int    `json:"product_id"`
}

func FormatCategory(category Category) CategoryFormatter {
	formatter := CategoryFormatter{
		ID:        category.ID,
		Name:      category.Name,
		Slug:      category.Slug,
		ProductID: category.ProductID,
	}

	return formatter
}
