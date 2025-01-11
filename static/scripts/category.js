import { addEventOnPosts, handleScroll } from "./posts.js";
let path = "/api/filter" + window.location.pathname;

addEventOnPosts(path);

window.addEventListener("scrollend", () => {
  setInterval(handleScroll(path));
});
