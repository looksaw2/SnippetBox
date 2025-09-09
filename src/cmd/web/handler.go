package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	//检查URL
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	files := []string{
		"./src/ui/html/base.tmpl",
		"./src/ui/html/pages/home.tmpl",
		"./src/ui/html/partials/nav.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) //Internal Error -> 500
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err) //Internal Error -> 500
		return
	}
	w.Write([]byte("This is the home ........."))
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	//从Query中得到对应的ID
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	w.Write([]byte(fmt.Sprintf("Display a specific  snippet  with ID %d", id)))
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	//检查是否为POST方法
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed) // 405是Method不被允许
		return
	}
	w.Write([]byte("Create a new snippet"))
}
