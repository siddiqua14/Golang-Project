package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	
)

type CatController struct {
	web.Controller
}
type CatImage struct {
	ID        string   `json:"id"`
	URL       string   `json:"url"`
	Width     int      `json:"width"`
	Height    int      `json:"height"`
	MimeType  string   `json:"mime_type"`
	Breeds    []Breed  `json:"breeds"`
	Categories []string `json:"categories"`
}
type Vote struct {
	ImageID string `json:"image_id"`
	Value   int    `json:"value"`
}

type Breed struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Origin       string `json:"origin"`
	WikipediaURL string `json:"wikipedia_url"`
}

func (c *CatController) GetCatImage() {
	apiKey, _ := web.AppConfig.String("catapi.key")
	apiURL, _ := web.AppConfig.String("catapi.url")

	imageURL, err := fetchCatImage(apiURL, apiKey)
	if err != nil {
		c.Data["CatImage"] = ""
	} else {
		c.Data["CatImage"] = imageURL
	}
	c.TplName = "index.tpl"
}


func (c *CatController) GetCatImagesAPI() {
	apiKey, _ := web.AppConfig.String("catapi.key")
	apiURL, _ := web.AppConfig.String("catapi.url")

	// Create a channel to receive cat images
	imageChan := make(chan []CatImage)
	errorChan := make(chan error)

	go fetchCatImages(apiURL, apiKey, imageChan, errorChan)

	select {
	case images := <-imageChan:
		c.Data["json"] = images
	case err := <-errorChan:
		c.Data["json"] = map[string]string{"error": err.Error()}
	}

	c.ServeJSON()
}
func (c *CatController) GetBreeds() {
	apiKey, _ := web.AppConfig.String("catapi.key")
	apiURL, _ := web.AppConfig.String("catapi.url")

	// Create a channel to receive breeds
	breedChan := make(chan []Breed)
	errorChan := make(chan error)

	go fetchBreeds(apiURL, apiKey, breedChan, errorChan)

	select {
	case breeds := <-breedChan:
		c.Data["json"] = breeds
	case err := <-errorChan:
		c.Data["json"] = map[string]string{"error": err.Error()}
	}

	c.ServeJSON()
}

func (c *CatController) GetBreedImages() {
	apiKey, _ := web.AppConfig.String("catapi.key")
	apiURL, _ := web.AppConfig.String("catapi.url")

	breedID := c.GetString("breed_id")

	// Create a channel to receive breed images
	imageChan := make(chan []CatImage)
	errorChan := make(chan error)

	go fetchBreedImages(apiURL, apiKey, breedID, imageChan, errorChan)

	select {
	case images := <-imageChan:
		c.Data["json"] = images
	case err := <-errorChan:
		c.Data["json"] = map[string]string{"error": err.Error()}
	}

	c.ServeJSON()
}
func fetchCatImage(apiURL, apiKey string) (string, error) {
	reqURL := apiURL + "/images/search?limit=1"
	req, _ := http.NewRequest("GET", reqURL, nil)
	req.Header.Add("x-api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	var result []CatImage
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	if len(result) > 0 {
		return result[0].URL, nil
	}

	return "", nil
}
// Fetch Cat Images Concurrently
func fetchCatImages(apiURL, apiKey string, imageChan chan []CatImage, errorChan chan error) {
	reqURL := apiURL + "/images/search?limit=10"
	req, _ := http.NewRequest("GET", reqURL, nil)
	req.Header.Add("x-api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errorChan <- err
		close(imageChan)
		close(errorChan)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		errorChan <- fmt.Errorf("API returned status code %d", resp.StatusCode)
		close(imageChan)
		close(errorChan)
		return
	}

	var result []CatImage
	err = json.Unmarshal(body, &result)
	if err != nil {
		errorChan <- err
		close(imageChan)
		close(errorChan)
		return
	}

	imageChan <- result
	close(imageChan)
	close(errorChan)
}
// Fetch Breeds Concurrently
func fetchBreeds(apiURL, apiKey string, breedChan chan []Breed, errorChan chan error) {
	reqURL := apiURL + "/breeds"
	req, _ := http.NewRequest("GET", reqURL, nil)
	req.Header.Add("x-api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errorChan <- err
		close(breedChan)
		close(errorChan)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		errorChan <- fmt.Errorf("API returned status code %d", resp.StatusCode)
		close(breedChan)
		close(errorChan)
		return
	}

	var result []Breed
	err = json.Unmarshal(body, &result)
	if err != nil {
		errorChan <- err
		close(breedChan)
		close(errorChan)
		return
	}

	breedChan <- result
	close(breedChan)
	close(errorChan)
}

// Fetch Breed Images Concurrently
func fetchBreedImages(apiURL, apiKey, breedID string, imageChan chan []CatImage, errorChan chan error) {
	reqURL := apiURL + "/images/search?breed_ids=" + breedID + "&limit=10"
	req, _ := http.NewRequest("GET", reqURL, nil)
	req.Header.Add("x-api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errorChan <- err
		close(imageChan)
		close(errorChan)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		errorChan <- fmt.Errorf("API returned status code %d", resp.StatusCode)
		close(imageChan)
		close(errorChan)
		return
	}

	var result []CatImage
	err = json.Unmarshal(body, &result)
	if err != nil {
		errorChan <- err
		close(imageChan)
		close(errorChan)
		return
	}

	imageChan <- result
	close(imageChan)
	close(errorChan)
}
// Handle voting on a cat image
func (c *CatController) CreateVote() {
	apiKey, _ := web.AppConfig.String("catapi.key")
	apiURL, _ := web.AppConfig.String("catapi.url")

	var vote Vote
	if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&vote); err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid request body: " + err.Error()}
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}

	if vote.Value != 1 && vote.Value != -1 {
		c.Data["json"] = map[string]string{"error": "Vote value must be 1 or -1"}
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}

	// Log the vote being created
	fmt.Printf("Creating vote for image %s with value %d\n", vote.ImageID, vote.Value)

	reqURL := fmt.Sprintf("%s/votes", apiURL)
	jsonData, _ := json.Marshal(vote)
	req, _ := http.NewRequest("POST", reqURL, bytes.NewBuffer(jsonData))
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to create vote: %v", err)
		fmt.Println(errMsg)
		c.Data["json"] = map[string]string{"error": errMsg}
		c.Ctx.ResponseWriter.WriteHeader(500)
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("API Response: %s\n", string(body))

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		errMsg := fmt.Sprintf("API returned status code %d: %s", resp.StatusCode, string(body))
		fmt.Println(errMsg)
		c.Data["json"] = map[string]string{"error": errMsg}
		c.Ctx.ResponseWriter.WriteHeader(resp.StatusCode)
		c.ServeJSON()
		return
	}

	fmt.Println("Vote created successfully")
	c.Data["json"] = map[string]string{"status": "success"}
	c.ServeJSON()
}

