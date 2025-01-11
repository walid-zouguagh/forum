package controllers

import (
	"net/http"

	"forum/api"
	"forum/controllers/auth"
	"forum/utils"
)

func CreatedPosts(w http.ResponseWriter, r *http.Request) {
	pages := []string{"views/pages/profile.html"}
	if r.Method != http.MethodGet {
		utils.ErrorHandler(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed), "this Page doesn't support your Method", nil)
		return
	}
	utils.ExecuteTemplate(w, pages, true)
}

func Categories(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ErrorHandler(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed), "this Page doesn't support your Method", nil)
		return
	}
	categoriesName := r.PathValue("nameCategory")
	category := api.TakeCategories(categoriesName)
	if category < 1 {
		utils.ErrorHandler(w, http.StatusNotFound, "Page not Found", "The page you are looking for is not available!", nil)
		return
	}
	pages := []string{"views/pages/categories.html"}
	utils.ExecuteTemplate(w, pages, auth.Islogin(r))
}
