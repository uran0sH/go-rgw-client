package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateBucket(bucket string) {
	client := http.Client{}
	request, err := http.NewRequest("GET", ip + "/createbucket/" + bucket, nil)
	if err != nil {
		fmt.Println(err)
	}
	rep, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	res, _ := ioutil.ReadAll(rep.Body)
	fmt.Printf("%v", string(res))
	_ = rep.Body.Close()
}
