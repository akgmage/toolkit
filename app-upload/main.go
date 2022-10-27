package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akgmage/toolkit"
)

func main() {
	mux := routes()

	log.Println("Starting server on port 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
// uses standard library to handle 3 routes
func routes() http.Handler{
	mux := http.NewServeMux()
	// register handler for default route
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	// register handler for uplaodFiles
	mux.HandleFunc("/upload", uploadFiles)
	// register handler for uploadOneFile
	mux.HandleFunc("/upload-one", uploadOneFile)

	return mux
	
}
// upload multiple files
func uploadFiles(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	t := toolkit.Tools{
		MaxFileSize: 1024 * 1024 *1024,
		AllowedFileTypes: []string{"image/jpeg", "image/png", "image/gif"},
	}

	files, err := t.UploadFiles(r, "./uploads")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	out := ""
	for _, item := range files {
		out += fmt.Sprintf("Uploaded %s to the uploads folder, renamed to %s\n", item.OriginalFileName, item.NewFileName)
	}
	_, _ = w.Write([]byte(out))
}
// upload one file
func uploadOneFile(w http.ResponseWriter, r *http.Request) {

}