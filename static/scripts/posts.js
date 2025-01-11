import { CreatPost, HandulLike } from "./like.js";

export async function addEventOnPosts(path) {
  document.addEventListener("DOMContentLoaded", function () {
    const postsContainer = document.getElementById("posts-container");
    PostCategory();

    postsContainer.addEventListener("submit", function (event) {
      const postElement = event.target.closest(".post-container");
      if (event.target.classList.contains("comment_form")) {
        event.preventDefault();
        const form = event.target;

        const postId = postElement.getAttribute("data-post-id");
        const commentText = form.querySelector(".comment").value.trim();

        if (commentText === "") {
          alert("Comment cannot be empty.");
          return;
        }

        handleComment(postId, commentText);
        form.reset();
      }
    });

    postsContainer.addEventListener("click", function (event) {
      const postElement = event.target.closest(".post-container");

      const postId = postElement.getAttribute("data-post-id");

      const CommentClass = event.target.classList.contains("see_comments");

      if (event.target.classList.contains("see_comments")) {
        event.target.disabled = true;
      }
      const divcomments = document.querySelector(".divcomments" + postId);
      if (CommentClass && postId) {
        GetComments(postId, divcomments);
      }
    });

    loadMorePosts(path);
    window.addEventListener("scrollend", () => {
      handleScroll(path);
    });
  });
}

function handleComment(postId, comment) {
  fetch(`/api/${postId}/comment/new`, {
    method: "POST",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    body: `comment=${comment}`,
  })
    .then((response) => {
      if (!response.ok) {
        alert("faild to add Comment");
        return;
      }
      return response.json();
    })
    .then((data) => {
      let comment_form = ".divcomments" + postId;
      let commentElement = document.querySelector(comment_form);
      GetComments(postId, commentElement);
    })
    .catch((error) => window.location.href = "/login");
}

let offset = 0;
const limit = 20;
let loading = false;
const onehour = 1000 * 60 * 60;
const oneday = 1000 * 60 * 60 * 24;
const onemin = 1000 * 60;

export async function loadMorePosts(path) {
  if (loading) return;
  loading = true;

  try {
    const response = await fetch(`${path}?offset=${offset}`);
    const posts = await response.json();
    console.log(posts);

    if (!posts || posts.length === 0) return;
    createPosts(posts);

    offset += limit;
  } catch (error) {
    console.error("Error loading posts:", error);
  } finally {
    loading = false;
  }
}

export function createPosts(posts) {

  const postsContainer = document.getElementById("posts-container");
  posts.forEach((post) => {
    const postElement = createEle("div", "post-container");
    postElement.dataset.postId = post.Id;

    /* h2 will contain the image and name of the persen who posted */
    const posterName = createEle("h2", "poster");
    const posterImg = createEle("img");
    posterImg.src =
      "/static/images/466006304_871124095226532_8631138819273739648_n.jpg";
    const nameContainer = createEle("span", "usrname");
    nameContainer.innerText = post.UserName;
    const createdat = createEle("span", "creationdate");

    createdat.innerText = formDate(post.Date);
    posterName.appendChild(posterImg);
    posterName.appendChild(nameContainer);
    postElement.appendChild(posterName);
    postElement.appendChild(createdat);

    /* creating a div that will contain all the elements bellow */
    const pc = createEle("div", "pc");

    /* creating an h3 element to contain the post title */
    const title = createEle("h3", "title");
    title.innerText = post.Title;

    /* creating a p element that will contain the content of the post */
    const content = createEle("p", "content");
    content.innerText = post.Content;



    const Image = createEle("img");
    Image.className = "image";
    if (post.Image != "") {
      const imageUrl = `http://localhost:8001/${post.Image}`;
      fetch(imageUrl)
        .then(response => {
          if (response.ok) {
            Image.src = imageUrl;
          }else {
            console.error(`Image not found at ${imageUrl}`);
          
           
          }
        })
    }
    // Optionally, log the image path for debugging
    console.log(post.Image);
    // Append the image to the parent element
    pc.append(Image);



    const categories_container = createEle("div", "categories");

    for (let cate of post.Categories) {
      const span = createEle("a", "category");
      span.href = `/category/${cate}`;
      span.innerText = cate;
      categories_container.appendChild(span);
    }
    /* creating like and dislike button */
    const like_dislike_container = createEle("div", "like-dislike-container");
    //

    const [likebnt, dislikebnt, likeNbr, dislikeNbr] = HandulLike(
      post.Reactions.Action,
      post.Reactions.Likes,
      post.Reactions.Dislikes,
      "posts",
      post.Id
    );

    likebnt.classList.add("like-btn");

    /* create an img element to contain like icon */
    const likeIcon = createEle("img", "likeNbr");
    likeIcon.className = "likeIcon";
    likeIcon.src = "/static/images/like.png";

    // const likeNbr = createEle("span");
    // likeNbr.innerText = post.Reactions.Likes;
    likeNbr.className = "likeNbr";
    likebnt.appendChild(likeIcon);
    // likebnt.appendChild(likeNbm);

    /* creationg of the dislike button */
    // const dislikebnt = createEle("button");

    dislikebnt.classList.add("dislike-btn");
    /* creating an img tag to containg dislike icon */
    const dislikeIcon = createEle("img", "dislikeIcon");
    dislikeIcon.src = "/static/images/dislike.png";

    // const dislikeNbr = createEle("span");
    dislikeNbr.className = "dislikeNbr";
    // dislikeNbr.innerText = post.Reactions.Dislikes;

    dislikebnt.appendChild(dislikeIcon);
    // dislikebnt.appendChild(dislikeNbr);

    /* appending like and dislike buttons to like container */
    like_dislike_container.append(likebnt, likeNbr, dislikebnt, dislikeNbr);

    /* appending like container to the post contaner */

    ////add div comments
    const divcomments = createEle("div", `divcomments${post.Id} divcomments`);

    /* adding a button to see comments */
    const seecomments = createEle("button", "see_comments");
    seecomments.innerText = "see comments";

    /* creating the form that sends comments */
    const comment_form = createEle("form", "comment_form");
    comment_form.method = "POST";

    const title_impt = createEle("input", "comment");
    title_impt.name = "comment";
    title_impt.type = "text";
    title_impt.placeholder = "Add your comment";
    title_impt.required = true;

    const submit_comment = createEle("button", "send_comment");
    submit_comment.type = "submit";

    const send_icon = createEle("img", "sendimg");
    send_icon.src = "/static/images/send-message.png";
    submit_comment.appendChild(send_icon);
    comment_form.append(title_impt, submit_comment);

    pc.append(
      title,
      content,
      categories_container,
      like_dislike_container,
      divcomments,
      seecomments,
      comment_form
    );
    postElement.appendChild(pc);

    postsContainer.appendChild(postElement);
  });
}

