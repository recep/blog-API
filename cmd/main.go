package main

import (
	"net/http"

	. "blog-API/routes"
	. "blog-API/utils"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db := DbConn()
	db.Exec("create table if not exists blog (ID integer,CreatedAt integer , UpdatedAt integer ,Title text,Body text)")

	router := mux.NewRouter()
	ReturnRoutes(router)
	http.ListenAndServe(":8081", router)
}
