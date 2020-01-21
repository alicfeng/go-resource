package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/alicfeng/go-resource/bean"
)

func main() {
	http.HandleFunc("/", resourceHandler)
	e := http.ListenAndServe(":8888", nil)
	if e != nil {
		fmt.Println(e.Error())
	}

}

/**
资源路由事件业务处理
*/
func resourceHandler(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(bean.Resource)
	if err != nil {
		fmt.Println("Json parse error")
		return
	}
	_, _ = io.WriteString(w, string(data))
}
