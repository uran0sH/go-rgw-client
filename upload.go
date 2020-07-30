package main

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func readFromFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var content []byte
	var cache = make([]byte, 1024)
	for {
		n, err := file.Read(cache)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		content = append(content, cache[:n]...)
	}
	return content, nil
}

func putObject(object Object) {
	client := http.Client{}
	body := bytes.NewReader(object.Content)
	request, err := http.NewRequest("POST", ip + "/upload/" + object.Bucket + "/" + object.Name, body)
	if err != nil {
		fmt.Println(err)
	}
	checkSum := md5.New()
	checkSum.Write(object.Content)
	hash := base64.StdEncoding.EncodeToString(checkSum.Sum(nil))
	fmt.Println(hash)
	request.Header.Add("Content-MD5", hash)
	request.Header.Add("c-meta-hello", "hello meta")
	rep, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	res, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", string(res))
	_ = rep.Body.Close()
}

func NewObject(filename string, bucket string, name string) (*Object, error) {
	data, err := readFromFile(filename)
	if err != nil {
		return nil, err
	}
	obj := &Object{
		Content: data,
		Bucket:  bucket,
		Name:    name,
	}
	return obj, nil
}
