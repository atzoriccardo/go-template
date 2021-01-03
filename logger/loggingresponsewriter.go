package logger

import (
	"net/http"
)

type logginResponseWriter struct {
	writer     http.ResponseWriter
	statusCode int
	length     int
}

func (r *logginResponseWriter) WriteHeader(code int) {
	r.statusCode = code
	r.writer.WriteHeader(code)
}

func (r *logginResponseWriter) Write(b []byte) (int, error) {
	n, err := r.writer.Write(b)

	if err != nil {
		return n, err
	}

	r.length += n

	return n, err
}

func (r *logginResponseWriter) Header() http.Header {
	return r.writer.Header()
}

func (r *logginResponseWriter) Length() int {
	return r.Length()
}
