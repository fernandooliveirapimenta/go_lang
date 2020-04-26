package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Post struct {
	Id int
	Title string
	Body string
}

func main() {
	fmt.Println("Iniciando server..")

	// No intelij deve conter workdir sn_web
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {

			post := Post{
				Id: 1,
				Title: "Unamed Post",
				Body: "No content",
			}

			if title := r.FormValue("title"); title != "" {
				post.Title = title
			}

			t := template.Must(template.ParseFiles(
	"templates/index.html"))
			err := t.ExecuteTemplate(w, "index.html", post)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

		//fmt.Fprintf(w, "Hello world")

	})
	fmt.Println(http.ListenAndServe(":8080", nil))
	fmt.Println("Server iniciado")
}
