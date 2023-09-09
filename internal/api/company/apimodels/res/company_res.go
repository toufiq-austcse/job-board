package res

type CompanyDetailsRes struct {
	Name               string `json:"name"`
	Location           string `json:"location"`
	LogoURL            string `json:"logo_url"`
	WebsiteURL         string `json:"website_url"`
	Email              string `json:"email"`
	Size               string `json:"size"`
	Industry           string `json:"industry"`
	Established        string `json:"established"`
	Description        string `json:"description"`
	CultureDescription string `json:"culture_description"`
	HiringDescription  string `json:"hiring_description"`
	Slug               string `json:"slug"`
}
