package res

type JobRes struct {
	ID          int           `json:"id"`
	Title       string        `json:"title"`
	ApplyTo     string        `json:"apply_to"`
	Description string        `json:"description"`
	Taxonomies  []JobTaxonomy `json:"taxonomies"`
}

type JobTaxonomy struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Slug  string `json:"slug"`
}
