package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", resourceHandler)
	e := http.ListenAndServe(":8888", nil)
	if e != nil {
		fmt.Println(e.Error())
	}

}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(Resource)
	if err != nil {
		fmt.Println("Json parse error")
		return
	}
	_, _ = io.WriteString(w, string(data))
}
