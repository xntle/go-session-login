package main

import (
	"context"
	"net/http"
)

// Define a custom type for context keys to avoid collisions with other context keys in the application
type contextKey string

// Create a constant for the "user" context key to store user data in the request context
const contextKeyUser = contextKey("user")

// Middleware function to authenticate users before passing requests to the next handler
func (app *app) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the session from the request, ignoring any errors during session retrieval
		session, _ := app.session.Get(r, "anleapp")

		// Try to get the userID from the session. If it’s not available or cannot be cast to int, proceed without authentication
		userID, ok := session.Values["userID"].(int)
		if !ok {
			next.ServeHTTP(w, r) // Continue without modifying the context
			return
		}

		// Find the user by userID from the database or user store
		user, err := app.users.Find(userID)
		if err != nil {
			next.ServeHTTP(w, r) // If user retrieval fails, continue without authentication
			return
		}

		// Add the user information to the request context for use in later handlers
		ctx := context.WithValue(r.Context(), contextKeyUser, user)

		// Pass the request with the new context to the next handler in the chain
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
