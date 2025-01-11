package router

import (
	"net/http"

	"forum/api"
	m "forum/middleware"
)

func APIRouter() *http.ServeMux {
	/* --------------------------- api handlers --------------------------- */
	apiMux := http.NewServeMux()
	apiMux.HandleFunc(`/api/{PostId}/comments`, m.CheckMethod(api.Comments, "GET"))                  // comments list
	apiMux.HandleFunc("/api/category/list", m.CheckMethod(api.CategoryList, "GET"))                  // get all categories
	apiMux.HandleFunc("/api/posts", m.CheckMethod(api.FetchPosts, "GET"))                            // get all posts
	apiMux.HandleFunc("/api/filter/category/{Category}", m.CheckMethod(api.FilterByCategory, "GET")) // filter by category

	apiMux.HandleFunc("/api/{PostId}/comment/new", m.CheckMethod(m.Authorization(api.NewComment), "POST")) // create comment
	apiMux.HandleFunc("/api/created/posts", m.CheckMethod(m.Authorization(api.CreatedPosts), "GET"))       // filter by created posts
	apiMux.HandleFunc("/api/liked/posts", m.CheckMethod(m.Authorization(api.LikedPosts), "GET"))           // not complated
	apiMux.HandleFunc("/api/comment/reaction/{CommentId}", m.Authorization(api.CommentReaction))           // (like or dislike) a comment
	apiMux.HandleFunc("/api/posts/reaction/{PostId}", m.Authorization(api.PostReaction))                   // create new post
	return apiMux
}
