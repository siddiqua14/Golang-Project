package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
	_ "catapi/main"
)

func TestMainApp(t *testing.T) {
	// Helper function to test routes.
	testMainRoute := func(req *http.Request) *httptest.ResponseRecorder {
		rr := httptest.NewRecorder()
		web.BeeApp.Handlers.ServeHTTP(rr, req)
		return rr
	}

}
