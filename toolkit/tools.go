package toolkit

import (
	"crypto/rand"
	"errors"
	"net/http"
	"strings"
)

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJLKMNOPQRSTUVWXYZ0123456789_+"
// Tools is the type used to instantiate this module, 
// Any variable of this type will have access to all the methods with the receiver *Tools
type Tools struct {
	MaxFileSize int
	AllowedFileTypes []string
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

	var uploadedFiles []*UploadedFile

	if t.MaxFileSize == 0 {
		t.MaxFileSize = 1024 * 1024 * 1024
	}

	err := r.ParseMultipartForm(int64(t.MaxFileSize))

	if err != nil {
		return nil, errors.New("The uploaded file is too big")
	}

	for _, fHeaders := range r.MultipartForm.File {
		for _, hdr := range fHeaders {
			uploadedFiles, err = func(uploadedFiles []*UploadedFile) ([]*UploadFiles, error) {
				var uploadedFile UploadedFile
				infile, err := hdr.Open()
				if err != nil {
					return nil, err
				}
				defer infile.Close()
				buff := make([]byte, 512)
				_, err = infile.Read(buff)
				if err != nil {
					return nil, err
				}
				// TODO: check to see if file type is permitted
				allowed := false
				fileType := http.DetectContentType(buff)
				// Only file types allowed
				allowedTypes := []string{"image/jpeg", "image/png", "image/gif"}
				// atleast one value is in there if it passes
				if len(t.AllowedFileTypes) > 0 {
					// range through allowedTypes and do a comparision
					for _, x := range t.AllowedFileTypes {
						// is file type of currently uploaded file equal to one of the 
						// things in the slice of string allowedTypes
						if strings.EqualFold(fileType, x) {
							allowed = true
						}
					}
				} else {
					allowed = true
				}

				if !allowed {
					return nil, errors.New("the uploaded file type is not permitted")
				}

			}(uploadedFiles)
		}
	}
}