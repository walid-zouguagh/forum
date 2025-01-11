package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"forum/utils"
)

func Comments(w http.ResponseWriter, r *http.Request) {
	/* -------------------------- Get all comments for 1 post with PostId -------------------------- */
	w.Header().Set("Content-type", "application/json")
	var comments []utils.CommentType
	/* -------------------------- Check post_id -------------------------- */
	post_id, err := strconv.Atoi(r.PathValue("PostId"))
	if utils.HandleError(utils.Error{Err: err, Code: http.StatusNotFound}, w) || CheckPost(post_id) {
		return
	}

	userId := 0
	cookie, err := r.Cookie("token")
	if err == nil {
		userId = TakeuserId(cookie.Value)
	}
	/* -------------------------- Prepare query -------------------------- */
	query := `SELECT id,user_id,comment,date FROM comments WHERE post_id = ? ORDER BY id DESC`
	stmt, stmt_err := utils.DB.Prepare(query)
	if utils.HandleError(utils.Error{Err: stmt_err, Code: http.StatusInternalServerError}, w) {
		return
	}
	/* -------------------------- Get comments from DB -------------------------- */
	rows, rows_err := stmt.Query(post_id)
	if utils.HandleError(utils.Error{Err: rows_err, Code: http.StatusInternalServerError}, w) {
		return
	}
	for rows.Next() {
		var comment utils.CommentType
		var user_id int
		err := rows.Scan(&comment.Id, &user_id, &comment.Comment, &comment.Date)
		if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		comment.Username, err = GetUsername(user_id)
		if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
			return
		}
		comment.Reactions = GetReaction(userId, comment.Id, "comment_id")
		comments = append(comments, comment)
	}

	/* -------------------------- Set result in json response -------------------------- */
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}

func GetUsername(id int) (string, error) {
	/* ----------------------------- this function get username ----------------------------- */
	var username string
	query := `SELECT (username) FROM users WHERE id= ?`
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		return "", err
	}
	err = stmt.QueryRow(id).Scan(&username)
	return username, err
}
