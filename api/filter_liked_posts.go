package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"forum/utils"
)

func LikedPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	offset := r.URL.Query().Get("offset")

	limitInt := 20
	offsetInt := 0
	if o, err := strconv.Atoi(offset); err == nil {
		offsetInt = o
	}
	cookie, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	userid := TakeuserId(cookie.Value)
	if userid < 1 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	query := `
		SELECT 
			users.username, 
			posts.id, 
			posts.title, 
			posts.content, 
			posts.date, 
			posts.categories
		FROM 
			users
		INNER JOIN 
			posts 
		ON 
			posts.user_id = users.id
		INNER JOIN 
			reactions 
		ON 
			reactions.post_id = posts.id
		WHERE 
			reactions.user_id = ? AND reactions.type = ?
		ORDER BY 
			posts.id DESC
		LIMIT ? OFFSET ?;
	`
	rows, err := utils.DB.Query(query, userid, "like", limitInt, offsetInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	defer rows.Close()
	posts := []utils.PostsResult{}
	for rows.Next() {
		post := utils.PostsResult{}
		categories := ""

		if err := rows.Scan(&post.UserName, &post.Id, &post.Title, &post.Content, &post.Date, &categories); err != nil {
			continue
		}
		post.Categories = strings.Split(categories, ",")
		post.Reactions = GetReaction(userid, post.Id, "post_id")
		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)
}

func TakeuserId(secion string) int {
	query := `
		SELECT user_id FROM sessions WHERE token = ?
	`
	id := 0
	err := utils.DB.QueryRow(query, secion).Scan(&id)
	if err != nil {
		return -1
	}
	return id
}

func GetReaction(userid, id int, colom string) utils.Reactions {
	reaction := utils.Reactions{}
	query := `
		SELECT count(*) FROM reactions WHERE %s = ? AND type = ?
	`
	query = fmt.Sprintf(query, colom)
	utils.DB.QueryRow(query, id, "like").Scan(&reaction.Likes)
	utils.DB.QueryRow(query, id, "dislike").Scan(&reaction.Dislikes)
	if userid < 1 {
		return reaction
	}
	query = `
		SELECT type FROM reactions WHERE %s = ? AND user_id = ?
	`
	query = fmt.Sprintf(query, colom)
	utils.DB.QueryRow(query, id, userid).Scan(&reaction.Action)
	return reaction
}
