package utils

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ConfPass string `json:"password_config"`
}

type Category struct {
	Id   int
	Name string
}

type Comment struct {
	Id      int
	UserId  int
	PostId  int
	Comment string
	Date    string
}

type PostsResult struct {
	Id         int
	UserName   string
	Title      string
	Content    string
	Image      string
	Categories []string
	Date       string
	Reactions  Reactions
}
type CommentType struct {
	Id        int
	Username  string
	UserImage string
	Comment   string
	Date      string
	Reactions Reactions
}

type Reactions struct {
	Likes    int
	Dislikes int
	Action   string
}

type PostCategory struct {
	PostId     int
	CategoryId int
}

type Error struct {
	Err  error
	Code int
}

type ErrorData struct {
	Msg1       string
	Msg2       string
	StatusCode int
}
