package res

import "time"

type JobDetailsRes struct {
	ID          int           `json:"id"`
	Title       string        `json:"title"`
	Slug        string        `json:"slug"`
	Status      *string       `json:"status"`
	ApplyTo     string        `json:"apply_to"`
	Description string        `json:"description"`
	Taxonomies  []JobTaxonomy `json:"taxonomies"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

type JobTaxonomy struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Slug  string `json:"slug"`
}

type JobCompany struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	LogoUrl  string `json:"logo_url"`
}

type JobInListJobRes struct {
	ID         int           `json:"id"`
	Title      string        `json:"title"`
	Slug       string        `json:"slug"`
	Status     *string       `json:"status"`
	Company    JobCompany    `json:"company"`
	Taxonomies []JobTaxonomy `json:"taxonomies"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
}
