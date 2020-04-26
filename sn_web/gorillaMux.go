package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type PostModel struct {
	Id int
	Title string
	Body string
}

var dbMsql, erroMsql =
	sql.Open("mysql",
		"root:MySql2019!@/go_course?charset=utf8")

func main() {
	checkErrr(erroMsql)

	fmt.Println("Iniciando server..")

	r := mux.NewRouter()

	// Servindo arquivos staticos
	r.PathPrefix("/static").Handler(
		http.StripPrefix("/static",
			http.FileServer(http.Dir("static/"))))

	r.HandleFunc("/", HomeHandler )
	r.HandleFunc("/{id}/post", ViewHandler)

	fmt.Println(http.ListenAndServe(":8080", r))
}

func GetPostById(id string) *PostModel {
	row := dbMsql.QueryRow("SELECT * from posts where id=?",id)
	post := new(PostModel)
	e := row.Scan(&post.Id, &post.Title, &post.Body)
	checkErrr(e)
	return post
}

func ListPosts() []PostModel{
	rows, err := dbMsql.Query("SELECT * from posts")
	checkErrr(err)

	items := []PostModel{}
	for rows.Next() {
		post := PostModel{}
		er := rows.Scan(&post.Id, &post.Title, &post.Body)
		checkErrr(er)
		items = append(items, post)
	}

	fmt.Println(items)
	return items
}

func ViewHandler(w http.ResponseWriter, r *http.Request){

	id := mux.Vars(r)["id"]
	t := template.Must(
		template.ParseFiles("templates/layout.html","templates/post.html"))
	if err := t.ExecuteTemplate(
		w, "layout.html", GetPostById(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"templates/layout.html", "templates/list.html"))
	err := t.ExecuteTemplate(w, "layout.html", ListPosts())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func checkErrr(err error){
	if err != nil {
		panic(err)
	}
}
