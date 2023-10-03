package main

import (
	"fmt"
	"sync"
)

func main() {
	// Deklarasi WaitGroup untuk menunggu semua goroutine selesai
	var wg sync.WaitGroup

	// Mutex digunakan untuk mengamankan akses ke data yang bersamaan oleh goroutine
	var mu sync.Mutex

	// Buat dua slice interface dengan data yang berbeda
	interface1 := make([]interface{}, 3)
	interface2 := make([]interface{}, 3)

	for i := 0; i < 3; i++ {
		interface1[i] = fmt.Sprintf("coba%d", i+1)
		interface2[i] = fmt.Sprintf("bisa%d", i+1)
	}

	// Simulasikan 4 goroutine untuk masing-masing dari dua tipe data
	for i := 0; i < 4; i++ {
		// Menambahkan 2 ke WaitGroup, karena ada 2 goroutine yang akan dijalankan dalam setiap iterasi
		wg.Add(2)

		// Mengunci mutex untuk menghindari akses bersamaan ke data oleh goroutine
		mu.Lock()

		// Menjalankan goroutine untuk tipe data interface1
		go arranged(&wg, interface1, i+1)

		// Mengunci mutex lagi sebelum menjalankan goroutine berikutnya
		mu.Lock()

		// Menjalankan goroutine untuk tipe data interface2
		go arranged(&wg, interface2, i+1)
	}

	// Menunggu semua goroutine selesai
	wg.Wait()
}

func arranged(wg *sync.WaitGroup, data []interface{}, id int) {
	// Deklarasi Mutex di dalam goroutine untuk mengamankan akses ke data
	var mu sync.Mutex

	// Menggunakan defer untuk memastikan Mutex di-unlock ketika goroutine selesai
	defer mu.Unlock()
	defer wg.Done()

	// Menampilkan pesan dengan ID goroutine dan data yang dioperasikan
	fmt.Printf("Goroutine %d: %+v\n", id, data)
}
