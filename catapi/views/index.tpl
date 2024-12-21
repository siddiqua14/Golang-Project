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
        <div class="button-container" id="catTabs" role="tablist">

            <button id="votingButton" class="button" title="Voting" onclick="showVotingLayout()" data-bs-toggle="tab"
                data-bs-target="#voting" type="button">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M7 14l5-5 5 5" />
                </svg>Voting
            </button>


            <button id="breedButton" class="button" title="Breed" onclick="showBreedLayout()" data-bs-toggle="tab"
                data-bs-target="#breeds" type="button"><svg width="16" height="16" viewBox="0 0 24 24" fill="none"
                    stroke="currentColor" stroke-width="2">
                    <circle cx="11" cy="11" r="8" />
                    <path d="M21 21l-4.35-4.35" />
                </svg>
                Breeds</button>

            <button id="favButton" class="button" title="Favorite" onclick="showFavoriteLayout()">❤️</button>
        </div>

        <!-- Main Card Container -->
        <div id="catCard" class="cat-card">
            <div id="votingLayout" class="layout show active" role="tabpanel">
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
            <div id="breedLayout" class="layout" style="display: none;" role="tabpanel">
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
                            Wikipedia
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
                        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor"
                            stroke-width="2">
                            <rect x="3" y="3" width="7" height="7" />
                            <rect x="14" y="3" width="7" height="7" />
                            <rect x="3" y="14" width="7" height="7" />
                            <rect x="14" y="14" width="7" height="7" />
                        </svg>
                    </button>
                    <button id="listBtn" class="layout-btn list-btn">
                        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor"
                            stroke-width="2">
                            <line x1="3" y1="6" x2="21" y2="6" />
                            <line x1="3" y1="12" x2="21" y2="12" />
                            <line x1="3" y1="18" x2="21" y2="18" />
                        </svg>
                    </button>
                </div>
                <div id="favoriteImagesContainer" class="grid-layout">
                    <!-- Images will be added here -->
                </div>
            </div>
        </div>
    </div>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/bootstrap/5.3.2/js/bootstrap.bundle.min.js"></script>
</body>

</html>