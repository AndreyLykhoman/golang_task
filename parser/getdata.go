package main
//TODo: rename file and write secon servis that can to control to work with db and json
//ToDo: rebase to main folder getdata

import (
	"fmt"
	"net/http"
	"crypto/tls"
	"io/ioutil"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	result, err := client.Get("https://habrahabr.ru/post/150732/")
	if err != nil {
		fmt.Println(err)
	}

	defer result.Body.Close()
	body, err := ioutil.ReadAll(result.Body)
	if result != nil{
		s := string(body)
		fmt.Println(s)
	}


}
