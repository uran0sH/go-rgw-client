package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func registerUser(username, password string) error{
	user := User{Username: username, Password: password}
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	client := http.Client{}
	body := bytes.NewReader(data)
	request, err := http.NewRequest("POST", ip + "/register", body)
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")
	rep, err := client.Do(request)
	if err != nil {
		return err
	}
	res, _ := ioutil.ReadAll(rep.Body)
	fmt.Printf("%s", res)
	_ = rep.Body.Close()
	return nil
}

type Result struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data string `json:"data"`
}

func login(username, password string) (string, error) {
	user := User{Username: username, Password: password}
	data, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	client := http.Client{}
	body := bytes.NewReader(data)
	request, err := http.NewRequest("POST", ip + "/login", body)
	if err != nil {
		return "", err
	}
	rep, err := client.Do(request)
	if err != nil {
		return "", err
	}
	res, _ := ioutil.ReadAll(rep.Body)
	var result Result
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}
	return result.Data, nil
}


func test(tokenString string) {
	request, err := http.NewRequest("GET", ip + "/test", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := http.Client{}
	request.Header.Add("Authorization", "Bearer " + tokenString)
	rep, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, _ := ioutil.ReadAll(rep.Body)
	fmt.Printf("%s", res)
}
