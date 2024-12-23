document.addEventListener("DOMContentLoaded", () => {
  const votingButton = document.getElementById("votingButton");
  const breedButton = document.getElementById("breedButton");
  const favButton = document.getElementById("favButton");
  const heartButton = document.getElementById("heartButton");
  const catImageContainer = document.getElementById("catImageContainer");
  const catImage = document.getElementById("catImage");
  const gridBtn = document.getElementById("gridBtn");
  const listBtn = document.getElementById("listBtn");
  const favoriteImagesContainer = document.getElementById("favoriteImagesContainer");
  const likeButton = document.getElementById("likeButton");
  const dislikeButton = document.getElementById("dislikeButton");
  const breedSearch = document.getElementById("breedSearch");
  const breedList = document.getElementById("breedList");
  const breedName = document.getElementById("breedName");
  const breedOrigin = document.getElementById("breedOrigin");
  const breedDescription = document.getElementById("breedDescription");
  const breedWiki = document.getElementById("breedWiki");
  const breedImagesContainer = document.getElementById("breedImagesContainer");
  const breedImagesSlider = document.getElementById("breedImagesSlider");
  const sliderIndicators = document.getElementById("sliderIndicators");

  let currentBreeds = [];
  let favoriteImages = JSON.parse(localStorage.getItem("favoriteImages")) || [];
  let currentSlideIndex = 0;
  let slideInterval;
  let isTransitioning = false;
  let catImages = [];

  // Add a data attribute to store current image ID
  if (!catImage.hasAttribute('data-image-id')) {
    catImage.setAttribute('data-image-id', '');
  }

  votingButton.addEventListener("click", function() {
    // Add active class to the voting button (Bootstrap's active class)
    votingButton.classList.add("active");
    fetchNewCatImage();  // Fetch new images when the voting layout is shown
  });

  async function fetchNewCatImage() {
    try {
      const response = await fetch("/");
      const html = await response.text();

      // Create a temporary div to parse the HTML
      const parser = new DOMParser();
      const doc = parser.parseFromString(html, 'text/html');
      const newImageSrc = doc.getElementById("catImage").src;

      console.log("Fetched new image:", newImageSrc);

      // Extract image ID from URL
      const imageId = newImageSrc.split('/').pop().split('.')[0];
      console.log("New image ID:", imageId);

      catImageContainer.classList.add("fade-out");
      setTimeout(() => {
        catImage.src = newImageSrc;
        catImage.setAttribute('data-image-id', imageId);
        catImageContainer.classList.remove("fade-out");
      }, 500);
    } catch (error) {
      console.error("Error fetching image:", error);
    }
  }


  async function createVote(value) {
    const imageId = catImage.getAttribute('data-image-id');
    console.log("Attempting to vote for image:", imageId, "with value:", value);

    if (!imageId) {
      console.error("No image ID available");
      return;
    }

    try {
      console.log("Sending vote request...");
      const response = await fetch("/vote", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          image_id: imageId,
          value: value
        })
      });

      console.log("Vote response status:", response.status);
      const data = await response.json();
      console.log("Vote response data:", data);

      if (data.error) {
        console.error("Error creating vote:", data.error);
        return;
      }

      console.log("Vote successful, fetching new image...");
      // Fetch new image after successful vote
      fetchNewCatImage();
    } catch (error) {
      console.error("Error creating vote:", error);
    }
  }


  likeButton.addEventListener("click", () => {
    console.log("Like button clicked");
    createVote(1);
  });

  dislikeButton.addEventListener("click", () => {
    console.log("Dislike button clicked");
    createVote(-1);
  });

  // Initial image ID setup
  const initialImageSrc = catImage.src;
  if (initialImageSrc) {
    const imageId = initialImageSrc.split('/').pop().split('.')[0];
    catImage.setAttribute('data-image-id', imageId);
    console.log("Initial image ID set to:", imageId);
  }

  // Heart Button Functionality
  heartButton.addEventListener("click", async () => {
    const imageUrl = document.getElementById("catImage").src;
    const imageId = catImage.getAttribute("data-image-id");

    if (imageUrl) {
      // Add to API favorites
      await addToFavorites(imageId);

      // Fetch updated favorite images from the API and display them
      await fetchFavoriteImages();

      // Fetch a new cat image to display
      fetchNewCatImage();
    }
  });

  // Event listeners
  gridBtn.addEventListener("click", () => switchLayout("grid"));
  listBtn.addEventListener("click", () => switchLayout("list"));
  function switchLayout(type) {
    // Remove both classes first
    favoriteImagesContainer.className = `${type}-layout`;

    // Update button states
    if (type === "list") {
      listBtn.classList.add("active");
      gridBtn.classList.remove("active");
    } else {
      gridBtn.classList.add("active");
      listBtn.classList.remove("active");
    }

    // Adjust image sizes based on layout
    const images = favoriteImagesContainer.getElementsByClassName('favorite-image');
    Array.from(images).forEach(img => {
      if (type === 'list') {
        img.style.maxHeight = '80vh';
      } else {
        img.style.maxHeight = '200px';
      }
    });
  }
  // Frontend JavaScript
  function displayFavoriteImages(favoriteImages) {
    favoriteImagesContainer.innerHTML = ""; // Clear the container

    favoriteImages.forEach((favorite) => {
      if (favorite.image && favorite.image.url) {
        // Create wrapper div for image and delete button
        const wrapper = document.createElement("div");
        wrapper.className = "position-relative d-inline-block m-2";

        // Create image element
        const img = document.createElement("img");
        img.src = favorite.image.url;
        img.className = "favorite-image";
        img.alt = "Cat image";

        // Create delete button
        const deleteBtn = document.createElement("button");
        deleteBtn.className = "delete-btn";
        deleteBtn.innerHTML = "×";
      

        // Add click handler for delete
        deleteBtn.onclick = async () => {
          try {
            const response = await fetch(`/deleteFavorite/${favorite.id}`, {
              method: 'DELETE'
            });

            if (response.ok) {
              wrapper.remove(); // Remove the image and button from DOM
            } else {
              alert('Failed to delete favorite');
            }
          } catch (error) {
            console.error('Error:', error);
            alert('Failed to delete favorite');
          }
        };

        // Add image and button to wrapper
        wrapper.appendChild(img);
        wrapper.appendChild(deleteBtn);
        favoriteImagesContainer.appendChild(wrapper);
      }
    });
  }
  

  // Function to add an image to favorites via the API
  async function addToFavorites(imageId) {
    try {
      const response = await fetch("/createFavorite", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          image_id: imageId,
        }),
      });

      const data = await response.json();
      if (!response.ok) {
        throw new Error(data.error || "Failed to add to favorites");
      }

      console.log("Image added to favorites successfully");
    } catch (error) {
      console.error("Error adding to favorites:", error);
    }
  }

  // Function to fetch the favorite images from the API
  async function fetchFavoriteImages() {
    try {
      const response = await fetch("/getFavorites");
      const data = await response.json();

      if (response.ok) {
        // Display the favorite images
        displayFavoriteImages(data);
      } else {
        console.error("Failed to fetch favorites:", data.error);
      }
    } catch (error) {
      console.error("Error fetching favorite images:", error);
    }
  }


  window.showVotingLayout = function () {
    document.getElementById("votingLayout").style.display = "block";
    document.getElementById("breedLayout").style.display = "none";
    document.getElementById("favoriteLayout").style.display = "none";
    
  };

  window.showFavoriteLayout = async function () {
    try {
      // Fetch favorite images from the API
      const response = await fetch("/getFavorites");
      const data = await response.json();

      if (response.ok && data.length > 0) {
        // Clear the container before adding new images
        favoriteImagesContainer.innerHTML = "";
        displayFavoriteImages(data);
        // Iterate over the favorite images and create img elements
        data.forEach((favorite) => {
          if (favorite.image && favorite.image.url) {
            const img = document.createElement("img");
            img.src = favorite.image.url; // Use the image URL from the API response
            img.alt = "Favorite Cat Image";
            img.classList.add("favorite-image");
            favoriteImagesContainer.appendChild(img);
          }
        });

        // Show the favorite layout and hide others
        document.getElementById("favoriteLayout").style.display = "block";
        document.getElementById("votingLayout").style.display = "none";
        document.getElementById("breedLayout").style.display = "none";
      } else {
        // Handle case where there are no favorite images
        favoriteImagesContainer.innerHTML = "";
        alert("No favorite images yet.");
      }
    } catch (error) {
      console.error("Error fetching favorite images:", error);
      alert("Failed to load favorite images. Please try again.");
    }
  };

  async function fetchBreeds() {
    try {
      const response = await fetch("/api/breeds");
      const breeds = await response.json();
      currentBreeds = breeds;
      return breeds;
    } catch (error) {
      console.error("Error fetching breeds:", error);
      return [];
    }
  }

  async function fetchBreedImages(breedId) {
    try {
      const response = await fetch(`/api/breed-images?breed_id=${breedId}`);
      const images = await response.json();
      return images;
    } catch (error) {
      console.error("Error fetching breed images:", error);
      return [];
    }
  }

  async function initializeBreedSearch() {
    const breeds = await fetchBreeds();
    breedList.innerHTML = "";

    breeds.forEach((breed) => {
      const breedItem = document.createElement("div");
      breedItem.className = "breed-item";
      breedItem.textContent = breed.name;
      breedItem.onclick = () => selectBreed(breed);
      breedList.appendChild(breedItem);
    });

    // Auto-select the first breed
    if (breeds.length > 0) {
      selectBreed(breeds[0]);
    }
  }

  async function selectBreed(breed) {
    breedSearch.value = breed.name; // Set breed name in the input field
    breedList.style.display = "none"; // Close the dropdown

    breedName.textContent = breed.name;
    breedDescription.textContent = breed.description;
    breedOrigin.textContent = breed.origin;
    breedWiki.href = breed.wikipedia_url || "#";
    breedWiki.style.display = breed.wikipedia_url ? "block" : "none";

    const images = await fetchBreedImages(breed.id);
    setupSlider(images, breed.name);
  }

  function setupSlider(images, breedName) {
    const sliderWrapper = document.createElement("div");
    sliderWrapper.className = "slider-wrapper";
    breedImagesSlider.innerHTML = "";
    sliderIndicators.innerHTML = "";
    currentSlideIndex = 0;

    if (slideInterval) {
      clearInterval(slideInterval);
    }

    if (images.length > 0) {
      images.forEach((image, index) => {
        const img = document.createElement("img");
        img.src = image.url;
        img.alt = `${breedName} image ${index + 1}`;
        img.className = `breed-image ${index === 0 ? "active" : ""}`;
        sliderWrapper.appendChild(img);

        const indicator = document.createElement("div");
        indicator.className = `slider-indicator ${index === 0 ? "active" : ""}`;
        indicator.onclick = () => {
          if (!isTransitioning) {
            goToSlide(index);
          }
        };
        sliderIndicators.appendChild(indicator);
      });

      breedImagesSlider.appendChild(sliderWrapper);

      startAutoSlide(images.length);
    }
  }

  function startAutoSlide(totalSlides) {
    slideInterval = setInterval(() => {
      const nextIndex = (currentSlideIndex + 1) % totalSlides;
      goToSlide(nextIndex);
    }, 2000); // Change slide every 3 seconds
  }

  function goToSlide(index) {
    if (isTransitioning || index === currentSlideIndex) return;

    isTransitioning = true;

    const images = document.querySelectorAll(".breed-image");
    const indicators = document.querySelectorAll(".slider-indicator");

    images[currentSlideIndex].classList.remove("active");
    indicators[currentSlideIndex].classList.remove("active");

    images[index].classList.add("active");
    indicators[index].classList.add("active");

    currentSlideIndex = index;

    startAutoSlide(images.length);

    setTimeout(() => {
      isTransitioning = false;
    }, 500);
  }

  breedSearch.addEventListener("input", (e) => {
    const searchTerm = e.target.value.toLowerCase();
    const filteredBreeds = currentBreeds.filter((breed) =>
      breed.name.toLowerCase().includes(searchTerm)
    );

    breedList.innerHTML = "";
    filteredBreeds.forEach((breed) => {
      const breedItem = document.createElement("div");
      breedItem.className = "breed-item";
      breedItem.textContent = breed.name;
      breedItem.onclick = () => selectBreed(breed);
      breedList.appendChild(breedItem);
    });

    breedList.style.display = filteredBreeds.length > 0 ? "block" : "none";
  });

  breedSearch.addEventListener("focus", () => {
    breedList.style.display = "block";
  });

  document.addEventListener("click", (e) => {
    if (!breedSearch.contains(e.target) && !breedList.contains(e.target)) {
      breedList.style.display = "none";
    }
  });

  function showBreedLayout() {
    document.getElementById("votingLayout").style.display = "none";
    document.getElementById("favoriteLayout").style.display = "none";
    document.getElementById("breedLayout").style.display = "block";
    initializeBreedSearch();
  }

  window.showBreedLayout = showBreedLayout;

  showVotingLayout();
});
