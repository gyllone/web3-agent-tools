package dependencies

import (
	"compress/flate"
	"compress/gzip"
	"io"
	"net/http"
)

func DecompressResponse(resp *http.Response) (io.ReadCloser, error) {
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return resp.Body, err
		}
		return reader, nil
	case "deflate":
		reader := flate.NewReader(resp.Body)
		return reader, nil
	default:
		// No compression or unsupported compression format
		return resp.Body, nil
	}
}
