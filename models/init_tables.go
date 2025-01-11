package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitTables(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username varchar(255) NOT NULL UNIQUE,
			email varchar(255) NOT NULL UNIQUE,
			password varchar(255) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS sessions(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL UNIQUE,
			token varchar(255) NOT NULL UNIQUE,
			date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id)
		);

		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			title varchar(250) NOT NULL,
			content TEXT NOT NULL,
			image varchar(250),
			categories TEXT NOT NULL,
			date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id)
		);

		
		CREATE TABLE IF NOT EXISTS categories (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name varchar(255) NOT NULL UNIQUE
		);

		CREATE TABLE IF NOT EXISTS comments (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			post_id INTEGER NOT NULL,
			comment TEXT NOT NULL,
			date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (post_id) REFERENCES posts(id)
		);
		
		CREATE TABLE IF NOT EXISTS reactions (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			post_id INTEGER,
			comment_id INTEGER,
			type varchar(50) NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (post_id) REFERENCES posts(id),
			FOREIGN KEY (comment_id) REFERENCES comments(id),
			UNIQUE(user_id,post_id),
			UNIQUE(user_id,comment_id),
			CHECK (
				(post_id IS NULL AND comment_id IS NOT NULL) OR
				(post_id IS NOT NULL AND comment_id IS NULL)
			)
		);

		CREATE TABLE IF NOT EXISTS posts_categories (
			post_id INTEGER NOT NULL,
			category_id INTEGER NOT NULL,
			PRIMARY KEY (post_id,category_id),
			FOREIGN KEY (post_id) REFERENCES posts(id),
			FOREIGN KEY (category_id) REFERENCES categories(id)
		);

		INSERT INTO categories(name) VALUES ("Dev"),("Tech"),("News"), ("Gaming"), ("Music") ON CONFLICT (name) DO NOTHING;
	`
	_, err := db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
