package toolkit

import (
	"crypto/rand"
	"net/http"
)

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJLKMNOPQRSTUVWXYZ0123456789_+"
// Tools is the type used to instantiate this module, 
// Any variable of this type will have access to all the methods with the receiver *Tools
type Tools struct {
	MaxFileSize int
}
// RandomString generates a random string of length n using randomStringSource as
// the source for the string
func (t *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(randomStringSource)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x % y]
	}
	return string(s)
}
// UploadedFile is a struct used to save info about an uploaded file
type UploadedFile struct {
	NewFileName 		string
	OriginalFileName 	string
	FileSize 			int64
}

func  (t *Tools) UploadFiles(r *http.Request, uploadDir string, rename ...bool) ([]*UploadedFile, error) {
	renameFile := true
	if len(rename) > 0 {
		renameFile = rename[0]
	}

	var uploadedFile []*UploadedFile

	if t.MaxFileSize == 0 {
		t.MaxFileSize = 1024 * 1024 * 1024
	}

	err := r.ParseMultipartForm(int64(t.MaxFileSize))
}