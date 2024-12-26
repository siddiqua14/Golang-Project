package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
    "strings"
	"fmt"
    "errors"
    "time"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
	"catapi/controllers"
   //"github.com/stretchr/testify/mock"
)
// MockHTTPClient for testing
type MockHTTPClient struct {
    DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`{"message": "SUCCESS"}`)),
	}, nil
}
var (
	apiURL string
	apiKey string
)

func init() {
	setupTestConfig()
}

// setupTestConfig initializes the test configuration
func setupTestConfig() {
	// Setup test configuration
	if err := web.LoadAppConfig("ini", "conf/app.conf"); err != nil {
		// If app.conf doesn't exist, set configuration directly
		web.BConfig.AppName = "catapi"
		web.AppConfig.Set("catapi.url", "https://api.thecatapi.com/v1")
		web.AppConfig.Set("catapi.key", "live_UeBfmyQ9TgLkkVLKsIF6FdYu9vaXTfddUioxblmRAkLgNBf8b1ko08b0KMOvHmfC")
	}

	// Load configuration values
	var err error
	apiURL, err = web.AppConfig.String("catapi.url")
	if err != nil {
		apiURL = "https://api.thecatapi.com/v1" // default value
	}
	apiKey, err = web.AppConfig.String("catapi.key")
	if err != nil {
		apiKey = "live_UeBfmyQ9TgLkkVLKsIF6FdYu9vaXTfddUioxblmRAkLgNBf8b1ko08b0KMOvHmfC" // default value
	}
}



// setupController creates and initializes a controller for testing
func setupController(w http.ResponseWriter, r *http.Request) *controllers.CatController {
    controller := &controllers.CatController{}
    ctx := context.NewContext()
    ctx.Reset(w, r)
    ctx.Input.SetData("RequestBody", r.Body)
    controller.Init(ctx, "", "", nil)
    controller.Ctx = ctx
    return controller
}

func TestCreateVote(t *testing.T) {
    tests := []struct {
        name         string
        voteData    map[string]interface{}
        expectedCode int
        expectedBody map[string]string
        mockResponse func() (*http.Response, error)
    }{
        {
            name: "Valid Vote",
            voteData: map[string]interface{}{
                "image_id": "test123", 
                "value":    1,
            },
            expectedCode: 200,
            expectedBody: map[string]string{"status": "success"},
            mockResponse: func() (*http.Response, error) {
                return &http.Response{
                    StatusCode: 201,
                    Body:       ioutil.NopCloser(bytes.NewBufferString(`{"message": "SUCCESS"}`)),
                }, nil
            },
        },
        {
            name: "API Returns Error Status",
            voteData: map[string]interface{}{
                "image_id": "test123",
                "value":    1,
            },
            expectedCode: 400,
            expectedBody: map[string]string{"error": "API returned status code 400: Bad Request"},
            mockResponse: func() (*http.Response, error) {
                return &http.Response{
                    StatusCode: 400,
                    Body:       ioutil.NopCloser(bytes.NewBufferString("Bad Request")),
                }, nil
            },
        },
        {
            name: "HTTP Client Error",
            voteData: map[string]interface{}{
                "image_id": "test123",
                "value":    1,
            },
            expectedCode: 500,
            expectedBody: map[string]string{"error": "Failed to create vote: network error"},
            mockResponse: func() (*http.Response, error) {
                return nil, fmt.Errorf("network error")
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            w := httptest.NewRecorder()
            jsonData, _ := json.Marshal(tt.voteData)
            r, _ := http.NewRequest("POST", "/api/votes", bytes.NewBuffer(jsonData))
            r.Header.Set("Content-Type", "application/json")

            controller := setupController(w, r)

            if tt.mockResponse != nil {
                httpClient := &MockHTTPClient{
                    DoFunc: func(req *http.Request) (*http.Response, error) {
                        return tt.mockResponse()
                    },
                }
                controller.SetHTTPClient(httpClient)
            }

            controller.CreateVote()

            assert.Equal(t, tt.expectedCode, w.Code)

            var response map[string]string
            json.Unmarshal(w.Body.Bytes(), &response)
            assert.Equal(t, tt.expectedBody, response)
        })
    }
}

