package main

import (
	"fmt"
	"sync"
)

func main() {
    // Deklarasi WaitGroup untuk menunggu semua goroutine selesai
	var wg sync.WaitGroup

	// Buat dua slice interface dengan data yang berbeda
	interface1 := make([]interface{}, 3)
	interface2 := make([]interface{}, 3)

	for i := 0; i < 3; i++ {
		interface1[i] = fmt.Sprintf("coba%d", i+1)
		interface2[i] = fmt.Sprintf("bisa%d", i+1)
	}
	// Simulasikan 4 goroutine untuk masing-masing dari dua tipe data
	for i := 0; i < 4; i++ {
		wg.Add(2)
		go disorder(&wg, interface1, i+1)
		go disorder(&wg, interface2, i+1)
	}

	wg.Wait()
}

func disorder(wg *sync.WaitGroup, data []interface{}, id int) {
	defer wg.Done()

	fmt.Printf("Goroutine %d: %+v\n", id, data)
}