package main

import (
	"classmate-finder/helpers"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Mendefinisikan struktur untuk data teman sekelas
type Classmate struct {
	ID           int
	Name         string
	Address      string
	Occupation   string
	ReasonToGo   string
}

func main() {
    // Baca file menggunakan ReadFileJSON function dari "helpers" package
    classmatesJSON, err := helpers.ReadFileJSON("Classmates.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Parse data JSON menjadi slice dari struct Classmate 
	var classmates []Classmate
	err = json.Unmarshal(classmatesJSON, &classmates)
	if err != nil {
		fmt.Println("Error parsing JSON data:", err)
		return
	}
	
	// Dapatkan argumen dari baris perintah
	args := os.Args[1:]

	// Periksa jika tidak ada argumen yang diberikan
	if len(args) == 0 {
		fmt.Println("Silakan berikan Nama atau ID sebagai argumen!")
		return
	}

	// Panggil fungsi searchClassmate
	searchClassmates(args, classmates)
}
func searchClassmates(args []string, classmates []Classmate) {
	for _, arg := range args {
		// Parse argumen sebagai bilangan bulat (ID)
		id, err := strconv.Atoi(arg)
		if err != nil {
			// Tangani jika argumen bukan bilangan bulat (nama)
			found := false
			for _, classmate := range classmates {
				if arg == classmate.Name || arg == fmt.Sprint(classmate.ID) {
					displayClassmate(classmate)
					found = true
					break
				}
			}
			if !found {
				fmt.Printf("Teman sekelas dengan nama '%s' tidak ditemukan.\n", arg)
			}
		} else if id >= 1 && id <= len(classmates) {
			// Tangani jika argumen adalah indeks yang valid
			classmate := classmates[id-1]
			displayClassmate(classmate)
		} else {
			// Tangani jika indeks berada di luar jangkauan
			fmt.Printf("Indeks %d berada di luar jangkauan. Jangkauan yang valid adalah index 1 sampai %d.\n", id, len(classmates))
		}
	}
}

// displayClassmate menampilkan data teman sekelas
func displayClassmate(classmate Classmate) {
	fmt.Println("ID:", classmate.ID)
	fmt.Println("Nama:", classmate.Name)
	fmt.Println("Alamat:", classmate.Address)
	fmt.Println("Pekerjaan:", classmate.Occupation)
	fmt.Println("Alasan memilih Go:", classmate.ReasonToGo)
}
