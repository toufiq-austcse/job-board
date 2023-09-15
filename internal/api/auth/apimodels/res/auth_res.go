package res

type SignUpResModel struct {
	Token       Token       `json:"token"`
	CompanyInfo CompanyInfo `json:"company_info"`
}

type LoginResModel struct {
	Token       Token       `json:"token"`
	CompanyInfo CompanyInfo `json:"company_info"`
}

type CompanyInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	ExpireAt    int64  `json:"expire_at"`
}

type TokenVerificationRes struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	LogoUrl     string `json:"logo_url"`
	WebsiteUrl  string `json:"website_url"`
	Email       string `json:"email"`
	Size        string `json:"size"`
	Slug        string `json:"slug"`
	Industry    string `json:"industry"`
	Established string `json:"established"`
}
