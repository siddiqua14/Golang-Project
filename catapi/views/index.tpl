<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <title>Cat Voting</title>
    <link rel="stylesheet" href="/static/css/catapi.css">
    <link rel="icon" href="/static/favicon.ico" type="image/x-icon">
    <script src="/static/js/catapi.js" defer></script>

    <link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.4/css/all.min.css" rel="stylesheet">
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.11.2/font/bootstrap-icons.css">
</head>

<body>
    <div class="Cat-container">
        <div class="button-container" id="catTabs" role="tablist">

            <button id="votingButton" class="button" title="Voting" onclick="showVotingLayout()" data-bs-toggle="tab"
                data-bs-target="#voting" type="button">
                <svg class="w-20 h-20 hover:stroke-primary hover:fill-primary" fill="none" stroke="currentColor"
                    stroke-width="10" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg"
                    xmlns:xlink="http://www.w3.org/1999/xlink" width="24px" height="24px" viewBox="0 0 377.03 377.031"
                    xml:space="preserve">
                    <g>
                        <g>
                            <path
                                d="M58.876,315.796h48.549c13.426,0,24.355-10.93,24.355-24.355V184.819c0-0.918,0.335-1.74,0.985-2.391
                    c0.631-0.611,1.463-0.956,2.4-0.956l13.961,0.115c7.287,0.096,11.628-2.668,14-5.002c2.869-2.858,4.437-6.665,4.437-10.738
                    c0-4.035-1.472-8.052-4.389-11.934l-62.701-83.806c-4.16-5.604-10.241-8.836-16.687-8.836c-6.436,0-12.527,3.194-16.706,8.779
                    l-62.682,83.48C1.482,157.414,0,161.449,0,165.503c0,7.86,6.292,15.835,18.322,15.835H31.04c1.913,0,3.471,1.568,3.471,3.472
                    v106.621C34.521,304.876,45.451,315.796,58.876,315.796z M28.171,160.483c-1.587,0-0.306-3.423,2.859-7.65l47-62.605
                    c3.174-4.227,8.3-4.217,11.465,0.009l46.952,62.778c3.165,4.227,4.198,7.641,2.324,7.631l-3.404-0.028
                    c-6.531-0.077-12.651,2.41-17.279,6.99c-4.638,4.58-7.191,10.7-7.191,17.222v106.622c0,1.922-1.559,3.471-3.471,3.471H58.876
                    c-1.913,0-3.481-1.549-3.481-3.471V184.819c0-13.406-10.93-24.336-24.346-24.336C31.049,160.483,29.758,160.483,28.171,160.483z">
                            </path>
                            <path
                                d="M318.154,61.234h-48.549c-13.426,0-24.347,10.93-24.347,24.355V192.21c0,0.92-0.344,1.742-0.994,2.391
                    c-0.631,0.613-1.463,0.957-2.4,0.957l-13.961-0.115c-7.287-0.096-11.628,2.668-14,5.002c-2.868,2.859-4.437,6.664-4.437,10.738
                    c0,4.035,1.473,8.053,4.389,11.953l62.701,83.787c4.16,5.604,10.242,8.836,16.687,8.836c6.437,0,12.527-3.193,16.706-8.779
                    l62.673-83.48c2.926-3.883,4.408-7.918,4.408-11.973c0-7.859-6.292-15.834-18.321-15.834H345.99c-1.912,0-3.471-1.568-3.471-3.473
                    V85.599C342.51,72.164,331.58,61.234,318.154,61.234z M348.859,216.548c1.587,0,0.306,3.424-2.869,7.65L299,286.804
                    c-3.175,4.227-8.31,4.217-11.465-0.01l-46.952-62.77c-3.165-4.227-4.208-7.639-2.324-7.631l3.404,0.029
                    c6.531,0.076,12.651-2.41,17.28-6.98c4.638-4.58,7.19-10.701,7.19-17.223V85.599c0-1.922,1.559-3.471,3.472-3.471h48.549
                    c1.922,0,3.48,1.549,3.48,3.471V192.22c0,13.408,10.93,24.338,24.346,24.338C345.98,216.548,347.271,216.548,348.859,216.548z">
                            </path>
                        </g>
                    </g>
                </svg>Voting


            </button>


            <button id="breedButton" class="button" title="Breed" onclick="showBreedLayout()" data-bs-toggle="tab"
                data-bs-target="#breeds" type="button"><svg width="20" height="20" viewBox="0 0 24 24" fill="none"
                    stroke="currentColor" stroke-width="2">
                    <circle cx="11" cy="11" r="8" />
                    <path d="M21 21l-4.35-4.35" />
                </svg>
                Breeds</button>

            <button id="favButton" class="button" title="Favorite" onclick="showFavoriteLayout()" data-bs-toggle="tab"
                data-bs-target="#breeds" type="button"><i class='far fa-heart'></i>Favs</button>
        </div>

        <!-- Main Card Container -->
        <div id="catCard" class="cat-card">


            <div id="votingLayout" class="layout show active" role="tabpanel">
                <div id="catImageContainer" class="image-container">
                    <img id="loadingImage" src="../static/img/loading.png" alt="Loading..." style="display: none;">
                    {{if .CatImage}}
                    <img id="catImage" src="{{.CatImage}}" alt="Cute Cat Image">
                    {{else}}
                    <p>No Image Available</p>
                    {{end}}
                </div>
                <div class="vote-buttons">
                    <button id="heartButton" class="button" title="Favorite"><svg
                            class="w-6 h-6 hover:stroke-primary hover:fill-primary" version="1.0" id="Layer_1"
                            xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"
                            viewBox="0 0 64 64" enable-background="new 0 0 64 64" xml:space="preserve"
                            stroke="currentColor">
                            <g>
                                <path fill="#231F20" d="M48,6c-4.418,0-8.418,1.791-11.313,4.687l-3.979,3.961c-0.391,0.391-1.023,0.391-1.414,0
                        c0,0-3.971-3.97-3.979-3.961C24.418,7.791,20.418,6,16,6C7.163,6,0,13.163,0,22c0,3.338,1.024,6.436,2.773,9
                        c0,0,0.734,1.164,1.602,2.031s24.797,24.797,24.797,24.797C29.953,58.609,30.977,59,32,59s2.047-0.391,2.828-1.172
                        c0,0,23.93-23.93,24.797-24.797S61.227,31,61.227,31C62.976,28.436,64,25.338,64,22C64,13.163,56.837,6,48,6z M58.714,30.977
                        c0,0-0.612,0.75-1.823,1.961S33.414,56.414,33.414,56.414C33.023,56.805,32.512,57,32,57s-1.023-0.195-1.414-0.586
                        c0,0-22.266-22.266-23.477-23.477s-1.823-1.961-1.823-1.961C3.245,28.545,2,25.424,2,22C2,14.268,8.268,8,16,8
                        c3.866,0,7.366,1.566,9.899,4.101l0.009-0.009l4.678,4.677c0.781,0.781,2.047,0.781,2.828,0l4.678-4.677l0.009,0.009
                        C40.634,9.566,44.134,8,48,8c7.732,0,14,6.268,14,14C62,25.424,60.755,28.545,58.714,30.977z">
                                </path>
                                <path fill="#231F20" d="M48,12c-0.553,0-1,0.447-1,1s0.447,1,1,1c4.418,0,8,3.582,8,8c0,0.553,0.447,1,1,1s1-0.447,1-1
                        C58,16.478,53.522,12,48,12z"></path>
                            </g>
                        </svg></button>

                    <div class="right-buttons">
                        <button id="likeButton" class="button" title="Like"><svg
                                class="w-6 h-6 hover:stroke-primary hover:fill-primary" version="1.0" id="Layer_1"
                                xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"
                                viewBox="0 0 64 64" enable-background="new 0 0 64 64" xml:space="preserve"
                                stroke="currentColor">
                                <g>
                                    <path fill="#231F20" d="M64,28c0-3.314-2.687-6-6-6H41l0,0h-0.016H41l2-18c0.209-2.188-1.287-4-3.498-4h-4.001
                        C33,0,31.959,1.75,31,4l-8,18c-2.155,5.169-5,6-7,6c-1,0-2,0-2,0v-2c0-2.212-1.789-4-4-4H4c-2.211,0-4,1.788-4,4v34
                        c0,2.21,1.789,4,4,4h6c2.211,0,4-1.79,4-4v-2c1,0,3.632,0.052,6.21,2.697C23.324,63.894,27.043,64,29,64h23c3.313,0,6-2.688,6-6
                        c0-1.731-0.737-3.288-1.91-4.383C58.371,52.769,60,50.577,60,48c0-1.731-0.737-3.288-1.91-4.383C60.371,42.769,62,40.577,62,38
                        c0-1.731-0.737-3.288-1.91-4.383C62.371,32.769,64,30.577,64,28z M12,60c0,1.104-0.896,2-2,2H4c-1.104,0-2-0.896-2-2V26
                        c0-1.105,0.896-2,2-2h6c1.104,0,2,0.895,2,2V60z M58,32H48c-0.553,0-1,0.446-1,1c0,0.552,0.447,1,1,1h8c2.209,0,4,1.79,4,4
                        c0,2.209-1.791,4-4,4H46c-0.553,0-1,0.446-1,1c0,0.552,0.447,1,1,1h8c2.209,0,4,1.79,4,4c0,2.209-1.791,4-4,4H44
                        c-0.553,0-1,0.446-1,1c0,0.552,0.447,1,1,1h8c2.209,0,4,1.79,4,4c0,2.209-1.791,4-4,4H29c-1,0-4.695,0.034-7.358-2.699
                        C18.532,56.109,16.112,56.003,14,56V30h2c4,0,6.57-1.571,9.25-8L33,4c0.521-1.104,1.146-2,2.251-2H39c1.104,0,2.126,0.834,2,2
                        l-1.99,18c-0.132,1.673,0.914,2,1.99,2h17c2.209,0,4,1.79,4,4C62,30.209,60.209,32,58,32z"></path>
                                    <path fill="#231F20" d="M7,54c-1.657,0-3,1.342-3,3c0,1.656,1.343,3,3,3s3-1.344,3-3C10,55.342,8.657,54,7,54z M7,58
                        c-0.553,0-1-0.449-1-1c0-0.553,0.447-1,1-1s1,0.447,1,1C8,57.551,7.553,58,7,58z"></path>
                                </g>
                            </svg></button>
                        <button id="dislikeButton" class="button" title="Dislike"><svg
                                class="w-6 h-6 hover:stroke-primary hover:fill-primary" version="1.0" id="Layer_1"
                                xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"
                                viewBox="0 0 64 64" enable-background="new 0 0 64 64" xml:space="preserve"
                                stroke="currentColor">
                                <g>
                                    <path d="M64,36c0,3.312-2.687,6-6,6H41v-0.002L40.984,42H41l2,18c0.209,2.186-1.287,4-3.498,4h-4.001
                        C33,64,31.959,62.248,31,60l-8-18c-2.155-5.171-5-6-7-6c-1,0-2,0-2,0v2c0,2.21-1.789,4-4,4H4c-2.211,0-4-1.79-4-4V4
                        c0-2.212,1.789-4,4-4h6c2.211,0,4,1.788,4,4v2c1,0,3.632-0.054,6.21-2.699C23.324,0.104,27.043,0,29,0h23c3.313,0,6,2.686,6,6
                        c0,1.729-0.737,3.286-1.91,4.381C58.371,11.229,60,13.421,60,16c0,1.729-0.737,3.286-1.91,4.381C60.371,21.229,62,23.421,62,26
                        c0,1.729-0.737,3.286-1.91,4.381C62.371,31.229,64,33.421,64,36z M12,4c0-1.105-0.896-2-2-2H4C2.896,2,2,2.895,2,4v34
                        c0,1.104,0.896,2,2,2h6c1.104,0,2-0.896,2-2V4z M58,32H48c-0.553,0-1-0.448-1-1c0-0.554,0.447-1,1-1h8c2.209,0,4-1.791,4-4
                        c0-2.21-1.791-4-4-4H46c-0.553,0-1-0.448-1-1c0-0.554,0.447-1,1-1h8c2.209,0,4-1.791,4-4c0-2.21-1.791-4-4-4H44
                        c-0.553,0-1-0.448-1-1c0-0.554,0.447-1,1-1h8c2.209,0,4-1.791,4-4c0-2.21-1.791-4-4-4H29c-1,0-4.695-0.036-7.358,2.697
                        C18.532,7.889,16.112,7.995,14,8v26h2c4,0,6.57,1.569,9.25,8L33,60c0.521,1.103,1.146,2,2.251,2H39c1.104,0,2.126-0.834,2-2
                        l-1.99-18c-0.132-1.675,0.914-2,1.99-2h17c2.209,0,4-1.791,4-4C62,33.79,60.209,32,58,32z"></path>
                                    <path d="M7,38c-1.657,0-3-1.344-3-3c0-1.658,1.343-3,3-3s3,1.342,3,3C10,36.656,8.657,38,7,38z M7,34
                        c-0.553,0-1,0.447-1,1c0,0.551,0.447,1,1,1s1-0.449,1-1C8,34.447,7.553,34,7,34z"></path>
                                </g>
                            </svg></button>

                    </div>
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
                        <h2 class="breed-name">
                        <span id="breedName" class="breed-name">Breed Name</span>
                        <span id="breedOrigin" class="breed-origin"></span>
                        <span class="breed-id" id="breedId"></span>
                    </h2>
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
                <div class="layout-controls d-flex gap-3 mb-3">
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
                    <button class="delete-btn">×</button>
                    <!-- Images will be added here -->
                </div>
            </div>
        </div>
    </div>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js"></script>
</body>

</html>