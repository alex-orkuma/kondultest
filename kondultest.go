package kondultest

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

// HelloHTTP cloud function sample
func HelloHTTP(rw http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(rw, "Hello World!")
		return
	}
	if d.Name == "" {
		fmt.Fprint(rw, "Hello World!")
		return
	}
	fmt.Fprint(rw, "Hello", html.EscapeString(d.Name))
}
