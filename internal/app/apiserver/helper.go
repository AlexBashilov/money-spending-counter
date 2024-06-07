package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func covertIntArrToStringarr(year, month, day int) string {
	intArr := [3]int{
		year,
		month,
		day,
	}

	var buffer bytes.Buffer

	for i := 0; i < len(intArr); i++ {
		buffer.WriteString(strconv.Itoa(intArr[i]))
		if i != len(intArr)-1 {
			buffer.WriteString("-")
		}
	}
	return buffer.String()
}
