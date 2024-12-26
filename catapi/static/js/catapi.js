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
  const breedId= document.getElementById("breedId");
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
  let lastImageId = ""; // Keep track of the last displayed image ID

  // Add a data attribute to store current image ID
  if (!catImage.hasAttribute('data-image-id')) {
    catImage.setAttribute('data-image-id', '');
  }

  votingButton.addEventListener("click", function () {
    // Add active class to the voting button (Bootstrap's active class)
    votingButton.classList.add("active");
    fetchNewCatImage();  // Fetch new images when the voting layout is shown
  });

  async function fetchNewCatImage() {
    const loadingImage = document.getElementById("loadingImage");
  
    try {
      loadingImage.style.display = "block";
      catImage.style.display = "none";

      const response = await fetch("/");
      const html = await response.text();
  
      // Parse the HTML to get the new image URL
      const parser = new DOMParser();
      const doc = parser.parseFromString(html, "text/html");
      const newImageSrc = doc.getElementById("catImage").src;
  
      // Extract the new image ID from the URL
      const newImageId = newImageSrc.split("/").pop().split(".")[0];
  
      // Check if the new image is the same as the current one
      if (newImageId === lastImageId) {
        console.warn("Fetched image is the same as the current image. Fetching again...");
        return fetchNewCatImage(); // Retry fetching a new image
      }
  
      console.log("Fetched new image:", newImageSrc);
      console.log("New image ID:", newImageId);
  
      // Preload the new image before displaying it
      const img = new Image();
      img.src = newImageSrc;
  
      img.onload = () => {
        // Update the lastImageId to the new image ID
        lastImageId = newImageId;
        loadingImage.style.display = "none";
        // Apply fade-out animation to the container
        catImageContainer.classList.add("fade-out");
  
        setTimeout(() => {
          // Update the image and remove the fade-out class
          catImage.src = newImageSrc;
          catImage.setAttribute("data-image-id", newImageId);
          catImage.style.display = "block";
          catImageContainer.classList.remove("fade-out");
          catImageContainer.classList.add("fade-in");
          setTimeout(() => {
            catImageContainer.classList.remove("fade-in");
          }, 300);
        }, 300); // Matches the fade-out animation duration
      };
  
      img.onerror = () => {
        console.error("Failed to preload the image. Retrying...");
        fetchNewCatImage(); // Retry fetching if the image fails to load
      };
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
    // Clear the container before appending new images
    favoriteImagesContainer.innerHTML = "";

    favoriteImages.forEach((favorite) => {
      if (favorite.image && favorite.image.url) {
        // Create a wrapper div for the image and delete button
        const wrapper = document.createElement("div");
        wrapper.className = "position-relative d-inline-block m-2";

        // Create an img element for the cat image
        const img = document.createElement("img");
        img.src = favorite.image.url;
        img.className = "favorite-image";
        img.alt = "Cat image";

        // Create a delete button for each favorite
        const deleteBtn = document.createElement("button");
        deleteBtn.className = "delete-btn";
        deleteBtn.textContent = "×";

        // Delete favorite image on button click
        deleteBtn.onclick = async () => {
          try {
            const response = await fetch(`/deleteFavorite/${favorite.id}`, {
              method: "DELETE",
            });

            if (response.ok) {
              wrapper.remove(); // Remove the image and button from DOM
            } else {
              alert("Failed to delete favorite");
            }
          } catch (error) {
            console.error("Error:", error);
            alert("Failed to delete favorite");
          }
        };

        // Append the image and delete button to the wrapper
        wrapper.appendChild(img);
        wrapper.appendChild(deleteBtn);

        // Append the wrapper to the container
        favoriteImagesContainer.appendChild(wrapper);
      }
    });
  }

  // Function to fetch favorite images from the API
  async function fetchFavoriteImages() {
    try {
      const response = await fetch("/getFavorites");
      const data = await response.json();

      if (response.ok) {
        // Display favorite images
        displayFavoriteImages(data);
      } else {
        console.error("Failed to fetch favorites:", data.error);
      }
    } catch (error) {
      console.error("Error fetching favorite images:", error);
    }
  }
  // Show the favorite layout
  window.showFavoriteLayout = async function () {
    try {
      // Fetch favorite images from the API
      const response = await fetch("/getFavorites");
      const data = await response.json();

      if (response.ok && data.length > 0) {
        // Display favorite images
        displayFavoriteImages(data);

        // Show the favorite layout and hide others
        document.getElementById("favoriteLayout").style.display = "block";
        document.getElementById("votingLayout").style.display = "none";
        document.getElementById("breedLayout").style.display = "none";
      } else {
        // Handle case where there are no favorite images
        favoriteImagesContainer.innerHTML = "<p>No favorite images yet.</p>";
      }
    } catch (error) {
      console.error("Error fetching favorite images:", error);
      alert("Failed to load favorite images. Please try again.");
    }
  };

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
      await fetchFavoriteImages(); // Update favorite images list
    } catch (error) {
      console.error("Error adding to favorites:", error);
    }
  }




  window.showVotingLayout = function () {
    document.getElementById("votingLayout").style.display = "block";
    document.getElementById("breedLayout").style.display = "none";
    document.getElementById("favoriteLayout").style.display = "none";

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

    breedId.textContent = breed.id;
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
    }, 4000); // Change slide every 4 seconds
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
