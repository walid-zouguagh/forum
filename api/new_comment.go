package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"forum/utils"
)

func NewComment(w http.ResponseWriter, r *http.Request) {
	var comment utils.Comment
	/* ----------------------------- token validation ----------------------------- */
	token, tokenErr := r.Cookie("token")
	if tokenErr != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	/* ----------------------------- handle post_id ----------------------------- */
	postId := r.PathValue("PostId")
	var postIdErr error
	comment.PostId, postIdErr = strconv.Atoi(postId)
	if postIdErr != nil || CheckPost(comment.PostId) {
		err := errors.New("post_id not valid")
		if utils.HandleError(utils.Error{Err: err, Code: http.StatusBadRequest}, w) {
			return
		}
	}
	/* ----------------------------- Prepare query for get user_id ----------------------------- */
	// getUserIdQuery := `SELECT user_id FROM sessions WHERE token=?;`
	// stmt, stmt_err := utils.DB.Prepare(getUserIdQuery)
	// if utils.HandleError(utils.Error{Err: stmt_err, Code: http.StatusInternalServerError}, w) {
	// 	return
	// }

	// /* ----------------------------- Get user_id from DB ----------------------------- */
	// queryErr := stmt.QueryRow(token.Value).Scan(&comment.UserId)
	// if utils.HandleError(utils.Error{Err: queryErr, Code: http.StatusInternalServerError}, w) {
	// 	return
	// }
	comment.UserId = TakeuserId(token.Value)
	/* ----------------------------- Handle comment data ----------------------------- */
	comment.Comment = strings.TrimSpace(r.FormValue("comment"))
	if len(comment.Comment) < 1 || len(comment.Comment) > 500 {
		err := errors.New("bad request empty comment")
		if utils.HandleError(utils.Error{Err: err, Code: http.StatusBadRequest}, w) {
			return
		}
	}

	/* ----------------------------- Prepare create comment query ----------------------------- */
	query := `INSERT INTO comments(user_id,post_id,comment) VALUES (?,?,?);`
	stmt, stmt_err := utils.DB.Prepare(query)
	if utils.HandleError(utils.Error{Err: stmt_err, Code: http.StatusInternalServerError}, w) {
		return
	}
	_, err := stmt.Exec(comment.UserId, comment.PostId, comment.Comment)
	if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
		return
	}

	/* ----------------------------- this response if comment is created ----------------------------- */
	response := map[string]int{http.StatusText(http.StatusCreated): http.StatusCreated}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func CheckPost(postId int) bool {
	/* ----------------------------- Check post ----------------------------- */
	var exist int
	query := `SELECT EXISTS(SELECT 1 FROM posts WHERE id=?);`
	err := utils.DB.QueryRow(query, postId).Scan(&exist)
	if exist == 0 {
		err = errors.New("post not exist")
	}
	return err != nil
}

func CheckCommat(id int) error {
	query := `
		SELECT EXISTS (
    SELECT 1 FROM comments WHERE id = ? )
	`
	exist := false
	err := utils.DB.QueryRow(query, id).Scan(&exist)
	return err
}
