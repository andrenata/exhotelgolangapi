package category

type CategoryFormatter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func FormatCategory(category Category) CategoryFormatter {
	formatter := CategoryFormatter{
		ID:   category.ID,
		Name: category.Name,
		Slug: category.Slug,
	}

	return formatter
}

func FormatCategories(category []Category) []CategoryFormatter {
	CategoryFormatter := []CategoryFormatter{}
	for _, category := range category {
		formatter := FormatCategory(category)
		CategoryFormatter = append(CategoryFormatter, formatter)
	}
	return CategoryFormatter
}
