package integration_tests

import (
	"net/http"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	resp, err := c.Get(srv.URL + "/v1/healthz")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

  assertStatusCode(t, resp.StatusCode, http.StatusOK)
}

