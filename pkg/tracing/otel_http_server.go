package tracing

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const (
	attrRequestQuery = "http.request.query"
	attrRequestBody  = "http.request.body"
	attrResponseBody = "http.response.body"
)

// Middleware traces request query string and body
func HTTPRequestTrace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		span := trace.SpanFromContext(r.Context())
		traceRequest(span, r)

		next.ServeHTTP(w, r)
	})
}

func traceRequest(span trace.Span, r *http.Request) {
	if span == nil {
		return
	}

	if r.URL.RawQuery != "" {
		attr := attribute.String(attrRequestQuery, r.URL.RawQuery)
		span.SetAttributes(attr)
	}

	rBody := r.Body

	body, err := io.ReadAll(rBody)
	if err != nil {
		return
	}
	defer rBody.Close()

	if len(body) != 0 {
		attr := attribute.String(attrRequestBody, string(body))
		span.SetAttributes(attr)
	}

	r.Body = io.NopCloser(bytes.NewReader(body))
}

// Constructs middleware that traces response body. Truncates the output according to the provided bytes count limit
func NewHTTPResponseTrace(bytesLimit int) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			span := trace.SpanFromContext(r.Context())
			if span == nil {
				next.ServeHTTP(w, r)

				return
			}

			bw := newBodyResponseWriter(w)

			next.ServeHTTP(bw, r)

			decoded := decode(bw.Header(), bw.body)
			truncated := truncate(decoded, bytesLimit)
			attr := attribute.String(attrResponseBody, string(truncated))
			span.SetAttributes(attr)
		})
	}
}

type bodyResponseWriter struct {
	http.ResponseWriter
	body []byte
}

func (brw *bodyResponseWriter) Write(buf []byte) (int, error) {
	n, err := brw.ResponseWriter.Write(buf)
	if err != nil {
		return 0, err //nolint:wrapcheck
	}

	brw.body = append(brw.body[:len(brw.body):len(brw.body)], buf...)

	return n, nil
}

func newBodyResponseWriter(w http.ResponseWriter) *bodyResponseWriter {
	return &bodyResponseWriter{
		ResponseWriter: w,
	}
}
