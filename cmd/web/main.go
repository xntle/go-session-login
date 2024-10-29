package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xntle/go-session-login/internal/models/sqlite"
)

type app struct {
	posts *sqlite.PostModel	
	users *sqlite.UserModel	
	session *sessions.CookieStore
}
func main() {

	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	} 

	// Initialize a new cookie store for session management using a secret key to secure cookies.
	// The key should be kept secret to prevent tampering with session data.
	session := sessions.NewCookieStore([]byte("OFOeQti9yOYj1eQRFi9vhz6rru6WOP1Z"))

	// Set the HttpOnly option to true, which helps prevent client-side scripts from accessing the cookie.
	// This is a security measure to protect against cross-site scripting (XSS) attacks.
	session.Options.HttpOnly = true

	// Set SameSite cookie attribute to Lax mode, restricting cookies from being sent with cross-site
	// requests except for top-level navigations. This adds protection against CSRF attacks.
	session.Options.SameSite = http.SameSiteLaxMode
	
	app := app{
		posts: &sqlite.PostModel{
			DB: db,
		},	
		users: &sqlite.UserModel{
			DB: db,
		},	
		session: session,
	}

	srv := http.Server{
		Addr: ":8000",
		Handler: app.routes(),
	}

	log.Println("Starting server on :8000")
	srv.ListenAndServe()
}