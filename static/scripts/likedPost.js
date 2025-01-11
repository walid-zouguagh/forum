import { addEventOnPosts, handleScroll } from "./posts.js";
let path = "/api/liked/posts";

addEventOnPosts(path);

window.addEventListener("scrollend", () => {
  setInterval(handleScroll(path));
});
