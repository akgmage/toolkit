package main

import (
	"log"
	"net/http"
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
	}
}
// upload one file
func uploadOneFile(w http.ResponseWriter, r *http.Request) {

}