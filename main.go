package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//Post type
type Post struct {
	ID    int
	Title string
}

func main() {
	db, err := sql.Open("mysql", "<user>:<password>@tcp(127.0.0.1:3306)/<database-name>")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	posts, err := db.Query("SELECT id, title FROM posts")
	if err != nil {
		panic(err.Error())
	}
	for posts.Next() {
		var post Post
		err := posts.Scan(&post.ID, &post.Title)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(post)
	}
}
