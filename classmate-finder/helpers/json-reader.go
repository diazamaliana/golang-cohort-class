package helpers

import (
    "os"
)

// ReadFileJSON reads a JSON file and returns its content as a byte slice.
func ReadFileJSON(filename string) ([]byte, error) {
    // Read the JSON file
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return data, nil
}
