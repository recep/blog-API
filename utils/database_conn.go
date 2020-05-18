package utils

import (
	"database/sql"

	. "blog-API/helpers"
)

func DbConn() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "./database/blog.db")
	CheckErr(err)
	return
}
