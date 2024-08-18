import { renderRawJson, renderLinkPreviews, handleViewToggle } from "./modules.js";

document.getElementById("data-form").addEventListener("submit", function (event) {
   event.preventDefault();
   const form = event.target;
   const formData = new FormData(form);
   const outputType = formData.get("outputType");
   const links = formData.get("links");
   const filename = formData.get("filename");
   const actionUrl = outputType === "zip" ? "/api/v1/zip" : "/api/v1/json";

   // Construct query parameters
   const queryParams = new URLSearchParams({
      url: links,
      filename: filename,
   }).toString();

   // Set URL with query parameters
   const urlWithParams = `${actionUrl}?${queryParams}`;

   fetch(urlWithParams)
      .then((response) => {
         if (response.ok) {
            if (outputType === "json") {
               return response.json();
            } else {
               // For ZIP file, handle as blob
               return response.blob();
            }
         } else {
            throw new Error("Network response was not ok.");
         }
      })
      .then((data) => {
         const resultDiv = document.getElementById("result");
         if (outputType === "json") {
            // Display JSON data and clear previous content
            resultDiv.innerHTML = "";
            resultDiv.innerHTML = JSON.stringify(data);
            resultDiv.dataset.jsonData = JSON.stringify(data); // Store data for later use
            handleViewToggle(); // Update view based on checkbox
        } else {
            // Handle file download
            const url = window.URL.createObjectURL(data);
            const a = document.createElement("a");
            a.href = url;
            a.download = filename ? `${filename}.zip` : "data.zip"; // Set default file name for ZIP
            document.body.appendChild(a);
            a.click();
            a.remove();
            window.URL.revokeObjectURL(url);
         }
      })
      .catch((error) => {
         console.error("Error:", error);
      });
});

document.addEventListener("DOMContentLoaded", function () {
   
   const viewToggle = document.getElementById("viewToggle");
   
   // Add event listener for checkbox change
   viewToggle.addEventListener("change", handleViewToggle);
});
