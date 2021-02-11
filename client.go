package cinode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type dataResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func (c *Client) getToken() error {
	req, err := http.NewRequest("GET", tokenURL, nil)

	if err != nil {
		log.Fatal(err)
		return err
	}

	req.SetBasicAuth(c.Auth.AccessID, c.Auth.AccessSecret)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&c.Auth.AccessToken); err != nil {
		return err
	}

	return nil
}

// RefreshToken is used to refresh expired access token
func (c *Client) RefreshToken() error {
	reqBody, err := json.Marshal(map[string]string{
		"refreshToken": c.Auth.RefreshToken,
	})

	resp, err := c.HTTPClient.Post(tokenRefreshURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Could not refresh access token.\n%d", err)
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&c.Auth.AccessToken); err != nil {
		return err
	}

	return nil
}

func (c *Client) sendRequest(req *http.Request, result interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Auth.AccessToken))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Bar response, status: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return fmt.Errorf("Could not parse the response data!\n%v", err)
	}

	return nil
}
