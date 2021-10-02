package pages

type CreatePageInput struct {
	Title       string `json:"title" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Description string `json:"description"`
}

type UpdatePageInput struct {
	ID          int    `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Description string `json:"description"`
}

type PageIdInput struct {
	ID int `json:"id" binding:"required"`
}

type PageSlugInput struct {
	Slug string `json:"slug" binding:"required"`
}
