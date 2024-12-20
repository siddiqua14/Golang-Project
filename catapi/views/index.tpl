<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <title>Cat Voting</title>
    <link rel="stylesheet" href="/static/css/catapi.css">
    <link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.4/css/all.min.css" rel="stylesheet">
    <link rel="icon" href="/static/favicon.ico" type="image/x-icon">
    <script src="/static/js/catapi.js" defer></script>
</head>

<body>
    <div class="container">
        <h1>Cat Voting</h1>

        <!-- Buttons for switching between Voting and Breed -->
        <div class="button-container">
            <button id="votingButton" class="button" title="Voting" onclick="showVotingLayout()">Voting</button>
            <button id="breedButton" class="button" title="Breed" onclick="showBreedLayout()">Breed</button>
            <button id="favButton" class="button" title="Favorite" onclick="showFavoriteLayout()">❤️</button>
        </div>

        <!-- Main Card Container -->
        <div id="catCard" class="cat-card">
            <div id="votingLayout" class="layout">
                <div id="catImageContainer" class="image-container">
                    {{if .CatImage}}
                    <img id="catImage" src="{{.CatImage}}" alt="Cute Cat Image">
                    {{else}}
                    <p>No Image Available</p>
                    {{end}}
                </div>
                <div class="vote-buttons">
                    <button id="heartButton" class="button" title="Favorite">❤️</button>
                    <button id="likeButton" class="button" title="Like">👍</button>
                    <button id="dislikeButton" class="button" title="Dislike">👎</button>
                </div>
            </div>

            <!-- Breed Layout -->
            <div id="breedLayout" class="layout" style="display: none;">
                <h2>Select a Breed</h2>
                <div class="search-container">
                    <!-- Search bar for filtering breeds -->
                    <input type="text" id="breedSearch" class="dropdown" placeholder="Search for a breed..." />

                    <!-- List of breeds will appear as search results -->
                    <div id="breedList" class="breed-list"></div>
                </div>
                <!-- Breed information and images (will show after selecting a breed) -->
                <div id="breedInfoContainer">
                    <div class="breed-image-container">
                        <div id="breedImagesSlider" class="slide-transition"></div>
                    </div>
                    <div class="slider-indicators" id="sliderIndicators"></div>

                    <div class="breed-info">
                        <h2 id="breedName" class="breed-name">Breed Name</h2>
                        <p id="breedOrigin" class="breed-origin">Origin: Unknown</p>
                        <p id="breedDescription" class="breed-description">
                            Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                        </p>
                        <a id="breedWiki" href="#" target="_blank" class="breed-wiki">
                            Learn more on Wikipedia
                        </a>
                    </div>
                </div>

                <!-- Breed images slider (will show after selecting a breed) -->
                <div id="breedImagesContainer" class="slider-container" style="display: none;">
                    <div id="breedImagesSlider" class="slider"></div>
                    <div id="sliderIndicators" class="slider-indicators"></div>
                </div>
            </div>



            <!-- Favorite Layout (Initially Hidden) -->
            <div id="favoriteLayout" class="layout">
                <div class="layout-controls">
                    <button id="gridBtn" class="layout-btn grid-btn active">
                        <i class="fas fa-th"></i> Grid
                    </button>
                    <button id="listBtn" class="layout-btn list-btn">
                        <i class="fas fa-list"></i> List
                    </button>
                </div>
                <div id="favoriteImagesContainer" class="grid-layout">
                    <!-- Images will be added here -->
                </div>
            </div>
        </div>
    </div>
</body>

</html>