func (c *CatController) GetVotes() {
	apiKey, _ := web.AppConfig.String("catapi.key")
	apiURL, _ := web.AppConfig.String("catapi.url")

	reqURL := fmt.Sprintf("%s/votes", apiURL)
	req, _ := http.NewRequest("GET", reqURL, nil)
	req.Header.Set("x-api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to fetch votes"}
		c.Ctx.ResponseWriter.WriteHeader(500)
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		c.Data["json"] = map[string]string{"error": "Failed to fetch votes"}
		c.Ctx.ResponseWriter.WriteHeader(resp.StatusCode)
		c.ServeJSON()
		return
	}

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to parse votes"}
		c.Ctx.ResponseWriter.WriteHeader(500)
		c.ServeJSON()
		return
	}

	c.Data["json"] = result
	c.ServeJSON()
}

// Handle favoriting a cat image
func (c *CatController) CreateFavorite() {
    apiKey, _ := web.AppConfig.String("catapi.key")
    apiURL, _ := web.AppConfig.String("catapi.url")

    var favorite struct {
        ImageID string `json:"image_id"`
    }
    if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&favorite); err != nil {
        c.Data["json"] = map[string]string{"error": "Invalid request body: " + err.Error()}
        c.Ctx.ResponseWriter.WriteHeader(400)
        c.ServeJSON()
        return
    }

    reqURL := fmt.Sprintf("%s/favourites", apiURL)
    jsonData, _ := json.Marshal(favorite)
    req, _ := http.NewRequest("POST", reqURL, bytes.NewBuffer(jsonData))
    req.Header.Set("x-api-key", apiKey)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        errMsg := fmt.Sprintf("Failed to create favorite: %v", err)
        c.Data["json"] = map[string]string{"error": errMsg}
        c.Ctx.ResponseWriter.WriteHeader(500)
        c.ServeJSON()
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    if resp.StatusCode != 200 && resp.StatusCode != 201 {
        errMsg := fmt.Sprintf("API returned status code %d: %s", resp.StatusCode, string(body))
        c.Data["json"] = map[string]string{"error": errMsg}
        c.Ctx.ResponseWriter.WriteHeader(resp.StatusCode)
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]string{"status": "success"}
    c.ServeJSON()
}

// Handle fetching favorite cat images
func (c *CatController) GetFavorites() {
    apiKey, _ := web.AppConfig.String("catapi.key")
    apiURL, _ := web.AppConfig.String("catapi.url")

    reqURL := fmt.Sprintf("%s/favourites", apiURL)
    req, _ := http.NewRequest("GET", reqURL, nil)
    req.Header.Set("x-api-key", apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        c.Data["json"] = map[string]string{"error": "Failed to fetch favorites"}
        c.Ctx.ResponseWriter.WriteHeader(500)
        c.ServeJSON()
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)

    if resp.StatusCode != 200 {
        c.Data["json"] = map[string]string{"error": "Failed to fetch favorites"}
        c.Ctx.ResponseWriter.WriteHeader(resp.StatusCode)
        c.ServeJSON()
        return
    }

    var result interface{}
    if err := json.Unmarshal(body, &result); err != nil {
        c.Data["json"] = map[string]string{"error": "Failed to parse favorites"}
        c.Ctx.ResponseWriter.WriteHeader(500)
        c.ServeJSON()
        return
    }

    c.Data["json"] = result
    c.ServeJSON()
}