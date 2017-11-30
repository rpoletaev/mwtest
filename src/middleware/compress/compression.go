package compress

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

const (
	accEncHeader  = "Accept-Encoding"
	contEncHeader = "Content-Encoding"
	gzString      = "gzip"
)

type stub struct{}

func New() stub {
	return stub{}
}

type gzipWriter struct {
	io.Writer
	http.ResponseWriter
}

func (gzw gzipWriter) Write(content []byte) (int, error) {
	return gzw.Writer.Write(content)
}

func (s stub) Compression(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get(accEncHeader), gzString) {
			next.ServeHTTP(w, r)
			return
		}

		w.Header().Set(contEncHeader, gzString)
		gzWriter := gzip.NewWriter(w)
		defer gzWriter.Close()

		gzResponseWriter := gzipWriter{
			Writer:         gzWriter,
			ResponseWriter: w,
		}

		next.ServeHTTP(gzResponseWriter, r)
	})
}
