package integration_tests

import (
	"net/http"
	"strings"
	"testing"
)

func TestPanicCheck(t *testing.T) {
  // Server must recover from panics and respond with a 500 status code.

	resp, err := c.Post(srv.URL + "/v1/panic", "", strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

  assertStatusCode(t, resp.StatusCode, http.StatusInternalServerError)
}
