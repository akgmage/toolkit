package main

import "net/http"

func main() {
	mux := routes()
}

func routes() http.Handler{
	mux := http.NewServeMux()
	// register handler for default page
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	// register handler for uplaodFiles
	mux.HandleFunc("/upload", uploadFiles)
	// register handler for uploadOneFile
	mux.HandleFunc("/upload-one", uploadOneFile)

	return mux
	
}
// upload multiple files
func uploadFiles(w http.ResponseWriter, r *http.Request) {

}
// upload one file
func uploadOneFile(w http.ResponseWriter, r *http.Request) {

}