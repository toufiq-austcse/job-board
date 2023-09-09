package req

import "github.com/gin-gonic/gin"

type UpdateCompanyReqModel struct {
	Name               string `json:"name,omitempty"`
	Location           string `json:"location,omitempty"`
	LogoURL            string `json:"logo_url,omitempty"`
	WebsiteURL         string `json:"website_url,omitempty"`
	Size               string `json:"size,omitempty"`
	Industry           string `json:"industry,omitempty"`
	Established        string `json:"established,omitempty"`
	Description        string `json:"description,omitempty"`
	CultureDescription string `json:"culture_description,omitempty"`
	HiringDescription  string `json:"hiring_description,omitempty"`
}

func (model *UpdateCompanyReqModel) Validate(c *gin.Context) error {
	err := c.BindJSON(&model)
	if err != nil {
		return err
	}
	return nil
}
