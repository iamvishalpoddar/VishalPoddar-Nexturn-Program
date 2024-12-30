package middleware

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func AuthorizationMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
			fmt.Println("Missing or invalid Authorization header")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Decode the Base64-encoded credentials
		payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(authHeader, "Basic "))
		if err != nil {
			fmt.Println("Failed to decode Authorization header:", err)
			http.Error(w, "Invalid Authorization Header", http.StatusUnauthorized)
			return
		}

		// Split the username and password
		credentials := strings.SplitN(string(payload), ":", 2)
		if len(credentials) != 2 {
			fmt.Println("Invalid credentials format:", string(payload))
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}

		username, password := credentials[0], credentials[1]
		fmt.Printf("Received username: %s, password: %s\n", username, password)

		if username != "admin" || password != "admin" {
			fmt.Println("Invalid username or password")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Validate credentials against the database
		// var storedPassword string
		// query := "SELECT password FROM users WHERE username = ?"
		// err = db.QueryRow(query, username).Scan(&storedPassword)
		// if err != nil {
		// 	fmt.Println("User not found or error querying database:", err)
		// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
		// 	return
		// }

		// if storedPassword != password {
		// 	fmt.Println("Password mismatch")
		// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
		// 	return
		// }

		fmt.Println("Authentication successful")
		next.ServeHTTP(w, r)
	})
}
