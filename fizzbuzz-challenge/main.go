package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) {
	 // Lakukan looping dari 1 hingga n
	 for i := 1; i <= n; i++ {
        // Periksa apakah i adalah kelipatan 3 dan/atau 5
		if i%3 == 0 && i%5 == 0 {
            fmt.Println("FizzBuzz")
        } else if i%3 == 0 {
            fmt.Println("Fizz")
        } else if i%5 == 0 {
            fmt.Println("Buzz")
		
		// Jika i bukan kelipatan 3 atau 5, cetak nilainya
        } else {
            fmt.Println(i)
        }
    }
}

func main() {
	// Tentukan nilai maksimal looping
	fmt.Print("Enter a value for n: ")
	var input string
	fmt.Scanln(&input)

	// Ubah input string jadi integer
	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		return
	}

	// Panggil looping function dengan n sebagai parameter
	fizzBuzz(n)
}
