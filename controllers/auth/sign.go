package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"forum/api"
	"forum/utils"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SingIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	user := utils.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if !IsValidEmail(user.Email) || user.Password == "" || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Inavlid input for login"})
		return
	}
	id, err := Select(user.Email, user.Password)
	if err != nil {
		if id == -2 {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "there are error in server try anthor time"})
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "user or password not correct"})
		}

		return
	}
	uid, err := uuid.NewV4()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "there are error in server try anthor time"})
		return
	}
	err = CraeteSession(id, uid.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "there are error in server try anthor time"})
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    uid.String(),
		MaxAge:   int(time.Hour) * 24,
		HttpOnly: true,
		Path:     "/",
	})
}

func Select(userIfo, passwd string) (int, error) {
	query := `SELECT id , password FROM users
		WHERE email = ?
	`
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return -2, fmt.Errorf("problem in server try anther time")
	}
	defer stmt.Close()
	var hashpasswd string
	var id int
	err = stmt.QueryRow(userIfo).Scan(&id, &hashpasswd)
	if err == sql.ErrNoRows {
		return -1, fmt.Errorf("user or password not correct")
	} else if err != nil {
		return -1, fmt.Errorf("check your input")
	}
	if bcrypt.CompareHashAndPassword([]byte(hashpasswd), []byte(passwd)) != nil {
		return -1, fmt.Errorf("user or password not correct")
	}
	return id, nil
}

func CraeteSession(userid int, session string) error {
	query := `INSERT INTO sessions(user_id , token)
		VALUES(?,?)
		ON CONFLICT DO UPDATE SET token = EXCLUDED.token , date = CURRENT_TIMESTAMP
	`
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userid, session)

	return err
}

func Insert(user utils.User) (int, int, error) {
	query := `INSERT INTO users (username , email , password) 
		VALUES (?, ? , ?)
	`
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		return http.StatusInternalServerError, -1, fmt.Errorf("sorry but there are error in server try anther time")
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.Username, user.Email, user.Password)
	if err == nil {
		latdID, _ := res.LastInsertId()
		return http.StatusOK, int(latdID), nil
	}
	if strings.Contains(err.Error(), "username") {
		return http.StatusFound, -1, fmt.Errorf("user name already used try anther user name")
	} else if strings.Contains(err.Error(), "email") {
		return http.StatusFound, -1, fmt.Errorf("email already used try anther email")
	}
	return http.StatusInternalServerError, -1, fmt.Errorf("sorry but there are error in server try anther time")
}

func SingUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := utils.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if !IsValidUsername(user.Username) || !IsValidEmail(user.Email) || user.Password != user.ConfPass || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Inavlid input for Register"})
		return
	}
	if !isValidPassword(user.Password) {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "password is weak"})
		return
	}
	user.Password, err = HasPassowd(user.Password)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "there are error in server try anthor time"})

		return
	}
	statuscode, userId, err := Insert(user)
	if err != nil {
		w.WriteHeader(statuscode)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	uid, err := uuid.NewV4()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "there are error in server try anthor time"})
	}
	err = CraeteSession(userId, uid.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "there are error in server try anthor time"})
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    uid.String(),
		MaxAge:   int(time.Hour) * 24,
		HttpOnly: true,
		Path:     "/",
	})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("token")
	DeletSeoin(cookie.Value)
	http.SetCookie(w, &http.Cookie{
		Value:  "",
		Name:   "token",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DeletSeoin(token string) error {
	query := `
		DELETE FROM sessions WHERE token = ?
	`
	_, err := utils.DB.Exec(query, token)
	return err
}

func Islogin(r *http.Request) bool {
	userName := 0
	cookie, err := r.Cookie("token")
	if err != nil {
		return false
	}
	userName = api.TakeuserId(cookie.Value)
	return userName > 0
}
