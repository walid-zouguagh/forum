/* ------------------ dark mode start ------------------ */
let mode = localStorage.getItem("mode")
if (!mode) {
    localStorage.setItem("mode", "light")
    mode = localStorage.getItem("mode")
}
const mode_btn = document.getElementById("change_mode")
//  document.addEventListener("DOMContentLoaded", () => {})
mode_btn.addEventListener("click", () => {
    mode === "light" ? mode = "dark" : mode = "light";
    document.body.className = mode
    localStorage.setItem("mode", mode)
})

document.body.className = localStorage.getItem("mode")
/* ------------------ dark mode end ------------------ */