func TestGetVotes(t *testing.T) {
    tests := []struct {
        name         string
        expectedCode int
        expectedBody map[string]string
        mockResponse func() (*http.Response, error)
    }{
        {
            name:         "Successful Votes Fetch",
            expectedCode: 200,
            mockResponse: func() (*http.Response, error) {
                return &http.Response{
                    StatusCode: 200,
                    Body:       ioutil.NopCloser(bytes.NewBufferString(`[{"id": "123", "value": 1}]`)),
                }, nil
            },
        },
        {
            name:         "API Error",
            expectedCode: 500,
            expectedBody: map[string]string{"error": "Failed to fetch votes"},
            mockResponse: func() (*http.Response, error) {
                return &http.Response{
                    StatusCode: 500,
                    Body:       ioutil.NopCloser(bytes.NewBufferString(`{"error": "Internal Server Error"}`)),
                }, nil
            },
        },
        {
            name:         "HTTP Client Error",
            expectedCode: 500,
            expectedBody: map[string]string{"error": "Failed to fetch votes"},
            mockResponse: func() (*http.Response, error) {
                return nil, fmt.Errorf("network error")
            },
        },
        {
            name:         "Invalid JSON Response",
            expectedCode: 500,
            expectedBody: map[string]string{"error": "Failed to parse votes"},
            mockResponse: func() (*http.Response, error) {
                return &http.Response{
                    StatusCode: 200,
                    Body:       ioutil.NopCloser(bytes.NewBufferString(`{invalid json`)),
                }, nil
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            w := httptest.NewRecorder()
            r, _ := http.NewRequest("GET", "/api/votes", nil)
            
            controller := setupController(w, r)

            if tt.mockResponse != nil {
                httpClient := &MockHTTPClient{
                    DoFunc: func(req *http.Request) (*http.Response, error) {
                        return tt.mockResponse()
                    },
                }
                controller.SetHTTPClient(httpClient)
            }

            controller.GetVotes()

            assert.Equal(t, tt.expectedCode, w.Code)

            // For error cases, check specific error messages
            if tt.expectedBody != nil {
                var response map[string]string
                err := json.Unmarshal(w.Body.Bytes(), &response)
                assert.NoError(t, err, "Response should be valid JSON")
                assert.Equal(t, tt.expectedBody, response)
            } else {
                // For successful case, just verify it's valid JSON
                var response interface{}
                err := json.Unmarshal(w.Body.Bytes(), &response)
                assert.NoError(t, err, "Response should be valid JSON")
            }
        })
    }
}

