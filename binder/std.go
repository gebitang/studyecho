package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// total of maxMemory bytes
		err := r.ParseMultipartForm(32 << 20)
		if err == nil {
			panic(err)
		}
		data := map[string]interface{}{
			"form":      r.Form,
			"post_form": r.PostForm,
		}

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data["json_data"] = string(reqBody)

		fmt.Fprintln(w, data)
	})
	log.Fatal(http.ListenAndServe(":2020", nil))
}
