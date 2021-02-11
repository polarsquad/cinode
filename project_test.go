package cinode

import (
	"net/http"
	"testing"
)

func TestGetProjects(t *testing.T) {
	client, server := testClientwithFile(http.StatusOK, "test_data/project_list.txt")
	defer server.Close()

	projectList, err := client.GetProjects()
	if err != nil {
		t.Fatal(err)
	}

	if len(*projectList) != 3 {
		t.Error("Mismatch number of Projects in the list`´")
	}

	if (*projectList)[0].CompanyID != companyID {
		t.Error("Mismatch in Company ID")
	}

	for _, v := range (*projectList)[1:] {
		if v.CustomerID != 78902 {
			t.Error("Mismatch in Customer ID")
		}
	}
}
