package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"forum/utils"
)

func PostReaction(w http.ResponseWriter, r *http.Request) {
	var reactPost struct {
		user_id int    // get from token
		post_id int    // get from url
		action  string // get from form
	}
	var err error
	reactPost.post_id, err = strconv.Atoi(r.PathValue("PostId"))
	if utils.HandleError(utils.Error{Err: err, Code: http.StatusNotFound}, w) || CheckPost(reactPost.post_id) {
		return
	}

	/* ----------------------------------- handle action ----------------------------------- */
	reactPost.action = r.FormValue("action")
	/* ----------------------------------- Handle User Id ----------------------------------- */
	token, token_err := r.Cookie("token")
	if utils.HandleError(utils.Error{Err: token_err, Code: http.StatusUnauthorized}, w) {
		return
	}
	reactPost.user_id = TakeuserId(token.Value)
	action := CheckLIke(reactPost.post_id, reactPost.user_id, "post_id")
	if reactPost.action == "like" {
		if action == "dislike" {
			UpdateLike(reactPost.post_id, reactPost.user_id, "post_id", "like")
		} else if action == "like" {
			DeletLike(reactPost.post_id, reactPost.user_id, "post_id")
		} else {
			InsertLike(reactPost.post_id, reactPost.user_id, "post_id", "like")
		}
	} else if reactPost.action == "dislike" {
		if action == "like" {
			UpdateLike(reactPost.post_id, reactPost.user_id, "post_id", "dislike")
		} else if action == "dislike" {
			DeletLike(reactPost.post_id, reactPost.user_id, "post_id")
		} else {
			InsertLike(reactPost.post_id, reactPost.user_id, "post_id", "dislike")
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid format"})
		return
	}
}

func UpdateLike(id, userid int, which, typ string) {
	query := fmt.Sprintf(`
		UPDATE  reactions
		SET type = ?
		WHERE %s = ?
		AND user_id = ?
	`, which)
	utils.DB.Exec(query, typ, id, userid)
}

func InsertLike(id, userid int, which, typ string) {
	query := fmt.Sprintf(`
		INSERT INTO reactions(%s ,user_id ,type)
		VALUES( ? , ? , ? )
	`, which)
	utils.DB.Exec(query, id, userid, typ)
}

func DeletLike(id, userid int, which string) {
	query := fmt.Sprintf(`
		DELETE FROM reactions WHERE %s = ? AND user_id = ?
	`, which)
	utils.DB.Exec(query, id, userid)
}

func CheckLIke(id, userId int, column string) string {
	var like string
	query := fmt.Sprintf(`
		SELECT type FROM reactions WHERE %s = ? AND user_id = ?
	`, column)
	utils.DB.QueryRow(query, id, userId).Scan(&like)
	return like
}
