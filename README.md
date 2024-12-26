# Golang Project with Beego Framework

## Overview
This project implements a web application similar to the features provided by [The Cat API](https://thecatapi.com/). It is built using the Beego framework for the backend and Vanilla JavaScript, CSS, and Bootstrap for frontend interactions. The project includes:

- API integration using Go channels.
- Configuration handling for API keys and other settings using Beego config.
- High test coverage with Go unit tests (80%+).
- Interactive and responsive UI for fetching and voting on images.

## Features
- Fetches cat images dynamically from The Cat API.
- Allows users to vote on images.
- **Backend**: Built with Beego.
- **Frontend**: Built with Vanilla JavaScript and Bootstrap.
- Uses **Go channels** for asynchronous API calls.
- Supports both **Linux** and **Windows** environments.

## Prerequisites
Before you begin, ensure you have met the following requirements:

- **Go (Golang)**: Version 1.18 or higher.
- **Beego Framework**: Installed globally.
- **Git**: For cloning the repository.

## Installation and Setup

### Step 1: Install Go
   - Download and install Go from [https://go.dev/dl/](https://go.dev/dl/).
   - Verify installation:
     ```bash
     go version
     ```
### Step 2: Clone the repository
To get started, clone the project repository:
1. Navigate to your Go src `/go/src` directory to ensure the project is placed in the correct directory for your Go workspace.
2. Clone the repository:
```bash
git clone https://github.com/siddiqua14/Golang-Project.git
cd Golang-Project/catapi
```
Now the project is properly set up in your Go workspace at ~/go/src/Golang-Project/catapi, which is the correct location for Go to find and manage the project dependencies. 

### Step 3: Install Beego Framework
Beego is the framework used for this project, and Bee CLI is a development tool. If `Beego` is not installed in your workspace,
- Install them by running:
```bash
go get github.com/beego/beego/v2@latest
```
Ensure your `GOPATH` is set up correctly:
#### For Linux:
Add the following lines to your `~/.bashrc` or `~/.zshrc`:
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```
Reload your shell:
```bash
source ~/.bashrc
```
#### For Windows:
Set `GOPATH` and add it to your system `PATH`:
1. Open **Environment Variables** in System Properties.
2. Add a new variable:
   - Variable Name: `GOPATH`
   - Variable Value: `C:\Users\<YourUsername>\go` or in which workspace you like, setup path of that workspace
3. Edit the `Path` variable and add: `%GOPATH%\bin`
Verify your setup:
```bash
echo $GOPATH   # For Linux/MacOS
echo $env:GOPATH   # For Windows PowerShell
echo $env:Path 
```
##### To verify installation:
```bash
bee version
```

### Step 4: Install Dependencies

After installing the required Go modules and Beego dependencies:

```bash
go mod tidy
```
This will automatically resolve any missing dependencies and update your go.sum file with the required entries.


## Configuration

Open the configuration file located at `conf/app.conf` and update the following:

```ini
appname = catapi
httpport = 8000
runmode = dev
catapi.key = live_UeBfmyQ9TgLkkVLKsIF6FdYu9vaXTfddUioxblmRAkLgNBf8b1ko08b0KMOvHmfC
catapi.url = https://api.thecatapi.com/v1
```
## Running the Application

### Step 1: Start the Beego Server
Run the following command to start the Beego application:

```bash
bee run
```
### Step 2: Access the Application
Open your browser and navigate to: `http://localhost:8000`

![Screenshot from 2024-12-26 16-33-33](https://github.com/user-attachments/assets/0ee5ee18-aa6c-43d4-a1f6-2d8309220a3c)


## Application Layout and Features
### Voting Layout
#### Voting: 
The default layout is the voting section where you can view random cat images and vote on them. You can either Like, Dislike, or mark them as a Favorite using the respective buttons. Votes are posted to the API, and images are updated accordingly.

- Buttons:
   - Like: Upvote the image.
   - Dislike: Downvote the image.
   - Heart: Add the image to your favorite list.

#### Breeds Layout
- Breeds: In the breed section, you can search for different cat breeds using a search bar. The system fetches breed data from the API and displays detailed information, including breed images, description, origin, and a link to the breed’s Wikipedia page.

- Search Bar: Filter and search for specific breeds.

- Breed Information: After selecting a breed, you can view:
   - Breed Images
   - Breed Name
   - Breed Origin
   - Breed ID
   - Breed Description
   - Wikipedia Link
- Images Slider: A slider will display images of the selected breed.

#### Favorite Layout
- Favorites: Images that have been marked as favorites will be displayed in this section.

- Layout: The favorites are shown in a grid or list layout. You can switch between grid and list view.

- Delete: You can delete an image from the favorites by clicking the `×` button.

## API Endpoints
The application exposes the following routes. You can access these APIs on your local machine by visiting` http://localhost:8000`.

1. GET `/`
- Controller Method: `GetCatImage`
- Description: Displays a random cat image on the homepage.
2. GET `/api/catimage`
- Controller Method: `GetCatImagesAPI`
- Description: Fetches a random cat image via the API. This endpoint returns the image URL in JSON format.
3. GET `/votes`
- Controller Method: `GetVotes`
- Description: Retrieves all the votes that have been cast for cat images.
4. GET `/api/breeds`
- Controller Method: `GetBreeds`
- Description: Fetches a list of all cat breeds from the API.
5. GET `/getFavorites`
- Controller Method: `GetFavorites`
- Description: Retrieves a list of all cat images marked as favorites.
You can test these routes by running the application locally. Ensure that the app is running on `http://localhost:8000` and use tools like Postman or cURL to make requests to the API endpoints.
## Unit Testing

This project includes unit tests to ensure code reliability.

1. **Run all tests:**
```bash
go test ./tests -v
```
2. **Generate coverage report:**

```bash
go test -coverprofile coverage.out ./...
go tool cover -html coverage.out
```
or,
```bash
go test ./... -coverprofile=coverage.out && go tool cover -func=coverage.out
```

## Notes

- Ensure you have `bee` in your system `PATH` for running the application.
- If you encounter any issues, double-check your `GOPATH` and `PATH` settings.
- If there any issue with `GOPATH` in `Linux`: 
   1. Check if $GOPATH/bin is set correctly:
   ```bash
    echo $GOPATH
   ```
   It should print something like` /home/username/go`. If it doesn't, recheck the setup for `$GOPATH`.
   2. Verify Bee Binary: Check if the Bee binary is installed:
   ```bash
   ls $GOPATH/bin
   ```
   Look for a file named bee (without any file extension).
   3. Make sure the Go binary path is in your PATH: Ensure that `$GOPATH/bin` is added to your system’s PATH. You can confirm by running:

   ```bash
   echo $PATH
   ```
   If it's not there, manually add it to your .bashrc:
   ```bash
   echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc
   source ~/.bashrc
   ```
   4. Verify Bee Tool:

   After everything is set up, verify the Bee tool:
   ```bash
   bee version
   ```