export function formDate(date) {
  let CreattionDate = new Date(date).getTime();
  const currentTime = Date.now();
  const elapsed = currentTime - CreattionDate;

  const days = Math.floor(elapsed / oneday);
  const hours = Math.floor((elapsed % oneday) / onehour);
  const minutes = Math.floor((elapsed % onehour) / onemin);

  let timeText = "";

  if (days > 0) {
    timeText += `${days}d `;
  }
  if (hours > 0) {
    timeText += `${hours}h `;
  }
  if (minutes > 0 && (days == 0 || hours == 0)) {
    timeText += `${minutes}min`;
  }
  return timeText;
}

export function createEle(elename, className) {
  const ele = document.createElement(elename);
  ele.className = className;
  return ele;
}

const showPostFormButton = document.querySelector(".show-postForm");
const postForm = document.querySelector(".postForm");
const layout = document.querySelector(".lay-out"); // Optional dimmed background

// Show the form
if (showPostFormButton) {
  showPostFormButton.addEventListener("click", () => {
    postForm.style.display = "flex";
    layout.style.display = "block";
    document.body.style.overflow = "hidden";

    // Hide the form when clicking outside or on a cancel button
    layout.addEventListener("click", () => {
      postForm.style.display = "none";
      layout.style.display = "none";
      document.body.style.overflow = "scroll";
    });
  });
}

export function handleScroll(path) {
  const scrollPosition = window.scrollY + window.innerHeight;
  const threshold = document.body.scrollHeight - 1000;

  if (scrollPosition > threshold) {
    loadMorePosts(path);
  }
}

const form = document.getElementById("postForm");
if (form) {
  form.addEventListener("submit", async function (event) {
    event.preventDefault();
    // Declare validation flags
    const postsContainer = document.getElementById("posts-container");
    let isValidTitle = false;
    let isValidContent = false;
    let isValidCheckboxes = false;

    // Get form values
    let Title = document.getElementById("post").value;
    let Content = document.getElementById("content").value;
    let Image = document.querySelector(".Image");
    const file = Image.files[0];

    let categoryName = [];

    // Get selected checkboxes
    let checkboxes = document.querySelectorAll('input[name="options"]:checked');


    checkboxes.forEach((checkbox) => {
      categoryName.push(checkbox.getAttribute("data-name"));
    });

    if (Title === "") {
      document.getElementById("errorTitle").innerHTML = "Title is required";
      document.getElementById("errorTitle").style.color = "red";
      isValidTitle = false;
    } else if (Title.length > 100) {
      document.getElementById("errorTitle").innerHTML = "max title is 100";
      document.getElementById("errorTitle").style.color = "red";
      isValidTitle = false;
    } else {
      document.getElementById("errorTitle").innerHTML = "";
      isValidTitle = true;
    }

    if (Content === "") {
      document.getElementById("errorContent").innerHTML = "Content is required";
      document.getElementById("errorContent").style.color = "red";
      isValidContent = false;
    } else if (Content.length > 1000) {
      document.getElementById("errorContent").innerHTML = "max Content 1000";
      document.getElementById("errorContent").style.color = "red";
      isValidContent = false;
    } else {
      document.getElementById("errorContent").innerHTML = "";
      isValidContent = true;
    }

    if (categoryName.length === 0) {
      document.getElementById("errorcategory").innerHTML =
        "Please select at least one category";
      document.getElementById("errorcategory").style.color = "red";
      isValidCheckboxes = false;
    } else {
      document.getElementById("errorcategory").innerHTML = "";
      isValidCheckboxes = true;
    }

    if (isValidTitle && isValidContent && isValidCheckboxes) {

      let formData = new FormData();
      formData.append("Title", Title);
      formData.append("Content", Content);
      formData.append("options", categoryName);
      formData.append("Images", file); // Append the file
      try {
        const res = await fetch("http://localhost:8001/add-post", {
          method: "POST",
          // headers:,
          body: formData
        });
        if (res.ok) {
          const post = await res.json();
          const div = CreatPost(post);
          postsContainer.prepend(div);
          form.style.display = "none";
          layout.style.display = "none";
          document.body.style.overflow = "scroll";
        }
      } catch (error) {
        alert("Error: " + error.message);
      }
    }
  });
}

