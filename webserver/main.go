package main

import (
	"fmt"
	"net/http"
	"html/template"
)

// Data statis pengguna
var users = map[string]string{
	"john@example.com": "password123",
	"jane@example.com": "secret456",
}

// Data biodata pengguna
var biodata = map[string]map[string]string{
	"john@example.com": {
		"name":    "John Doe",
		"address":  "123 Jalan Contoh, Kota Contoh",
		"phone": "123-456-7890",
	},
	"jane@example.com": {
		"name":    "Jane Smith",
		"address":  "456 Jalan Contoh, Kota Contoh",
		"phone": "987-654-3210",
	},
}

func main() {
	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/profile", profileHandler)

    fmt.Println("Server berjalan di http://localhost:9090")
	http.ListenAndServe(":9090", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		// Periksa apakah email dan password cocok
		if storedPassword, ok := users[email]; ok && storedPassword == password {
			http.Redirect(w, r, "/profile?email="+email, http.StatusSeeOther)
			return
		} else {
			fmt.Fprintln(w, "Login gagal. Email atau password salah.")
			return
		}
	}

	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Login</title>
	</head>
	<body>
		<h1>Login</h1>
		<form method="POST" action="/">
			<label>Email:</label>
			<input type="text" name="email"><br>
			<label>Password:</label>
			<input type="password" name="password"><br>
			<input type="submit" value="Login">
		</form>
	</body>
	</html>
	`

	t, err := template.New("login").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	bio, ok := biodata[email]
	if !ok {
		http.Error(w, "Email tidak ditemukan", http.StatusNotFound)
		return
	}

	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Profil</title>
	</head>
	<body>
		<h1>Profil</h1>
		<h2>name: {{.name}}</h2>
		<h2>address: {{.address}}</h2>
		<h2>phone: {{.phone}}</h2>
	</body>
	</html>
	`

	t, err := template.New("profile").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, bio)
}
