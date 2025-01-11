import { addEventOnPosts, handleScroll } from "./posts.js";
let path = "/api/posts";
addEventOnPosts(path);

window.addEventListener("scrollend", () => {
  setInterval(handleScroll(path));
});