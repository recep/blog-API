package main

import (
	"log"
	"net/http"

	. "blog-API/helpers"
	. "blog-API/routes"
	. "blog-API/utils"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db := DbConn()
	_, err := db.Exec("create table if not exists blog (ID integer primary key,CreatedAt integer , UpdatedAt integer ,Title text,Body text)")
	CheckErr(err)

	router := mux.NewRouter()
	ReturnRoutes(router)
	log.Fatal(http.ListenAndServe(":8081", router))
}