func TestCreateFavorite(t *testing.T) {
    tests := []struct {
        name         string
        favoriteData map[string]interface{}
        expectedCode int
        expectedBody map[string]string
        mockResponse func() (*http.Response, error)
    }{
        {
            name: "Valid Favorite",
            favoriteData: map[string]interface{}{
                "image_id": "test123",
            },
            expectedCode: 200,
            expectedBody: map[string]string{"status": "success"},
            mockResponse: func() (*http.Response, error) {
                return &http.Response{
                    StatusCode: 201,
                    Body:       ioutil.NopCloser(bytes.NewBufferString(`{"message": "SUCCESS"}`)),
                }, nil
            },
        },
        {
            name: "Invalid Request Body",
            favoriteData: map[string]interface{}{
                "image_id": "",
            },
            expectedCode: 400,
            expectedBody: map[string]string{"error": "image_id is required"},
            mockResponse: nil,
        },
        {
            name: "API Error",
            favoriteData: map[string]interface{}{
                "image_id": "test123",
            },
            expectedCode: 500,
            expectedBody: map[string]string{"error": "API returned status code 500: Internal Server Error"},
            mockResponse: func() (*http.Response, error) {
                return &http.Response{
                    StatusCode: 500,
                    Body:       ioutil.NopCloser(bytes.NewBufferString("Internal Server Error")),
                }, nil
            },
        },
        {
            name: "Network Error",
            favoriteData: map[string]interface{}{
                "image_id": "test123",
            },
            expectedCode: 500,
            expectedBody: map[string]string{"error": "Failed to create favorite: network error"},
            mockResponse: func() (*http.Response, error) {
                return nil, fmt.Errorf("network error")
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            w := httptest.NewRecorder()
            jsonData, _ := json.Marshal(tt.favoriteData)
            r, _ := http.NewRequest("POST", "/api/favorites", bytes.NewBuffer(jsonData))
            r.Header.Set("Content-Type", "application/json")

            controller := &controllers.CatController{}
            controller.Init(context.NewContext(), "", "", nil)
            controller.Ctx = context.NewContext()
            controller.Ctx.Reset(w, r)

            if tt.mockResponse != nil {
                mockClient := &MockHTTPClient{
                    DoFunc: func(req *http.Request) (*http.Response, error) {
                        return tt.mockResponse()
                    },
                }
                controller.SetHTTPClient(mockClient)
            }

            controller.CreateFavorite()

            assert.Equal(t, tt.expectedCode, w.Code)

            var response map[string]string
            json.Unmarshal(w.Body.Bytes(), &response)
            assert.Equal(t, tt.expectedBody, response)
        })
    }
}

func TestGetFavorites(t *testing.T) {
    tests := []struct {
        name         string
        expectedCode int
        expectedBody interface{}
        mockResponse func() (*http.Response, error)
    }{
        {
            name:         "Successful Fetch",
            expectedCode: 200,
            expectedBody: []interface{}{
                map[string]interface{}{
                    "id": float64(1),
                    "image_id": "test123",
                },
            },
            mockResponse: func() (*http.Response, error) {
                jsonResponse := `[{"id": 1, "image_id": "test123"}]`
                return &http.Response{
                    StatusCode: 200,
                    Body:       ioutil.NopCloser(bytes.NewBufferString(jsonResponse)),
                }, nil
            },
        },
        {
            name:         "API Error",
            expectedCode: 500,
            expectedBody: map[string]interface{}{"error": "Failed to fetch favorites"},
            mockResponse: func() (*http.Response, error) {
                return &http.Response{
                    StatusCode: 500,
                    Body:       ioutil.NopCloser(bytes.NewBufferString("Internal Server Error")),
                }, nil
            },
        },
        {
            name:         "Network Error",
            expectedCode: 500,
            expectedBody: map[string]interface{}{"error": "Failed to fetch favorites"},
            mockResponse: func() (*http.Response, error) {
                return nil, fmt.Errorf("network error")
            },
        },
        {
            name:         "Invalid JSON Response",
            expectedCode: 500,
            expectedBody: map[string]interface{}{"error": "Failed to parse favorites"},
            mockResponse: func() (*http.Response, error) {
                return &http.Response{
                    StatusCode: 200,
                    Body:       ioutil.NopCloser(bytes.NewBufferString("invalid json")),
                }, nil
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Create a response recorder and request
            w := httptest.NewRecorder()
            r, _ := http.NewRequest("GET", "/api/favorites", nil)

            // Setup controller
            controller := setupController(w, r)

            // Create mock HTTP client
            mockClient := &MockHTTPClient{
                DoFunc: func(req *http.Request) (*http.Response, error) {
                    // Verify request URL and API key
                    assert.Equal(t, fmt.Sprintf("%s/favourites", apiURL), req.URL.String())
                    assert.Equal(t, apiKey, req.Header.Get("x-api-key"))
                    return tt.mockResponse()
                },
            }
            controller.SetHTTPClient(mockClient)

            // Call the method
            controller.GetFavorites()

            // Check response status code
            assert.Equal(t, tt.expectedCode, w.Code)

            // Parse and check response body
            var response interface{}
            err := json.Unmarshal(w.Body.Bytes(), &response)
            assert.NoError(t, err)
            assert.Equal(t, tt.expectedBody, response)
        })
    }
}

