package integration_tests

import (
	"net/http"
	"testing"
)

func TestErrCheck(t *testing.T) {
	resp, err := c.Get(srv.URL + "/v1/err")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

  assertStatusCode(t, resp.StatusCode, http.StatusInternalServerError)
}
