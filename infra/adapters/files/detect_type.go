package files

import (
	"bytes"
	"io"
	"net/http"
	"strings"
)

func DetectFileType(r io.Reader) string {
	buf := make([]byte, 512)
	n, err := r.Read(buf)
	if err != nil && err != io.EOF {
		return "unknown"
	}

	r = io.NopCloser(io.MultiReader(bytes.NewReader(buf[:n]), r))

	mimeType := http.DetectContentType(buf)
	mimeType = strings.Split(mimeType, ";")[0]

	switch mimeType {
	case "application/json":
		return "json"
	case "text/csv", "text/plain":
		return "csv"
	case "application/vnd.ms-excel",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":
		return "xlsx"
	default:
		return "unknown"
	}
}
