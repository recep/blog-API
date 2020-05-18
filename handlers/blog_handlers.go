package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	. "blog-API/helpers"
	. "blog-API/model"
	. "blog-API/utils"

	"github.com/gorilla/mux"
)

func ReturnAllPost(w http.ResponseWriter, r *http.Request) {
	db := DbConn()
	rows, err := db.Query("select * from blog")
	CheckErr(err)

	var posts []Jblog
	var post Jblog

	for rows.Next() {
		err := rows.Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt, &post.Title, &post.Body)
		CheckErr(err)
		posts = append(posts, post)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(posts)
	CheckErr(err)
}

func ReturnSinglePost(w http.ResponseWriter, r *http.Request) {
	db := DbConn()
	id := mux.Vars(r)["id"]

	row, err := db.Query("SELECT * FROM blog WHERE id=?", id)
	CheckErr(err)

	var post Jblog
	for row.Next() {
		err := row.Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt, &post.Title, &post.Body)
		CheckErr(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(post)
	CheckErr(err)
}

func CreateNewPost(w http.ResponseWriter, r *http.Request) {
	db := DbConn()
	tx, err := db.Begin()
	CheckErr(err)

	stmt, err := tx.Prepare("INSERT INTO blog (ID,CreatedAt,UpdatedAt,Title,Body) VALUES (?,?,?,?,?)")
	CheckErr(err)

	resp, err := ioutil.ReadAll(r.Body)
	CheckErr(err)

	var post Blog
	err = json.Unmarshal(resp, &post)
	CheckErr(err)

	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	_, err = stmt.Exec(post.ID, post.CreatedAt, post.UpdatedAt, post.Title, post.Body)
	CheckErr(err)

	err = tx.Commit()
	CheckErr(err)

	w.WriteHeader(http.StatusCreated)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	db := DbConn()
	tx, err := db.Begin()
	CheckErr(err)

	id := mux.Vars(r)["id"]

	resp, err := ioutil.ReadAll(r.Body)
	CheckErr(err)

	var post Blog
	err = json.Unmarshal(resp, &post)
	CheckErr(err)

	post.UpdatedAt = time.Now()

	stmt, err := tx.Prepare("update blog set ID=?,UpdatedAt=?,Title=?,Body=? where id=?")
	CheckErr(err)

	_, err = stmt.Exec(post.ID, post.UpdatedAt, post.Title, post.Body, id)
	CheckErr(err)

	err = tx.Commit()
	CheckErr(err)

	w.WriteHeader(http.StatusOK)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	db := DbConn()
	tx, err := db.Begin()
	CheckErr(err)

	stmt, err := tx.Prepare("delete from blog where ID=?")
	CheckErr(err)

	id := mux.Vars(r)["id"]
	_, err = stmt.Exec(id)
	CheckErr(err)

	err = tx.Commit()
	CheckErr(err)

	w.WriteHeader(http.StatusOK)
}
