import { createEle, formDate } from "./posts.js";

export function HandulLike(type, numlike, numdislike, path, id) {
  const likebtn = document.createElement("button");
  const dislikebtn = document.createElement("button");
  if (type === "like") {
    likebtn.classList.add("liked");
  } else if (type === "dislike") {
    dislikebtn.classList.add("disliked");
  }
  const likespan = document.createElement("span");
  const dislikespan = document.createElement("span");
  likespan.innerHTML = numlike;
  dislikespan.innerHTML = numdislike;
  likebtn.addEventListener("click", () => {
    if (likebtn.classList.length == 2) {
      likebtn.classList.remove("liked");
      numlike--;
    } else {
      if (dislikebtn.classList.length == 2) {
        dislikebtn.classList.remove("disliked");
        numdislike--;
      }
      numlike++;
      likebtn.classList.add("liked");
    }
    likespan.innerHTML = numlike;
    dislikespan.innerHTML = numdislike;
    fetch(`/api/${path}/reaction/${id}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: `action=like`,
    }).then((res) => {
      if (res.redirected) {
        window.location.href = "login";
      }
    });
  });
  dislikebtn.addEventListener("click", () => {
    if (dislikebtn.classList.length == 2) {
      dislikebtn.classList.remove("disliked");
      numdislike--;
    } else {
      if (likebtn.classList.length == 2) {
        likebtn.classList.remove("liked");
        numlike--;
      }
      numdislike++;
      dislikebtn.classList.add("disliked");
    }
    likespan.innerHTML = numlike;
    dislikespan.innerHTML = numdislike;
    fetch(`/api/${path}/reaction/${id}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: `action=dislike`,
    }).then((res) => {
      if (res.redirected) {
        window.location.href = "login";
      }
    });
  });
  return [likebtn, dislikebtn, likespan, dislikespan];
}

export function CreatPost(post) {
  const postElement = document.createElement("div");
  postElement.className = "post-container";
  postElement.dataset.postId = post.Id;

  /* h2 will contain the image and name of the persen who posted */
  const posterName = createEle("h2");
  posterName.className = "poster";
  const posterImg = createEle("img");
  posterImg.src =
    "/static/images/466006304_871124095226532_8631138819273739648_n.jpg";
  const nameContainer = createEle("span");
  nameContainer.className = "usrname";
  nameContainer.innerText = post.UserName;
  const createdat = createEle("span");

  createdat.innerText = formDate(post.Date);
  createdat.className = "creationdate";
  posterName.appendChild(posterImg);
  posterName.appendChild(nameContainer);
  postElement.appendChild(posterName);
  postElement.appendChild(createdat);

  /* creating a div that will contain all the elements bellow */
  const pc = createEle("div");
  pc.className = "pc";

  /* creating an h3 element to contain the post title */
  const title = createEle("h3");
  title.className = "title";
  title.innerText = post.Title;

  /* creating a p element that will contain the content of the post */
  const content = createEle("p");
  content.className = "content";
  content.innerText = post.Content;
  pc.append(title, content);

  const categories_container = createEle("div");
  categories_container.className = "categories";

  for (let cate of post.Categories) {
    const span = createEle("a");
    span.href = `/category/${cate}`;
    span.className = "category";
    span.innerText = cate;
    categories_container.appendChild(span);
  }
  pc.appendChild(categories_container);
  /* creating like and dislike button */
  const like_dislike_container = createEle("div");
  like_dislike_container.className = "like-dislike-container";
  //

  const [likebnt, dislikebnt, likeNbr, dislikeNbr] = HandulLike(
    post.Reactions.Action,
    post.Reactions.Likes,
    post.Reactions.Dislikes,
    "posts",
    post.Id
  );
  ///
  /* creating of the like button */
  // const likebnt = createEle("button");
  likebnt.classList.add("like-btn");

  /* create an img element to contain like icon */
  const likeIcon = createEle("img");
  likeIcon.className = "likeIcon";
  likeIcon.src = "/static/images/like.png";

  // const likeNbr = createEle("span");
  likeNbr.className = "likeNbr";
  // likeNbr.innerText = post.Reactions.Likes;

  likebnt.appendChild(likeIcon);
  // likebnt.appendChild(likeNbm);

  /* creationg of the dislike button */
  // const dislikebnt = createEle("button");

  dislikebnt.classList.add("dislike-btn");
  /* creating an img tag to containg dislike icon */
  const dislikeIcon = createEle("img");
  dislikeIcon.className = "dislikeIcon";
  dislikeIcon.src = "/static/images/dislike.png";

  // const dislikeNbr = createEle("span");
  dislikeNbr.className = "dislikeNbr";
  // dislikeNbr.innerText = post.Reactions.Dislikes;

  dislikebnt.appendChild(dislikeIcon);
  // dislikebnt.appendChild(dislikeNbr);

  /* appending like and dislike buttons to like container */
  like_dislike_container.append(likebnt, likeNbr, dislikebnt, dislikeNbr);

  /* appending like container to the post contaner */
  pc.appendChild(like_dislike_container);

  ////add div comments
  const divcomments = createEle("div");
  divcomments.className = `divcomments${post.Id} divcomments`;

  pc.appendChild(divcomments);

  /* adding a button to see comments */
  const seecomments = createEle("button");
  seecomments.className = "see_comments";
  seecomments.innerText = "see comments";
  pc.appendChild(seecomments);
  /* creating the form that sends comments */
  const comment_form = createEle("form");
  comment_form.method = "POST";
  comment_form.className = "comment_form";

  const title_impt = createEle("input");
  title_impt.className = "comment";
  title_impt.name = "comment";
  title_impt.type = "text";
  title_impt.placeholder = "Add your comment";
  title_impt.required = true;

  const submit_comment = createEle("button");
  submit_comment.className = "send_comment";
  submit_comment.type = "submit";

  const send_icon = createEle("img");
  send_icon.className = "sendimg";
  send_icon.src = "/static/images/send-message.png";
  submit_comment.appendChild(send_icon);
  comment_form.appendChild(title_impt);
  comment_form.appendChild(submit_comment);

  pc.appendChild(comment_form);
  postElement.appendChild(pc);
  return postElement;
}
