package tracing

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
)

const contentEncodingHeaderKey = "Content-Encoding"

func decode(header http.Header, body []byte) []byte {
	if header.Get(contentEncodingHeaderKey) == "gzip" {
		reader, err := gzip.NewReader(bytes.NewBuffer(body))
		if err != nil {
			return body
		}
		defer reader.Close()

		decoded, err := io.ReadAll(reader)
		if err != nil {
			return body
		}

		return decoded
	}

	return body
}

func truncate(body []byte, limit int) []byte {
	if limit < 0 || len(body) <= limit {
		return body
	}

	res := body[:limit]

	return res
}
