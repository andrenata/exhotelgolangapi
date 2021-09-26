package pages

type PageFormatter struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

func FormatPage(page Page) PageFormatter {
	formatter := PageFormatter{
		ID:          page.ID,
		Title:       page.Title,
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
