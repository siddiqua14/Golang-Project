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

## Installation

## Step 1: Install Go
   - Download and install Go from [https://go.dev/dl/](https://go.dev/dl/).
   - Verify installation:
     ```bash
     go version
     ```
## Step 2: Install Beego Framework
Beego is the framework used for this project, and Bee CLI is a development tool.
- Install them by running:
```bash
go get github.com/beego/beego/v2@latest
```
Ensure your `GOPATH` is set up correctly:
### For Linux:
Add the following lines to your `~/.bashrc` or `~/.zshrc`:
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```
Reload your shell:
```bash
source ~/.bashrc
```
### For Windows:
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
#### To verify installation:
```bash
bee version
```

## Step 3: Install Dependencies

After installing the required Go modules and Beego dependencies:

```bash
go mod tidy
```
This will automatically resolve any missing dependencies and update your go.sum file with the required entries.
### 1. Clone the repository
To get started, clone the project repository:

```bash
git clone https://github.com/siddiqua14/Golang-Project.git
cd Golang-Project/catapi
```

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
Open your browser and navigate to: `http://localhost:8080`

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

## Notes

- Ensure you have `bee` in your system `PATH` for running the application.
- If you encounter any issues, double-check your `GOPATH` and `PATH` settings.
