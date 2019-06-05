package main

import (
  "html/template"
  "net/http"
)

func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
  //threads, err := data.Threads()
  //if err == nil {
  //
  //}

  _, err := session(w, r)
  publicTmplFiles := []string{"templates/layout.html","templates/public.navbar.html", "templates/index.html"}
  privateTmplFiles := []string{"templates/layout.html", "templates/private.navbar.html", "templates/index.html"}

  var templates *template.Template
  if err != nil {
    templates = template.Must(template.ParseFiles(privateTmplFiles...))
  } else {
    templates = template.Must(template.ParseFiles(publicTmplFiles...))
  }

  templates.ExecuteTemplate(w, "layout", nil)
}
