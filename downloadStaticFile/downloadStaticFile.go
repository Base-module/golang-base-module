package downloadStaticFile

import (
	"fmt"
	"net/http"
	"os"
)

// DownloadStaticFile a file and tries to force the browser to avoid displaying it in the browser window by setting content disposition.
// It also allows specification of the display name
func DownloadStaticFile(w http.ResponseWriter, r *http.Request, pathName, displayName string) {
	// Check if file exists before serving
	if _, err := os.Stat(pathName); os.IsNotExist(err) {
		http.Error(w, fmt.Sprintf("file not found: %s", pathName), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", displayName))
	http.ServeFile(w, r, pathName)
}
