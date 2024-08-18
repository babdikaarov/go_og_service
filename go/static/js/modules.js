// Function to render raw JSON
export function renderRawJson(data) {
   const resultDiv = document.getElementById("result");
   resultDiv.innerHTML = "";
   resultDiv.textContent = JSON.stringify(data, null, 2);
}

// Function to render link previews
export function renderLinkPreviews(data) {
   const resultDiv = document.getElementById("result");
   resultDiv.innerHTML = "";
   resultDiv.innerHTML = data
      .map(
         (item) => `
        <div class="card_container">
            <div class="card_container_image">
                <img
                src="${item.image}"
                alt="Image"
                class="card_image"
                />
            </div>
            <div class="card_container_content">
                <h2 class="card_title">${item.title}</h2>
                <p class="card_description">${item.description}</p>
                <div class="card_container_footer">
                <img src="${item.icon}" alt="Icon" class="card_icon" />
                <a
                    href="${item.original_url}"
                    target="_blank"
                    class="card_link"
                >
                    ${item.original_url}
                </a>
                </div>
            </div>
        </div>
       
    `,
      )
      .join("");
}

// Function to handle view toggle
export function handleViewToggle() {
   const resultDiv = document.getElementById("result");
   const jsonData = resultDiv.dataset.jsonData;
   if (jsonData) {
      const data = JSON.parse(jsonData);
      if (document.getElementById("viewToggle").checked) {
         resultDiv.classList.remove("nowrap");
         resultDiv.classList.add("wrap");
         renderRawJson(data);
      } else {
         resultDiv.classList.remove("wrap");
         resultDiv.classList.add("nowrap");
         renderLinkPreviews(data.data);
      }
   }
}
