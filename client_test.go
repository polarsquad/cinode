package cinode

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const (
	companyID int32 = 1234
)

func TestRefreshToken(t *testing.T) {
	client, server := testClientwithFile(http.StatusOK, "test_data/client_test.txt")
	defer server.Close()

	err := client.RefreshToken()
	if err != nil {
		t.Fatalf("[ERROR] Refreshing Token, error: %v", err)
	}
	client.Auth.AccessToken = ""
	err = client.RefreshToken()
	if err == nil {
		t.Error("RefreshToken() should have caused Error with empty AccessToken")
		t.Errorf("Error: %v", err)
	}

	client.Auth.RefreshToken = ""
	err = client.RefreshToken()
	if err == nil {
		t.Error("RefreshToken() should have caused Error with empty AccessToken")
		t.Errorf("Error: %v", err)
	}
}

func testClient(httpStatus int, body io.Reader) (*Client, *httptest.Server) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(httpStatus)
		io.Copy(w, body)
		r.Body.Close()
		if closer, ok := body.(io.Closer); ok {
			closer.Close()
		}
	}))
	client := &Client{
		BaseURL:    testServer.URL + "/",
		CompanyID:  companyID,
		HTTPClient: http.DefaultClient,
		Auth: &Auth{
			Token: &Token{
				AccessToken:  "1qaz2wsx",
				RefreshToken: "3edc4rfv",
			},
		},
	}

	return client, testServer
}

func testClientwithFile(httpStatus int, filePath string) (*Client, *httptest.Server) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	return testClient(httpStatus, file)
}
