package integration_tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zoumas/aggrss/service/internal/server"
)

var (
	srv *httptest.Server
	c   *http.Client
)

func TestMain(m *testing.M) {
	srv = httptest.NewServer(server.ConfiguredRouter())
	c = srv.Client()
	defer srv.Close()

	m.Run()
}
