package main

import "net/http"

func (app *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.getHome)
	mux.HandleFunc("GET /posts/create", app.createPost)
	mux.HandleFunc("POST /posts/create", app.storePost)

	mux.HandleFunc("GET /register", app.getRegister)
	mux.HandleFunc("POST /register", app.storeRegister)
	mux.HandleFunc("GET /login", app.getLogin)
	mux.HandleFunc("POST /login", app.storeLogin)

	return app.authenticate(mux)

}

