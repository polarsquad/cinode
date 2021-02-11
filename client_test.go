package cinode

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
)

const (
	companyID int32 = 1234
)

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
