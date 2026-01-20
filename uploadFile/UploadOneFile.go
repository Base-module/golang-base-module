package uploadFile

import (
	"errors"
	"net/http"
)

func (t *UploadTools) UploadOneFile(r *http.Request, uploadDir string, rename ...bool) (*UploadedFile, error) {
	renameFile := true
	if len(rename) > 0 {
		renameFile = rename[0]
	}

	files, err := t.UploadFiles(r, uploadDir, renameFile)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, errors.New("no files uploaded")
	}

	return files[0], nil
}
