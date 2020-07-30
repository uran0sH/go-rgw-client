package main

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type UploadID struct {
	ID string `json:"uploadID"`
}
func createMultipart(bucket, object string) string {
	client := http.Client{}
	request, err := http.NewRequest("POST", ip + "/uploads/create/" + bucket + "/" + object, nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	rep, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	body, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var uploadID UploadID
	err = json.Unmarshal(body, &uploadID)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return uploadID.ID
}

func uploadPart(uploadID, partID string, object Object) string {
	client := http.Client{}
	body := bytes.NewReader(object.Content)
	apiUrl := ip + "/uploads/upload/" + object.Bucket + "/" + object.Name
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	data := url.Values{}
	data.Set("PartNumber", partID)
	data.Set("UploadId", uploadID)
	u.RawQuery = data.Encode()
	request, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		fmt.Println(err)
		return ""
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
		return ""
	}
	//res, err := ioutil.ReadAll(rep.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return ""
	//}
	fmt.Printf("%v", rep.StatusCode)
	//_ = rep.Body.Close()
	return rep.Header.Get("ETag")
}

type Part struct {
	PartID string `json:"PartID"`
	ETag   string `json:"ETag"`
}

type CompleteMultipart struct {
	Parts []Part
}

func complete(uploadID, bucket, object string, multipart CompleteMultipart) {
	data, err := json.Marshal(&multipart)
	if err != nil {
		fmt.Println(err)
	}
	body := bytes.NewReader(data)
	client := http.Client{}
	apiUrl := ip + "/uploads/complete/" + bucket + "/" + object
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Println(err)
	}
	v := url.Values{}
	v.Set("UploadId", uploadID)
	u.RawQuery = v.Encode()
	request, err := http.NewRequest("POST", u.String(), body)
	rep, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v", rep.StatusCode)
}

func abort(uploadID, bucket, object string) {
	client := http.Client{}
	request, err := http.NewRequest("POST", ip + "/uploads/abort/" + bucket + "/" + object, nil)
	rep, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v", rep.StatusCode)
}