package toolkit

import (
	"image"
	"image/png"
	"io"
	"mime/multipart"
	"os"
	"sync"
	"testing"
)

// Function start with name Test followed by receiver Tools
// and _ and name of function i.e. RandomString
// takes a param t of type *testing.T (pointer to testing.T)
func TestTools_RandomString(t *testing.T) {
	var testTools Tools

	s := testTools.RandomString(10)
	if len(s) != 10 {
		t.Error("Wrong length random string returned")
	}
}

var uploadTests = []struct{
	name string
	allowedTypes []string
	renameFile bool
	errorExpected bool
}{
	{name: "allowed no rename", allowedTypes: []string{"image/jpeg", "image/png"}, renameFile: false, errorExpected: false},
}

func TestTools_UploadFiles(t *testing.T) {
	for _, e := range uploadTests{
		// set up pipe to avoid buffering
		pr, pw := io.Pipe()
		// make sure things occur in particular sequence
		writer := multipart.NewWriter(pw)
		wg := sync.WaitGroup{}
		wg.Add(1)
		// fire off goroutine in background
		go func() {
			defer writer.Close()
			defer wg.Done()

			// create form data field file
			part, err := writer.CreateFormFile("file", "./testdata/img.png")
			if err != nil {
				t.Error(err)
			}

			f, err := os.Open("./testdata/img.png")
			if err!= nil {
				t.Error(err)
			}
			defer f.Close()

			img, _, err := image.Decode(f)
			if err != nil {
				t.Error("error decoding image", err)
			}

			err = png.Encode(part, img)
			if err != nil {
				t.Error(err)
			}
		}()
	}
}