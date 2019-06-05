package main

import (
  "github.com/manakuro/chitchat/data"
  "net/http"
)

func login(writer http.ResponseWriter, request *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(writer, nil)
}

func authenticate(writer http.ResponseWriter, request *http.Request) {
  err := request.ParseForm()
  user, err := data.UserByEmail(request.PostFormValue("email"))
  if err != nil {
    danger(err, "Cannot find user")
  }
  if user.Password == data.Encrypt(request.PostFormValue("password")) {
    session, err := user.CreateSession()
    if err != nil {
      danger(err, "Cannot create session")
    }
    cookie := http.Cookie{
      Name:     "_cookie",
      Value:    session.Uuid,
      HttpOnly: true,
    }
    http.SetCookie(writer, &cookie)
    http.Redirect(writer, request, "/", 302)
  } else {
    http.Redirect(writer, request, "/login", 302)
  }
}
