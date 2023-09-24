package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Minta pengguna untuk memasukkan kalimat
    fmt.Print("Masukkan kalimat: ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    sentence := scanner.Text()

    // Split kalimat menjadi array kata-kata
    words := splitSentence(sentence)

    // Mencetak setiap karakter dari setiap kata dalam array
    fmt.Println("kalimat:", `"`+sentence+`"`)

    // Inisialisasi array untuk menyimpan karakter-karakter
    characters := splitCharacters(words)

    // Melakukan perhitungan munculnya setiap karakter, termasuk spasi
    countChar := make(map[string]int)
    for _, char := range characters {
        countChar[char]++
    }

    // Mencetak hasil perhitungan
    fmt.Println(countChar)
}

func splitSentence(sentence string) []string {
	var words []string
	var word string
	inWord := false

	for _, char := range sentence {
		if char == ' ' {
			if inWord {
				words = append(words, word)
				word = ""
				inWord = false
			}
		} else {
			word += string(char)
			inWord = true
		}
	}

	if inWord {
		words = append(words, word)
	}

	return words
}

func splitCharacters(words []string) []string {
	var characters []string

    for _, word := range words {
        for i, char := range word {
            char := string(char)
            characters = append(characters, char)
            fmt.Println(char)

            // Mengecek apakah karakter ini adalah spasi terakhir dalam kalimat
            if i == len(word)-1 && word == words[len(words)-1] {
                break
            }
        }

        // Menambahkan spasi antara kata-kata
        if word != words[len(words)-1] {
            characters= append(characters, " ")
            fmt.Println(" ")
        }
    }

    return characters
}