func TestDeleteFavorite(t *testing.T) {
    tests := []struct {
        name         string
        favoriteID   string
        expectedCode int
        expectedBody map[string]interface{}
        mockResponse func() (*http.Response, error)
    }{
        {
            name:         "Successful Delete",
            favoriteID:   "123",
            expectedCode: 200,
            expectedBody: map[string]interface{}{"message": "Favorite deleted successfully"},
            mockResponse: func() (*http.Response, error) {
                return &http.Response{
                    StatusCode: 200,
                    Body:       ioutil.NopCloser(bytes.NewBufferString(`{"message": "SUCCESS"}`)),
                }, nil
            },
        },
        {
            name:         "Not Found",
            favoriteID:   "999",
            expectedCode: 404,
            expectedBody: map[string]interface{}{"error": "Failed to delete favorite: map[message:NOT_FOUND]"},
            mockResponse: func() (*http.Response, error) {
                return &http.Response{
                    StatusCode: 404,
                    Body:       ioutil.NopCloser(bytes.NewBufferString(`{"message": "NOT_FOUND"}`)),
                }, nil
            },
        },
        {
            name:         "Network Error",
            favoriteID:   "123",
            expectedCode: 500,
            expectedBody: map[string]interface{}{"error": "Failed to delete favorite"},
            mockResponse: func() (*http.Response, error) {
                return nil, fmt.Errorf("network error")
            },
        },
        {
            name:         "Invalid Request",
            favoriteID:   "",
            expectedCode: 405,
            expectedBody: map[string]interface{}{
                "error": "Failed to delete favorite: map[message:404 - please consult the documentation for correct url to call. https://docs.thecatapi.com/]",
            },
            mockResponse: nil,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Create a response recorder and request
            w := httptest.NewRecorder()
            r, _ := http.NewRequest("DELETE", "/api/favorites/"+tt.favoriteID, nil)

            // Setup controller
            controller := setupController(w, r)
            controller.Ctx.Input.SetParam(":id", tt.favoriteID)

            // Skip mock client setup for invalid request
            if tt.mockResponse != nil {
                mockClient := &MockHTTPClient{
                    DoFunc: func(req *http.Request) (*http.Response, error) {
                        // Verify request method, URL, and API key
                        assert.Equal(t, "DELETE", req.Method)
                        assert.Equal(t, fmt.Sprintf("%s/favourites/%s", apiURL, tt.favoriteID), req.URL.String())
                        assert.Equal(t, apiKey, req.Header.Get("x-api-key"))
                        return tt.mockResponse()
                    },
                }
                controller.SetHTTPClient(mockClient)
            }

            // Call the method
            controller.DeleteFavorite()

            // Check response status code
            assert.Equal(t, tt.expectedCode, w.Code)

            // Parse and check response body
            var response map[string]interface{}
            err := json.Unmarshal(w.Body.Bytes(), &response)
            assert.NoError(t, err)
            assert.Equal(t, tt.expectedBody, response)
        })
    }
}

