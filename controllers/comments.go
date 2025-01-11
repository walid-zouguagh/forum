package controllers

import (
	"forum/utils"
)

func CheckPost(postId int) bool {
	query := `SELECT id FROM posts WHERE id=?;`
	var id int
	err := utils.DB.QueryRow(query, postId).Scan(&id)
	return err != nil
}


