package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	
	"github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
	"github.com/beego/beego/v2/server/web/context"
	_ "catapi/routers"
)

func TestRouters(t *testing.T) {
	// Define a helper to test the routes
	testRoute := func(req *http.Request) *httptest.ResponseRecorder {
		// Create a new ResponseRecorder to record the response
		rr := httptest.NewRecorder()

		// Create a new context
		ctx := context.NewContext()
		ctx.Reset(rr, req)
		web.BeeApp.Handlers.ServeHTTP(rr, req)

		return rr
	}

	// Test valid routes
	t.Run("Test GET /api/catimage", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/api/catimage", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		resp := testRoute(req)
		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("Test GET /votes", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/votes", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		resp := testRoute(req)
		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("Test GET /api/breeds", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/api/breeds", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		resp := testRoute(req)
		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("Test GET /api/breed-images", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/api/breed-images", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		resp := testRoute(req)
		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("Test GET /getFavorites", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/getFavorites", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		resp := testRoute(req)
		assert.Equal(t, http.StatusOK, resp.Code)
	})
}