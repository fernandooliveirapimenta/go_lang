package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	fmt.Println("oi")
	res, err := http.Get("http://www.globo.com")
	fmt.Println(res)
	fmt.Println(err)
	res2, err2 := http.Get("http://www.globo.com/naoexistesjsj")
	fmt.Println(res2)
	fmt.Println(err2)

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("%s", body)
	res.Body.Close()



}
