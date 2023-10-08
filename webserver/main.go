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

type Biodata struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Age     int    `json:"age"`
}

// Data statis pengguna
var users []User
var biodata []Biodata

func main() {
	loadDataFromJSON("data/", "users.json", &users)
	loadDataFromJSON("data/", "biodata.json", &biodata)

	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/profile", profileHandler)

	fmt.Println("Server berjalan di http://localhost:9090")
	http.ListenAndServe(":9090", nil)
}

// loadDataFromJSON membaca data JSON dari sebuah file dan mengisi slice yang diberikan.
func loadDataFromJSON(path, filename string, data interface{}) error {
	// Baca file json menggunakan fungsi ReadFileJSON
	jsonData, err := helpers.ReadFileJSON(path, filename)
	if err != nil {
		return fmt.Errorf("Error reading %s: %v", filename, err)
	}

	// Unmarshal data JSON ke dalam slice users atau biodata
    if err := json.Unmarshal(jsonData, data); err != nil {
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
	var userBio Biodata
	found := false
	for _, bio := range biodata {
		if bio.Email == email {
			userBio = bio
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Email not found", http.StatusNotFound)
		return
	}

	helpers.ReadFileHTML(w, "pages/", "profile.html", userBio)
}
