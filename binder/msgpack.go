package main

import (
	"bytes"
	"fmt"
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
	"net/http"
)

func main() {
	type User struct {
		Name string
		Sex  string
	}

	b, err := msgpack.Marshal(&User{Name: "gebitang", Sex: "male"})
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Post("http://localhost:2020", "application/msgpack", bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", result)
}
