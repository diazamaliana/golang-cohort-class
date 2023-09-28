package helpers

import (
    "os"
)

// ReadFileJSON membaca file JSON dan mengembalikan kontennya sebagai potongan byte.
func ReadFileJSON(filename string) ([]byte, error) {
    // Baca JSON file
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return data, nil
}
