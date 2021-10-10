package pages

type CreatePageInput struct {
	Name       string `json:"name" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Description string `json:"description"`
}

type UpdatePageInput struct {
	ID          int    `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Description string `json:"description"`
}

type PageIdInput struct {
	ID int `json:"id" binding:"required"`
}

type PageSlugInput struct {
	Slug string `json:"slug" binding:"required"`
}
