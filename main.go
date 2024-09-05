package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to MySQL database
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/users_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Simple login handler
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// Get username and password from form data
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Insecure SQL query vulnerable to SQL injection
		query := fmt.Sprintf("SELECT * FROM users WHERE username = '%s' AND password = '%s'", username, password)

		// Execute the query
		row := db.QueryRow(query)

		// Check if a row is returned
		var id int
		var user, pass string
		err := row.Scan(&id, &user, &pass)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// If no error, login is successful
		fmt.Fprintf(w, "Login successful for user: %s", username)
	})

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
