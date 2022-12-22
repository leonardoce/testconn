package testconn

import (
	"log"
	"net/http"
	"time"
)

// LoggingDecorator insert a new logging line for every request
type LoggingDecorator struct {
	inner http.Handler
}

// NewLoggingDecorator creates a new logging decorator handler
func NewLoggingDecorator(handler http.Handler) LoggingDecorator {
	return LoggingDecorator{
		inner: handler,
	}
}

// ServeHTTP implements the http.Handler interface
func (d LoggingDecorator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	writer := TrackingResponseWriter{
		inner: w,
	}

	before := time.Now()
	d.inner.ServeHTTP(&writer, r)
	elapsed := time.Since(before)

	log.Printf("%6dms [%s] %v %s", elapsed.Milliseconds(), r.RemoteAddr, writer.statusCode, r.RequestURI)
}

// TrackingResponseWriter implements the ResponseWriter interface
// keeping track of the status code written in the response, and then
// delegating the actual response writer to an inner writer
type TrackingResponseWriter struct {
	statusCode int
	inner      http.ResponseWriter
}

// Header implements the http.ResponseWriter interface
func (w *TrackingResponseWriter) Header() http.Header {
	return w.inner.Header()
}

// Write implements the http.ResponseWriter intarface
func (w *TrackingResponseWriter) Write(data []byte) (int, error) {
	if w.statusCode == 0 {
		w.statusCode = http.StatusOK
		w.inner.WriteHeader(http.StatusOK)
	}

	return w.inner.Write(data)
}

// WriteHeader implements the http.ResponseWriter intarface
func (w *TrackingResponseWriter) WriteHeader(statusCode int) {
	if w.statusCode == 0 {
		w.statusCode = statusCode
	}

	w.inner.WriteHeader(statusCode)
}
