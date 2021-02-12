package cinode

// import (
// 	"net/http"
// 	"testing"
// )

// func TestRefreshToken(t *testing.T) {
// 	client, server := testClientwithFile(http.StatusOK, "test_data/client_test.txt")
// 	defer server.Close()

// 	err := client.RefreshToken()
// 	if err != nil {
// 		t.Fatalf("[ERROR] Initialising client, error: %v", err)
// 	}
// 	client.Auth.AccessToken = ""
// 	err = client.RefreshToken()
// 	if err == nil {
// 		t.Error("Refresh token should have caused Error with empty AccessToken")
// 	}
// }
