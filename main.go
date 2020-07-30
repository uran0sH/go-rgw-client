package main

import (
	"fmt"
	"strconv"
)

type Object struct {
	Content []byte
	Name    string
	Bucket  string
}
const ip = "http://192.168.101.8:8080"
func main() {
	//CreateBucket("test")
	//obj, err := NewObject("./hello", "test", "hello")
	//if err != nil {
	//	fmt.Print(err)
	//}
	//putObject(*obj)
	//err := registerUser("hwy", "123456")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//_, err := login("hwy", "123456")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//test("abc")
	uploadID := createMultipart("test", "hello")
	var etags [10]string
	for i := 0; i < 10; i++ {
		obj, err := NewObject("./hello", "test", "hello")
		if err != nil {
			fmt.Print(err)
		}
		etag := uploadPart(uploadID, strconv.Itoa(i), *obj)
		etags[i] = etag
	}
	var multipart CompleteMultipart
	for i := 0; i < 5; i++ {
		partID := strconv.Itoa(i)
		etag := etags[i]
		part := Part{
			PartID: partID,
			ETag: etag,
		}
		multipart.Parts = append(multipart.Parts, part)
	}
	complete(uploadID, "test", "hello", multipart)
	//abort(uploadID, "test", "hello")
}
