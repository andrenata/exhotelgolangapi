package pages

type PageFormatter struct {
	ID          int    `json:"id"`
	Name       string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

func FormatPage(page Page) PageFormatter {
	formatter := PageFormatter{
		ID:          page.ID,
		Name:       page.Name,
		Slug:        page.Slug,
		Description: page.Description,
	}
	return formatter
}

func FormatPages(pages []Page) []PageFormatter {

	PageFormatter := []PageFormatter{}

	for _, page := range pages {
		formatter := FormatPage(page)
		PageFormatter = append(PageFormatter, formatter)
	}

	return PageFormatter
}