func TestFetchBreeds_MultipleScenarios(t *testing.T) {
    apiURL := "https://api.example.com"
    apiKey := "test-api-key"

    t.Run("Successful Response", func(t *testing.T) {
        mockResponseBody := `[{"id":"beng","name":"Bengal"},{"id":"siam","name":"Siamese"}]`
        mockClient := &MockHTTPClient{
            DoFunc: func(req *http.Request) (*http.Response, error) {
                return &http.Response{
                    StatusCode: 200,
                    Body:       ioutil.NopCloser(strings.NewReader(mockResponseBody)),
                }, nil
            },
        }

        breedChan := make(chan []controllers.Breed, 1)
        errorChan := make(chan error, 1)

        go controllers.FetchBreeds(apiURL, apiKey, mockClient, breedChan, errorChan)

        select {
        case breeds := <-breedChan:
            assert.Equal(t, 2, len(breeds), "Expected 2 breeds")
            assert.Equal(t, "Bengal", breeds[0].Name, "First breed should be Bengal")
        case err := <-errorChan:
            t.Fatalf("Unexpected error: %v", err)
        }
    })

    t.Run("Client Do Error", func(t *testing.T) {
        mockClient := &MockHTTPClient{
            DoFunc: func(req *http.Request) (*http.Response, error) {
                return nil, fmt.Errorf("network error")
            },
        }

        breedChan := make(chan []controllers.Breed, 1)
        errorChan := make(chan error, 1)

        go controllers.FetchBreeds(apiURL, apiKey, mockClient, breedChan, errorChan)

        select {
        case <-breedChan:
            t.Fatal("Expected error, but received breeds")
        case err := <-errorChan:
            assert.Error(t, err)
            assert.Contains(t, err.Error(), "network error")
        }
    })

    t.Run("Non-200 Status Code", func(t *testing.T) {
        mockClient := &MockHTTPClient{
            DoFunc: func(req *http.Request) (*http.Response, error) {
                return &http.Response{
                    StatusCode: 404,
                    Body:       ioutil.NopCloser(strings.NewReader("Not Found")),
                }, nil
            },
        }

        breedChan := make(chan []controllers.Breed, 1)
        errorChan := make(chan error, 1)

        go controllers.FetchBreeds(apiURL, apiKey, mockClient, breedChan, errorChan)

        select {
        case <-breedChan:
            t.Fatal("Expected error, but received breeds")
        case err := <-errorChan:
            assert.Error(t, err)
            assert.Contains(t, err.Error(), "API returned status code 404")
        }
    })

    t.Run("JSON Unmarshal Error", func(t *testing.T) {
        mockResponseBody := `[{"id":"beng","name":123}]` // Invalid JSON
        mockClient := &MockHTTPClient{
            DoFunc: func(req *http.Request) (*http.Response, error) {
                return &http.Response{
                    StatusCode: 200,
                    Body:       ioutil.NopCloser(strings.NewReader(mockResponseBody)),
                }, nil
            },
        }

        breedChan := make(chan []controllers.Breed, 1)
        errorChan := make(chan error, 1)

        go controllers.FetchBreeds(apiURL, apiKey, mockClient, breedChan, errorChan)

        select {
        case <-breedChan:
            t.Fatal("Expected error, but received breeds")
        case err := <-errorChan:
            assert.Error(t, err)
            assert.Contains(t, err.Error(), "cannot unmarshal")
        }
    })
}
func TestFetchBreedImages_MultipleScenarios(t *testing.T) {
    setupTestConfig()
    breedID := "beng"

    t.Run("Successful Response", func(t *testing.T) {
        mockResponseBody := `[{"id":"cat123", "url": "https://example.com/cat123.jpg"}]`
        mockClient := &MockHTTPClient{
            DoFunc: func(req *http.Request) (*http.Response, error) {
                return &http.Response{
                    StatusCode: 200,
                    Body:       ioutil.NopCloser(strings.NewReader(mockResponseBody)),
                }, nil
            },
        }

        imageChan := make(chan []controllers.CatImage, 1)
        errorChan := make(chan error, 1)

        go controllers.FetchBreedImages(apiURL, apiKey, breedID, mockClient, imageChan, errorChan)

        select {
        case images := <-imageChan:
            assert.Equal(t, 1, len(images), "Expected 1 image")
            assert.Equal(t, "https://example.com/cat123.jpg", images[0].URL, "Image URL should match the mock response")
        case err := <-errorChan:
            t.Fatalf("Unexpected error: %v", err)
        }
    })

    t.Run("Client Do Error", func(t *testing.T) {
        mockClient := &MockHTTPClient{
            DoFunc: func(req *http.Request) (*http.Response, error) {
                return nil, fmt.Errorf("network error")
            },
        }

        imageChan := make(chan []controllers.CatImage, 1)
        errorChan := make(chan error, 1)

        go controllers.FetchBreedImages(apiURL, apiKey, breedID, mockClient, imageChan, errorChan)

        select {
        case <-imageChan:
            t.Fatal("Expected error, but received images")
        case err := <-errorChan:
            assert.Error(t, err)
            assert.Contains(t, err.Error(), "network error")
        }
    })

    t.Run("Non-200 Status Code", func(t *testing.T) {
        testCases := []struct {
            statusCode int
            name       string
        }{
            {404, "Not Found"},
            {500, "Internal Server Error"},
            {403, "Forbidden"},
        }

        for _, tc := range testCases {
            t.Run(tc.name, func(t *testing.T) {
                mockClient := &MockHTTPClient{
                    DoFunc: func(req *http.Request) (*http.Response, error) {
                        return &http.Response{
                            StatusCode: tc.statusCode,
                            Body:       ioutil.NopCloser(strings.NewReader("Error response")),
                        }, nil
                    },
                }

                imageChan := make(chan []controllers.CatImage, 1)
                errorChan := make(chan error, 1)

                go controllers.FetchBreedImages(apiURL, apiKey, breedID, mockClient, imageChan, errorChan)

                select {
                case <-imageChan:
                    t.Fatal("Expected error, but received images")
                case err := <-errorChan:
                    assert.Error(t, err)
                    assert.Contains(t, err.Error(), fmt.Sprintf("API returned status code %d", tc.statusCode))
                }
            })
        }
    })

    t.Run("JSON Unmarshal Error", func(t *testing.T) {
        testCases := []struct {
            name           string
            mockBody       string
            expectedErrMsg string
        }{
            {
                name:           "Invalid JSON Structure",
                mockBody:       `[{"id":123, "url": 456}]`, // Incorrect types
                expectedErrMsg: "cannot unmarshal",
            },
            {
                name:           "Malformed JSON",
                mockBody:       `[{"id":"cat123", "url": "https://example.com/cat123.jpg"`,  // Incomplete JSON
                expectedErrMsg: "unexpected end of JSON input",
            },
        }
    
        for _, tc := range testCases {
            t.Run(tc.name, func(t *testing.T) {
                mockClient := &MockHTTPClient{
                    DoFunc: func(req *http.Request) (*http.Response, error) {
                        return &http.Response{
                            StatusCode: 200,
                            Body:       ioutil.NopCloser(strings.NewReader(tc.mockBody)),
                        }, nil
                    },
                }
    
                imageChan := make(chan []controllers.CatImage, 1)
                errorChan := make(chan error, 1)
    
                go controllers.FetchBreedImages(apiURL, apiKey, breedID, mockClient, imageChan, errorChan)
    
                select {
                case <-imageChan:
                    t.Fatal("Expected error, but received images")
                case err := <-errorChan:
                    assert.Error(t, err)
                    assert.Contains(t, err.Error(), tc.expectedErrMsg)
                }
            })
        }
    })
}

