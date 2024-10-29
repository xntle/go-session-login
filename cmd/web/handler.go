package main

import (
	"html/template"

	"net/http"
)

func (app *app) getHome(w http.ResponseWriter, r *http.Request){
	posts, err := app.posts.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t, err := template.ParseFiles("./asset/templates/home.page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, map[string]any{"Posts": posts})
}
func (app *app) getRegister(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./asset/templates/register.page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

func (app *app) storeRegister(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name := r.PostForm.Get("name")
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")

	err = app.users.Insert(name, email, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}

func (app *app) getLogin(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./asset/templates/login.page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

func (app *app) storeLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := app.users.Authenticate(
		r.PostForm.Get("email"), 
		r.PostForm.Get("password"),
	)



	if err != nil {
		http.Error(w, err.Error(), 500)
	
		return
	}

	session, _ := app.session.Get(r, "anleapp")
	// Set some session values.
	session.Values["userID"] = id
	// Save it before we write to the response/return from the handler.
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	
	http.Redirect(w, r, "/", http.StatusFound)
}

func (app *app) createPost(w http.ResponseWriter, r *http.Request){

	t, err := template.ParseFiles("./asset/templates/posts.create.page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

func (app *app) storePost(w http.ResponseWriter, r *http.Request){

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = app.posts.Insert(
		r.PostFormValue("title"),
		r.PostFormValue("content"))


	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}