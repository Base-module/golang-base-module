package downloadStaticFile

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

func TestTools_DownloadStaticFile(t *testing.T) {
	filePath := "../testdata/ABCD.jpg"
	r := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	// Get actual file size dynamically
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		t.Fatalf("failed to stat test file: %v", err)
	}
	expectedSize := strconv.FormatInt(fileInfo.Size(), 10)

	DownloadStaticFile(r, req, filePath, "CodeList.jpg")

	res := r.Result()
	defer res.Body.Close()

	actualSize := res.Header.Get("Content-Length")
	if actualSize == "" {
		t.Error("Content-Length header is missing")
	} else if actualSize != expectedSize {
		t.Errorf("wrong content length; expected %s, got %s", expectedSize, actualSize)
	}

	if res.Header["Content-Disposition"][0] != "attachment; filename=\"CodeList.jpg\"" {
		t.Error("wrong content disposition")
	}

	_, err = io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
}

func TestTools_DownloadStaticFile_NotFound(t *testing.T) {
	r := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	DownloadStaticFile(r, req, "./testdata/nonexistent.jpg", "CodeList.jpg")

	res := r.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected status code 404, got %d", res.StatusCode)
	}
}
