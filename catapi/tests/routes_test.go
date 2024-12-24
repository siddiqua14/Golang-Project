package tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"catapi/routers"
	"catapi/controllers"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func setupRequestAndRecorder(method, url string) (*httptest.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, nil)
	rec := httptest.NewRecorder()
	return req, rec
}

// Test for GetCatImage route
func TestGetCatImageRoute(t *testing.T) {
	web.BConfig.RunMode = "test"

	// Set up the mock request and response recorder
	req, rec := setupRequestAndRecorder("GET", "/")

	// Set up the controller context
	controller := &controllers.CatController{}
	controller.Ctx.Request = req
	controller.Ctx.ResponseWriter = rec
	controller.Ctx.Input = context.NewInput()
	controller.Ctx.Output = context.NewResponse()

	// Call the controller method for this route
	controller.GetCatImage()

	// Assert the response
	assert.Equal(t, 200, rec.Code)
	assert.Contains(t, rec.Body.String(), "CatImage")
}

// Test for GetCatImagesAPI route
func TestGetCatImagesAPIRoute(t *testing.T) {
	web.BConfig.RunMode = "test"

	// Set up the mock request and response recorder
	req, rec := setupRequestAndRecorder("GET", "/api/catimage")

	// Set up the controller context
	controller := &controllers.CatController{}
	controller.Ctx.Request = req
	controller.Ctx.ResponseWriter = rec
	controller.Ctx.Input = context.NewInput()
	controller.Ctx.Output = context.NewResponse()

	// Call the controller method for this route
	controller.GetCatImagesAPI()

	// Assert the response
	assert.Equal(t, 200, rec.Code)
	assert.Contains(t, rec.Body.String(), "https://example.com/cat1.jpg") // Adjust based on mock response
}

// Test for CreateVote route
func TestCreateVoteRoute(t *testing.T) {
	web.BConfig.RunMode = "test"

	// Set up the mock request and response recorder (POST method)
	req, rec := setupRequestAndRecorder("POST", "/vote")

	// Set up the controller context
	controller := &controllers.CatController{}
	controller.Ctx.Request = req
	controller.Ctx.ResponseWriter = rec
	controller.Ctx.Input = context.NewInput()
	controller.Ctx.Output = context.NewResponse()

	// Call the controller method for this route
	controller.CreateVote()

	// Assert the response
	assert.Equal(t, 200, rec.Code)
}

// Test for GetVotes route
func TestGetVotesRoute(t *testing.T) {
	web.BConfig.RunMode = "test"

	// Set up the mock request and response recorder (GET method)
	req, rec := setupRequestAndRecorder("GET", "/votes")

	// Set up the controller context
	controller := &controllers.CatController{}
	controller.Ctx.Request = req
	controller.Ctx.ResponseWriter = rec
	controller.Ctx.Input = context.NewInput()
	controller.Ctx.Output = context.NewResponse()

	// Call the controller method for this route
	controller.GetVotes()

	// Assert the response
	assert.Equal(t, 200, rec.Code)
}

// Test for GetBreeds route
func TestGetBreedsRoute(t *testing.T) {
	web.BConfig.RunMode = "test"

	// Set up the mock request and response recorder (GET method)
	req, rec := setupRequestAndRecorder("GET", "/api/breeds")

	// Set up the controller context
	controller := &controllers.CatController{}
	controller.Ctx.Request = req
	controller.Ctx.ResponseWriter = rec
	controller.Ctx.Input = context.NewInput()
	controller.Ctx.Output = context.NewResponse()

	// Call the controller method for this route
	controller.GetBreeds()

	// Assert the response
	assert.Equal(t, 200, rec.Code)
}

// Test for GetBreedImages route
func TestGetBreedImagesRoute(t *testing.T) {
	web.BConfig.RunMode = "test"

	// Set up the mock request and response recorder (GET method)
	req, rec := setupRequestAndRecorder("GET", "/api/breed-images")

	// Set up the controller context
	controller := &controllers.CatController{}
	controller.Ctx.Request = req
	controller.Ctx.ResponseWriter = rec
	controller.Ctx.Input = context.NewInput()
	controller.Ctx.Output = context.NewResponse()

	// Call the controller method for this route
	controller.GetBreedImages()

	// Assert the response
	assert.Equal(t, 200, rec.Code)
}

// Test for CreateFavorite route
func TestCreateFavoriteRoute(t *testing.T) {
	web.BConfig.RunMode = "test"

	// Set up the mock request and response recorder (POST method)
	req, rec := setupRequestAndRecorder("POST", "/createFavorite")

	// Set up the controller context
	controller := &controllers.CatController{}
	controller.Ctx.Request = req
	controller.Ctx.ResponseWriter = rec
	controller.Ctx.Input = context.NewInput()
	controller.Ctx.Output = context.NewResponse()

	// Call the controller method for this route
	controller.CreateFavorite()

	// Assert the response
	assert.Equal(t, 200, rec.Code)
}

// Test for GetFavorites route
func TestGetFavoritesRoute(t *testing.T) {
	web.BConfig.RunMode = "test"

	// Set up the mock request and response recorder (GET method)
	req, rec := setupRequestAndRecorder("GET", "/getFavorites")

	// Set up the controller context
	controller := &controllers.CatController{}
	controller.Ctx.Request = req
	controller.Ctx.ResponseWriter = rec
	controller.Ctx.Input = context.NewInput()
	controller.Ctx.Output = context.NewResponse()

	// Call the controller method for this route
	controller.GetFavorites()

	// Assert the response
	assert.Equal(t, 200, rec.Code)
}

// Test for DeleteFavorite route
func TestDeleteFavoriteRoute(t *testing.T) {
	web.BConfig.RunMode = "test"

	// Set up the mock request and response recorder (DELETE method)
	req, rec := setupRequestAndRecorder("DELETE", "/deleteFavorite/1")

	// Set up the controller context
	controller := &controllers.CatController{}
	controller.Ctx.Request = req
	controller.Ctx.ResponseWriter = rec
	controller.Ctx.Input = context.NewInput()
	controller.Ctx.Output = context.NewResponse()

	// Call the controller method for this route
	controller.DeleteFavorite()

	// Assert the response
	assert.Equal(t, 200, rec.Code)
}
