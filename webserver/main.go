package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webserver/helpers"
)

// Struct User untuk merepresentasikan seorang pengguna
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Data statis pengguna
var users []User

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
	loadDataFromJSON("data/", "users.json")

	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/profile", profileHandler)

    fmt.Println("Server berjalan di http://localhost:9090")
	http.ListenAndServe(":9090", nil)
}

// loadDataFromJSON membaca data JSON dari sebuah file dan mengisi slice yang diberikan.
func loadDataFromJSON(path, filename string) error {
	// Baca file json menggunakan fungsi ReadFileJSON
    jsonData, err := helpers.ReadFileJSON(path, filename)
    if err != nil {
        return fmt.Errorf("Error reading %s: %v", filename, err)
    }

	// Unmarshal data JSON ke dalam slice users atau biodata
    if err := json.Unmarshal(jsonData, &users); err != nil {
        return fmt.Errorf("Error unmarshaling %s: %v", filename, err)
    }

    return nil
}

// loginHandler menangani halaman login dan otentikasi.
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		// Periksa apakah email dan password cocok
		for _, user := range users {
			if user.Email == email && user.Password == password {
				http.Redirect(w, r, "/profile?email="+email, http.StatusSeeOther)
				return
			}
		}

		fmt.Fprintln(w, "Login gagal. Email atau password salah.")
		return
	}

	helpers.ReadFileHTML(w, "pages/", "login.html", nil)
}

// profileHandler menangani halaman profil pengguna.
func profileHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	bio, ok := biodata[email]
	if !ok {
		http.Error(w, "Email tidak ditemukan", http.StatusNotFound)
		return
	}

	helpers.ReadFileHTML(w, "pages/", "profile.html", bio)
}