async function PostCategory() {
  let category = document.getElementById("category");

  try {
    const res = await fetch("/api/category/list");

    const categories = document.querySelector("#categories");
    const data = await res.json();

    data.forEach((catg) => {
      let li = createEle("li");
      let a = createEle("a");
      a.className = "category";

      a.href = `/category/${catg.Name}`;
      a.innerText = catg.Name;
      if (category) {
        category.innerHTML += `
        <label class="catLabel" for="${catg.Name}">
          <input type="checkbox" name="options" id="${catg.Name}" value="${catg.Name}" data-name="${catg.Name}"> <splan>${catg.Name}</span>
        </label>
  
        `;
      }
      li.append(a);
      categories.append(li);
    });
  } catch {
    console.log("erroure");
  }
}

async function GetComments(idPost, str) {
  str.innerText = "";
  str.style.display = "block";
  try {
    const response = await fetch(`/api/${idPost}/comments`);

    if (response.ok) {
      const data = await response.json();

      if (data == null) {
        str.innerHTML = "there is no comments";
      } else {
        const comments = createEle("div");
        comments.className = "commentsDiv";

        data.forEach((e) => {
          const commentC = createEle("div");
          commentC.className = "commentC";
          commentC.dataset.commentId = e.Id;

          const commentHe = createEle("div");
          commentHe.className = "commentHe";

          const commentH = createEle("h3");
          commentH.className = "commentH";
          commentH.innerText = e.Username;

          const commentTime = createEle("p");
          commentTime.className = "commentTime";
          commentTime.innerText = formDate(e.Date);

          const commentP = createEle("p");
          commentP.className = "commentp";
          commentP.innerText = e.Comment;

          const like_dislike_container = createEle("div");
          like_dislike_container.className = "like-dislike-container-comment";

          /* creating of the like button */
          //
          const [likebnt, dislikebnt, likeNmb, dislikeNmb] = HandulLike(
            e.Reactions.Action,
            e.Reactions.Likes,
            e.Reactions.Dislikes,
            "comment",
            e.Id
          );

          //
          // const likebnt = createEle("button");
          // likebnt.className = "like-btn-comment";
          likebnt.classList.add("like-btn-comment");

          /* create an img element to contain like icon */
          const likeIcon = createEle("img", "likeicon-comment");
          likeIcon.src = "/static/images/like.png";

          likebnt.appendChild(likeIcon);

          // const likeNmb = createEle("span");
          likeNmb.className = "likeNbr";
          likeNmb.innerText = e.Reactions.Likes;
          // likebnt.appendChild(likeNmb);

          /* creationg of the dislike button */
          // const dislikebnt = createEle("button");
          // dislikebnt.className = ;
          dislikebnt.classList.add("dislike-btn-comment");
          /* creating an img tag to containg dislike icon */
          const dislikeIcon = createEle("img");
          dislikeIcon.className = "dislikeicon-comment";
          dislikeIcon.src = "/static/images/dislike.png";

          dislikebnt.appendChild(dislikeIcon);

          // const dislikeNmb = createEle("span");
          dislikeNmb.className = "dislikeNbr";
          dislikeNmb.innerText = e.Reactions.Dislikes;
          // dislikebnt.appendChild(dislikeNmb);

          /* appending like and dislike buttons to like container */
          like_dislike_container.append(
            likebnt,
            likeNmb,
            dislikebnt,
            dislikeNmb
          );

          commentHe.append(commentH, commentTime);
          commentC.append(commentHe, commentP, like_dislike_container);
          comments.appendChild(commentC);
        });
        str.appendChild(comments);
      }
    } else {
      console.error("Request failed with status:", response.status);
      // document.getElementById("responseMessage").innerText = "Error fetching comments.";
    }
  } catch (error) {
    console.error(error);
  }
}
