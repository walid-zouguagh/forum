import { createPosts } from "./posts";

async function PostCategory() {
  let category = document.getElementById("category");
  try {
    const res = await fetch("/api/category/list");
    const data = await res.json();
    data.forEach((catg) => {
      category.innerHTML += `
        <input type="checkbox" name="options" id="" value="${catg.Name}" data-name="${catg.Name}">${catg.Name}<br>
        `;
    });
  } catch {
    console.log("erroure");
  }
}
PostCategory();

document
  .getElementById("submit")
  .addEventListener("click", async function (event) {
    event.preventDefault(); // Prevent the form from submitting
    // console.log("ok");
    console.log(event);

    // Declare validation flags
    let isValidTitle = false;
    let isValidContent = false;
    let isValidCheckboxes = false;

    // Get form values
    let Title = document.getElementById("Title").value.trim();
    let Content = document.getElementById("Content").value.trim();
    let categoryName = [];

    // Validate Title
    if (Title === "") {
      document.getElementById("errorTitle").innerHTML = "Title is required";
      document.getElementById("errorTitle").style.color = "red";
      isValidTitle = false;
    } else {
      document.getElementById("errorTitle").innerHTML = "";
      isValidTitle = true;
    }

    // Validate Content
    if (Content === "") {
      document.getElementById("errorContent").innerHTML = "Content is required";
      document.getElementById("errorContent").style.color = "red";
      isValidContent = false;
    } else {
      document.getElementById("errorContent").innerHTML = "";
      isValidContent = true;
    }

    // Get selected checkboxes
    let checkboxes = document.querySelectorAll('input[name="options"]:checked');
    checkboxes.forEach((checkbox) => {
      categoryName.push(checkbox.getAttribute("data-name"));
    });

    // Checkbox validation
    if (categoryName.length === 0) {
      document.getElementById("errorcategory").innerHTML =
        "Please select at least one category";
      document.getElementById("errorcategory").style.color = "red";
      isValidCheckboxes = false;
    } else {
      document.getElementById("errorcategory").innerHTML = "";
      isValidCheckboxes = true;
    }

    // If all validations are successful, submit the form
    if (isValidTitle && isValidContent && isValidCheckboxes) {
      try {
        const res = await fetch("/add-post", {
          method: "POST",
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
          body: new URLSearchParams({
            Title: Title,
            Content: Content,
            options: categoryName.join(","), // Serialize as a comma-separated string
          }),
        });

        if (res.ok) {
          const post = await res.json();
          createPosts(post, true);
        } else {
          alert("Failed to submit post. Please try again later.");
          console.error("Response status:", res.status);
        }
      } catch (error) {
        alert("An error occurred: " + error.message);
        console.error("Fetch error:", error);
      }
    }
  });
