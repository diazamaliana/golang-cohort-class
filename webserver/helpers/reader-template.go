package helpers

import (
	"html/template"
	"net/http"
	"os"
)

// ReadFileJSON membaca file JSON dari path yang ditentukan dan mengembalikan datanya.
func ReadFileJSON(path string,filename string) ([]byte, error) {
    dataPath := path + filename

    data, err := os.ReadFile(dataPath)
    if err != nil {
        return nil, err
    }
    return data, nil
}

// ReadFileHTML membaca dan merender file HTML dari path yang ditentukan ke ResponseWriter.
func ReadFileHTML(w http.ResponseWriter, path string, tmplFile string, data interface{}) {
    tmplPath := path + tmplFile 

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tmpl.Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
