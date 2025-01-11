package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"forum/utils"
)

func FilterByCategory(w http.ResponseWriter, r *http.Request) {
	/* ---------------------------- check if user is logged and get here id ---------------------------- */
	var logged_user int
	token, token_err := r.Cookie("token")
	if token_err == nil {
		logged_user = TakeuserId(token.Value)
	}
	Category := r.PathValue("Category")
	category_id := TakeCategories(Category)
	// fmt.Println(category_id)
	if category_id < 1 {
		json.NewEncoder(w).Encode(nil)
		return
	}

	offset := r.URL.Query().Get("offset")

	limitInt := 20
	offsetInt := 0

	if o, err := strconv.Atoi(offset); err == nil {
		offsetInt = o
	}
	var posts []utils.PostsResult
	query := `
		SELECT p.id,p.user_id,p.title,p.content,p.categories,p.date, u.username
		FROM posts p JOIN posts_categories pc
		ON p.id = pc.post_id AND pc.category_id = ? 
		JOIN users u ON u.id = p.user_id  ORDER BY p.id DESC LIMIT ? OFFSET ?;
	`
	stmt, stmt_err := utils.DB.Prepare(query)
	if utils.HandleError(utils.Error{Err: stmt_err, Code: http.StatusInternalServerError}, w) {
		return
	}
	rows, rows_err := stmt.Query(category_id, limitInt, offsetInt)
	if utils.HandleError(utils.Error{Err: rows_err, Code: http.StatusInternalServerError}, w) {
		return
	}
	for rows.Next() {
		var user_id int
		var categories string
		var p utils.PostsResult // this is a post i'm use p for short var
		err := rows.Scan(&p.Id, &user_id, &p.Title, &p.Content, &categories, &p.Date, &p.UserName)
		if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		p.Reactions = GetReaction(logged_user, p.Id, "post_id")
		p.Categories = strings.Split(categories, ",")
		posts = append(posts, p)
	}

	/* -------------------- Set result in json response -------------------------- */
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func TakeCategories(Category string) int {
	category_id := -1
	utils.DB.QueryRow("SELECT id FROM categories WHERE name = ?", Category).Scan(&category_id)
	return category_id
}