// TestGetCatImage tests the GetCatImage method of the CatController

func TestCatController_GetCatImage(t *testing.T) {
    // Mock configuration values
    setupTestConfig()

    t.Run("Successful Image Fetch and Template Render", func(t *testing.T) {
        mockClient := &MockHTTPClient{
            DoFunc: func(req *http.Request) (*http.Response, error) {
                // Verify request details
                assert.Equal(t, apiURL+"/images/search?limit=1", req.URL.String())
                assert.Equal(t, apiKey, req.Header.Get("x-api-key"))

                mockResponse := []controllers.CatImage{
                    {URL: "https://cdn2.thecatapi.com/images/test-cat.jpg"},
                }
                responseBody, _ := json.Marshal(mockResponse)

                return &http.Response{
                    StatusCode: 200,
                    Body:       ioutil.NopCloser(bytes.NewBuffer(responseBody)),
                }, nil
            },
        }

        // Create controller and set mock client
        controller := &controllers.CatController{}
        controller.SetHTTPClient(mockClient)

        // Call method
        controller.GetCatImage()

        // Assertions
        assert.Equal(t, "https://cdn2.thecatapi.com/images/test-cat.jpg", controller.Data["CatImage"])
        assert.Equal(t, "index.tpl", controller.TplName)
    })

    t.Run("Image Fetch Error", func(t *testing.T) {
        mockClient := &MockHTTPClient{
            DoFunc: func(req *http.Request) (*http.Response, error) {
                return nil, fmt.Errorf("network error")
            },
        }

        controller := &controllers.CatController{}
        controller.SetHTTPClient(mockClient)

        // Call method
        controller.GetCatImage()

        // Assertions
        assert.Empty(t, controller.Data["CatImage"])
        assert.Equal(t, "index.tpl", controller.TplName)
    })

    t.Run("No Images Returned", func(t *testing.T) {
        mockClient := &MockHTTPClient{
            DoFunc: func(req *http.Request) (*http.Response, error) {
                return &http.Response{
                    StatusCode: 200,
                    Body:       ioutil.NopCloser(strings.NewReader("[]")),
                }, nil
            },
        }

        controller := &controllers.CatController{}
        controller.SetHTTPClient(mockClient)

        // Call method
        controller.GetCatImage()

        // Assertions
        assert.Empty(t, controller.Data["CatImage"])
        assert.Equal(t, "index.tpl", controller.TplName)
    })
}


