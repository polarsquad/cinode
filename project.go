package cinode

import (
	"fmt"
	"net/http"
)

/*
Project URL Paths

Projects: /v0.1/companies/{companyId}/projects
*/

// Project provides a structure for basic project information
type ProjectBaseModel struct {
	CompanyID   int32  `json:"companyId"`
	CustomerID  int32  `json:"customerId"`
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Links       []Link `json:"links"`
	SelfHref    string `json:"selfHref"`
}

// Link is sub model for example. ProjectBaseModel
type Link struct {
	Href    string   `json:"href"`
	Rel     string   `json:"rel"`
	Methods []string `json:"methods"`
}

// GetProjects gets a full list of projects
func (c *Client) GetProjects() (*[]ProjectBaseModel, error) {
	url := fmt.Sprintf("%s/%s/companies/%d/projects", c.BaseURL, apiVersion, c.CompanyID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var projectList []ProjectBaseModel

	err = c.sendRequest(req, &projectList)
	if err != nil {
		return nil, err
	}

	return &projectList, nil
}
