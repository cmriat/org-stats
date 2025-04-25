package orgstats

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-github/v39/github"
	"github.com/stretchr/testify/assert"
)

// TestFilterNonOrgMembers tests that non-organization members are filtered out
func TestFilterNonOrgMembers(t *testing.T) {
	// Create a mock server
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	// Mock the organization members endpoint
	mux.HandleFunc("/orgs/test-org/members", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// Return a list with one organization member
		w.Write([]byte(`[{"login":"org-member","id":1}]`))
	})

	// Create a client that uses our mock server
	client := github.NewClient(nil)
	url, _ := url.Parse(server.URL + "/")
	client.BaseURL = url
	client.UploadURL = url

	// Get organization members
	members, err := getOrgMembers(context.Background(), client, "test-org")

	// Verify the results
	assert.NoError(t, err)
	assert.NotNil(t, members)
	assert.Equal(t, 1, len(members))
	assert.True(t, members["org-member"])
	assert.False(t, members["non-org-member"])
}
