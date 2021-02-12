package cinode

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	baseURL         string = "https://api.cinode.com"
	apiVersion      string = "v0.1"
	tokenURL        string = "https://api.cinode.com/token"
	tokenRefreshURL string = "https://api.cinode.com/token/refresh"
)

// Client is...
type Client struct {
	BaseURL    string
	CompanyID  int32
	HTTPClient *http.Client
	Auth       *Auth
}

// Auth struct holds the needed info for authentication
// It's also used to store the fetched auth token
type Auth struct {
	AccessID     string
	AccessSecret string
	*Token
}

// Token stores the fetched authentication token and refresh token
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// NewClient is ...
func NewClient(accessID string, accessSecret string, companyID int32) (*Client, error) {
	if accessID == "" || accessSecret == "" || companyID == 0 {
		return nil, errors.New("Access ID, Access Secret or CompanyID missing")
	}

	c := &Client{
		BaseURL:   baseURL,
		CompanyID: companyID,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
		Auth: &Auth{
			AccessID:     accessID,
			AccessSecret: accessSecret,
			Token: &Token{
				AccessToken:  "1qaz2wsx",
				RefreshToken: "3edc4rfv",
			},
		},
	}
	err := c.getToken()
	if err != nil {
		return nil, fmt.Errorf("Could not initialise client.\n %d", err)
	}

	return c, nil
}
