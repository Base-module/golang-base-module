package uploadFile

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"testing"
)

func TestTools_UploadOneFile(t *testing.T) {
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
	uploadedFiles, err := testTools.UploadOneFile(request, "./testdata/uploads/", true)
	if err != nil {
		t.Error(err)
		return
	}

	filePath := fmt.Sprintf("./testdata/uploads/%s", uploadedFiles.NewFileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("expected file to exist: %s", err.Error())
	}

	_ = os.Remove(filePath)
}
