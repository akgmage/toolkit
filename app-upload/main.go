package main

import "net/http"

func main() {
	mux := routes()
}

func routes() http.Handler{
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	
	
}
// upload multiple files
func uploadFiles(w http.ResponseWriter, r *http.Request) {

}
// upload one file
func uploadOneFile(w http.ResponseWriter, r *http.Request) {

}