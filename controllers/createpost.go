package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"forum/api"
	"forum/utils"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"status"`
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := Error{Message: "Not Allowed", Code: http.StatusMethodNotAllowed}
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(err)
		return
	}
	post := utils.PostsResult{}
	cookie, err := r.Cookie("token") // Name the Cookie
	if err != nil {
		err := Error{Message: "Unauthorized", Code: http.StatusUnauthorized}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	errt := r.ParseMultipartForm(10 << 20) // Limit of 10 MB
	if errt != nil {
		fmt.Println("ffff", errt)
	}

	err = r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	Categoriesid := []int{}
	post.Categories = strings.Split(r.FormValue("options"), ",")
	post.Content = r.FormValue("Content")
	post.Title = r.FormValue("Title")
	file, handler, errImage := r.FormFile("Images")
	if errImage != nil {
		fmt.Println("errImage", errImage)
	}

	post.Image = ""
	if file != nil {
		ext := strings.ToLower(filepath.Ext(handler.Filename))
		allowedExtensions := map[string]bool{
			".jpeg": true,
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".svg":  true,
		}
		if !allowedExtensions[ext] {
			err := Error{Message: "Not Allowed", Code: http.StatusMethodNotAllowed}
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(err)
			return
		}
		const maxSize = 20 * 1024 * 1024
		if handler.Size > maxSize {
			err := Error{Message: "Not Allowed", Code: http.StatusMethodNotAllowed}
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(err)
			return
		}
		timestamp := time.Now().Format("20060102_150405")
		fmt.Println(timestamp)
		uniqueFilename := timestamp + handler.Filename
		post.Image = filepath.Join("static/images", uniqueFilename)

		fmt.Println(post.Image)
		dst, err := os.Create(post.Image)
		if err != nil {
			err := Error{Message: "Not Allowed", Code: http.StatusMethodNotAllowed}
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(err)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			err := Error{Message: "Not Allowed", Code: http.StatusMethodNotAllowed}
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(err)
			return
		}
	}

	for _, categ := range post.Categories {
		categid := api.TakeCategories(categ)
		if categid < 1 {
			err := Error{Message: "Title or Content is more than expact", Code: http.StatusBadRequest}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
		Categoriesid = append(Categoriesid, categid)
	}

	var userId int

	err = utils.DB.QueryRow("SELECT user_id FROM sessions WHERE token = ?", cookie.Value).Scan(&userId)
	if err != nil {
		err := Error{Message: "Error", Code: 500}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
		return
	}

	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
	if post.Title == "" || post.Content == "" {
		err := Error{Message: "Title or Content is empty", Code: http.StatusBadRequest}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	if len(post.Title) > 100 || len(post.Content) > 1000 {
		err := Error{Message: "Title or Content is more than expact", Code: http.StatusBadRequest}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	// result, err := utils.DB.Exec("INSERT INTO posts(user_id, title, content, categories) VALUES(?, ?, ?, ?)", userId, post.Title, post.Content, strings.Join(post.Categories, ","))
	// if err != nil {
	// 	err := Error{Message: "can insert in base donne", Code: http.StatusUnauthorized}
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	json.NewEncoder(w).Encode(err)
	// 	return
	// }

	var result sql.Result
	var errInsert error
	if post.Image != "" {
		result, errInsert = utils.DB.Exec("INSERT INTO posts(user_id, title, content, categories,image) VALUES(?, ?, ?, ?,?)", userId, post.Title, post.Content, strings.Join(post.Categories, ","),post.Image)
	} else {
		result, errInsert = utils.DB.Exec("INSERT INTO posts(user_id, title, content, categories) VALUES(?, ?, ?, ?)", userId, post.Title, post.Content, strings.Join(post.Categories, ","))
	}

	if errInsert != nil {
		err := Error{Message: "can insert in base donne", Code: http.StatusUnauthorized}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		return
	}

	last_post_id, err := result.LastInsertId()
	if err != nil {
		// http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		err := Error{Message: "Error", Code: http.StatusInternalServerError}
		// http.Redirect(w, r, "/login", http.StatusSeeOther)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	post.Id = int(last_post_id)
	for _, categ := range Categoriesid {
		_, err = utils.DB.Exec("INSERT INTO posts_categories(post_id, category_id) VALUES(?, ?)", post.Id, categ) // GetLast id in table posts
		if err != nil {
			err := Error{Message: "Bad Request", Code: http.StatusInternalServerError}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
	}
	json.NewEncoder(w).Encode(post)
}
