package mygo

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Params  map[string]string
}

// Sends a JSON response
func (context Context) Json(status int, data interface{}) {
	context.Writer.Header().Set("Content-Type", "application/json")
	context.Writer.WriteHeader(status)
	json.NewEncoder(context.Writer).Encode(data)
}

// sends plain text
func (context Context) String(status int, text string) {
	context.Writer.Header().Set("Content-Type", "text/plain")
	context.Writer.WriteHeader(status)
	context.Writer.Write([]byte(text))
}

// sends status code only
func (context Context) Status(code int) {
	context.Writer.WriteHeader(code)
}

// reads JSON body into a struct
func (context Context) BindJson(output interface{}) error {
	decoder := json.NewDecoder(context.Request.Body)
	return decoder.Decode(output)
}
