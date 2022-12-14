package toolkit

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"mime/multipart"
	"net/http/httptest"
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
	{name: "allowed rename", allowedTypes: []string{"image/jpeg", "image/png"}, renameFile: true, errorExpected: false},
	{name: "not allowed", allowedTypes: []string{"image/jpeg",}, renameFile: true, errorExpected: true},
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

		// read from pipe which receives data

		request := httptest.NewRequest("POST", "/", pr)
		// sets the  correct content type for the payload   
		request.Header.Add("Content-Type", writer.FormDataContentType())

		var testTools Tools
		testTools.AllowedFileTypes = e.allowedTypes

		uploadedFiles, err := testTools.UploadFiles(request, "./testdata/uploads/", e.renameFile)
		if err != nil && !e.errorExpected{
			t.Error(err)
		}

		if !e.errorExpected {
			// Stat returns a FileInfo describing the named file. If there is an error, it will be of type *PathError
			if _, err := os.Stat(fmt.Sprintf("./testdata/uploads/%s", uploadedFiles[0].NewFileName)); os.IsNotExist(err) {
				// Errorf is equivalent to Logf followed by Fail.
				t.Errorf("%s: expected file to exist: %s", e.name, err.Error())
			}
			// Remove removes the named file or directory. If there is an error, it will be of type *PathError.
			_ = os.Remove(fmt.Sprintf("./testdata/uploads/%s", uploadedFiles[0].NewFileName))
		}

		if !e.errorExpected && err != nil {
			t.Errorf("%s: error expected but none received", e.name)
		}
		wg.Wait()
	}
}

func TestTools_UploadOneFile(t *testing.T) {
		// set up pipe to avoid buffering
		pr, pw := io.Pipe()
		// make sure things occur in particular sequence
		writer := multipart.NewWriter(pw)
		// fire off goroutine in background
		go func() {
			defer writer.Close()
		
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

		// read from pipe which receives data

		request := httptest.NewRequest("POST", "/", pr)
		// sets the  correct content type for the payload   
		request.Header.Add("Content-Type", writer.FormDataContentType())

		var testTools Tools


		uploadedFiles, err := testTools.UploadOneFile(request, "./testdata/uploads/", true)
		if err != nil {
			t.Error(err)
		}

		// Stat returns a FileInfo describing the named file. If there is an error, it will be of type *PathError
		if _, err := os.Stat(fmt.Sprintf("./testdata/uploads/%s", uploadedFiles.NewFileName)); os.IsNotExist(err) {
			// Errorf is equivalent to Logf followed by Fail.
			t.Errorf("expected file to exist: %s", err.Error())
		}
		// Remove removes the named file or directory. If there is an error, it will be of type *PathError.
		_ = os.Remove(fmt.Sprintf("./testdata/uploads/%s", uploadedFiles.NewFileName))


}

func TestTools_CreateDirIfNotExist(t *testing.T) {
	var testTool Tools

	err := testTool.CreateDirIfNotExist("./testdata/mydir")

	if err != nil {
		t.Error(err)
	}

	err = testTool.CreateDirIfNotExist("./testdata/mydir")

	if err != nil {
		t.Error(err)
	}
	_ = os.Remove("./testdata/mydir")
}

var slugTests = []struct {
	name string
	s string
	expected string
	errorExpected bool
} {
	{name: "valid string", s: "now is the time", expected: "now-is-the-time", errorExpected: false},
	{name: "", s: "now is the time", expected: "", errorExpected: true},
	{name: "complex string", s: "Now is the time! + Goal &123 ^", expected: "now-is-the-time-goal-123", errorExpected: true},
	{name: "Japanese string", s: "?????????????????????", expected: "", errorExpected: true},
	{name: "Japanese string", s: "?????????????????????hello world", expected: "hello-world", errorExpected: false},
}

func TestTools_Slugify(t *testing.T) {
	var testTool Tools

	for _, e := range slugTests {
		slug, err := testTool.Slugify(e.s)
		if err != nil && !e.errorExpected {
			t.Errorf("%s: error received when none expected %s", e.name, err.Error())
		}
		if !e.errorExpected && slug != e.expected {
			t.Errorf("%s: wrong slug returned, expected %s but got %s", e.name, e.expected, slug)
		}
	}
}