package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		res := &response{Message: "Hello AWSome Builders!  Hello Octank Testing for BMW demo!"}

		for _, e := range os.Environ() {
			pair := strings.Split(e, "=")
			res.EnvVars = append(res.EnvVars, pair[0]+"="+pair[1])
		}
		sort.Strings(res.EnvVars)

		res.Message2 = "Inside main() of Go program inside the container image on Friday 9/3"

		// Beautify the JSON output
		out, _ := json.MarshalIndent(res, "", "  ")

		// Normally this would be application/json, but we don't want to prompt downloads
		w.Header().Set("Content-Type", "text/plain")

		io.WriteString(w, string(out))

		fmt.Println("Hello - the log message")
	})
	http.ListenAndServe(":8080", nil)
}

type response struct {
	Message string   `json:"message1"`
	EnvVars []string `json:"env"`
	Message2 string  `json:"message2"`
}
