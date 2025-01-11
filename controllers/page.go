package controllers

import (
	"net/http"

	"forum/api"
	"forum/controllers/auth"
	"forum/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ErrorHandler(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed), "this Page doesn't support your Method", nil)
		return
	}
	cookie, err := r.Cookie("token")
	if err == nil {
		userId := api.TakeuserId(cookie.Value)
		if userId > 0 {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
	pages := []string{"views/pages/login.html"}
	utils.ExecuteTemplate(w, pages, nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err == nil {
		userId := api.TakeuserId(cookie.Value)
		if userId > 0 {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
	pages := []string{"views/pages/register.html"}
	utils.ExecuteTemplate(w, pages, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { // handle invalid path (error 404)
		utils.ErrorHandler(w, http.StatusNotFound, "Page not Found", "The page you are looking for is not available!", nil)
		return
	}

	pages := []string{
		"views/pages/home.html",
	}
	utils.ExecuteTemplate(w, pages, auth.Islogin(r))
}

func LikedPostsPage(w http.ResponseWriter, r *http.Request) {
	pages := []string{
		"views/pages/likedPost.html",
	}
	utils.ExecuteTemplate(w, pages, true)
}
