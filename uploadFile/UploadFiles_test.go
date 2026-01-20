package uploadFile

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"testing"
)

var uploadTests = []struct {
	name          string
	allowedTypes  []string
	renameFile    bool
	errorExpected bool
}{
	{name: "allowed no rename", allowedTypes: []string{"image/jpeg", "image/png"}, renameFile: false, errorExpected: false},
	{name: "allowed rename", allowedTypes: []string{"image/jpeg", "image/png"}, renameFile: true, errorExpected: false},
	{name: "not allowed", allowedTypes: []string{"image/gif"}, renameFile: false, errorExpected: true},
}

func TestTools_UploadFiles(t *testing.T) {
	for _, e := range uploadTests {
		t.Run(e.name, func(t *testing.T) {
			pr, pw := io.Pipe()
			writer := multipart.NewWriter(pw)

			go func() {
				defer writer.Close()
				part, err := writer.CreateFormFile("file", "test.jpg")
				if err != nil {
					t.Error(err)
					return
				}

				f, err := os.Open("../testdata/ABCD.jpg")
				if err != nil {
					t.Error(err)
					return
				}
				defer f.Close()

				_, err = io.Copy(part, f)
				if err != nil {
					t.Error(err)
				}
			}()

			request := httptest.NewRequest("POST", "/", pr)
			request.Header.Add("Content-Type", writer.FormDataContentType())

			var testTools UploadTools
			testTools.AllowedFileTypes = e.allowedTypes

			uploadedFiles, err := testTools.UploadFiles(request, "./testdata/uploads/", e.renameFile)

			if e.errorExpected {
				if err == nil {
					t.Error("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			filePath := fmt.Sprintf("./testdata/uploads/%s", uploadedFiles[0].NewFileName)
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				t.Errorf("expected file to exist: %s", err.Error())
			}

			// clean up
			_ = os.Remove(filePath)
		})
	}
}
