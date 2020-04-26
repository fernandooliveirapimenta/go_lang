package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

type PostDomain struct {
	Id int
	Title string
	Body string
}

var db, erroCon =
	sql.Open("mysql",
		"root:MySql2019!@/go_course?charset=utf8")
func main() {
	checkErr(erroCon)

	fmt.Println("Iniciando server..")

	// insert
	//stmt, err :=
	// 	db.Prepare("Insert into posts(title, body) value (?,?)")
	//checkErr(err)
	//
	//_, erro:= stmt.Exec("My first post", "My first content")
	//checkErr(erro)
	//db.Close()



	rows, erroQ := db.Query("SELECT * FROM posts")
	checkErr(erroQ)

	items := []PostDomain{}
	for rows.Next() {
		post := PostDomain{}
		erros := rows.Scan(&post.Id, &post.Title, &post.Body)
		checkErr(erros)
		items = append(items, post)
	}
	er := db.Close()
	checkErr(er)

	fmt.Println(items)

	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {

			t := template.Must(template.ParseFiles(
				"templates/index.html"))
			err := t.ExecuteTemplate(w, "index.html", items)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

		})
	fmt.Println(http.ListenAndServe(":8080", nil))
	fmt.Println("Server iniciado")
}

	func checkErr(err error){
		if err != nil {
			panic(err)
		}
	}