func TestFetchCatImages(t *testing.T) {
	tests := []struct {
		name           string
		mockResponse   *http.Response
		mockError      error
		expectedError  bool
		expectedImages bool
	}{
		{
			name: "successful_fetch",
			mockResponse: &http.Response{
				StatusCode: 200,
				Body: ioutil.NopCloser(bytes.NewBufferString(`[
					{"url": "http://example.com/cat1.jpg"},
					{"url": "http://example.com/cat2.jpg"}
				]`)),
			},
			mockError:      nil,
			expectedError:  false,
			expectedImages: true,
		},
		{
			name:           "http_client_error",
			mockResponse:   nil,
			mockError:      errors.New("network error"),
			expectedError:  true,
			expectedImages: false,
		},
		{
			name: "api_error_status",
			mockResponse: &http.Response{
				StatusCode: 500,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`{"error": "server error"}`)),
			},
			mockError:      nil,
			expectedError:  true,
			expectedImages: false,
		},
		{
			name: "invalid_json_response",
			mockResponse: &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`invalid json`)),
			},
			mockError:      nil,
			expectedError:  true,
			expectedImages: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create mock client
			mockClient := &MockHTTPClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					// Verify request headers and URL
					assert.Equal(t, apiURL+"/images/search?limit=10", req.URL.String())
					assert.Equal(t, apiKey, req.Header.Get("x-api-key"))
					
					return tc.mockResponse, tc.mockError
				},
			}

			imageChan := make(chan []controllers.CatImage)
			errorChan := make(chan error)

			// Run FetchCatImages in a goroutine
			go controllers.FetchCatImages(mockClient, apiURL, apiKey, imageChan, errorChan)

			// Use select with timeout to prevent test from hanging
			select {
			case images := <-imageChan:
				if tc.expectedError {
					t.Error("Expected error but got images")
				}
				if tc.expectedImages {
					assert.NotNil(t, images)
					assert.Greater(t, len(images), 0)
				}
			case err := <-errorChan:
				if !tc.expectedError {
					t.Errorf("Expected success but got error: %v", err)
				}
			case <-time.After(2 * time.Second):
				t.Fatal("Test timed out")
			}

			// Verify channels are closed
			_, ok1 := <-imageChan
			assert.False(t, ok1, "Image channel should be closed")
			_, ok2 := <-errorChan
			assert.False(t, ok2, "Error channel should be closed")
		})
	}
}