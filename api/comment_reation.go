package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"forum/utils"
)

func CommentReaction(w http.ResponseWriter, r *http.Request) {
	var reactInfo struct {
		user_id    int    // get from token
		comment_id int    // get from url
		action     string // get from form
	}
	w.Header().Set("Content-Type", "application/json")

	/* ------------------------------ handle user_id ------------------------------ */
	/* get cookie */
	cookie, cookie_err := r.Cookie("token")
	if utils.HandleError(utils.Error{Err: cookie_err, Code: http.StatusUnauthorized}, w) {
		return
	}

	/* select user_id from database */
	get_user := `SELECT user_id FROM sessions WHERE token= ?`
	stm, stm_err := utils.DB.Prepare(get_user)
	if utils.HandleError(utils.Error{Err: stm_err, Code: http.StatusInternalServerError}, w) {
		return
	}
	err := stm.QueryRow(cookie.Value).Scan(&reactInfo.user_id)
	if utils.HandleError(utils.Error{Err: err, Code: http.StatusInternalServerError}, w) {
		return
	}
	reactInfo.action = r.FormValue("action")
	/* ------------------------------ handle comment_id ------------------------------ */
	reactInfo.comment_id, err = strconv.Atoi(r.PathValue("CommentId"))
	if utils.HandleError(utils.Error{Err: err, Code: http.StatusNotFound}, w) || CheckCommat(reactInfo.comment_id) != nil {
		return
	}
	action := CheckLIke(reactInfo.comment_id, reactInfo.user_id, "comment_id")
	if reactInfo.action == "like" {
		if action == "dislike" {
			UpdateLike(reactInfo.comment_id, reactInfo.user_id, "comment_id", "like")
		} else if action == "like" {
			DeletLike(reactInfo.comment_id, reactInfo.user_id, "comment_id")
		} else {
			InsertLike(reactInfo.comment_id, reactInfo.user_id, "comment_id", "like")
		}
	} else if reactInfo.action == "dislike" {
		if action == "like" {
			UpdateLike(reactInfo.comment_id, reactInfo.user_id, "comment_id", "dislike")
		} else if action == "dislike" {
			DeletLike(reactInfo.comment_id, reactInfo.user_id, "comment_id")
		} else {
			InsertLike(reactInfo.comment_id, reactInfo.user_id, "comment_id", "dislike")
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid format"})
		return
	}
}
