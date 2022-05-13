package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http" 
)

func Decoded(body string) string {

}

func RequestData() string {
	key  := "qdj239j39jt"
    resp, err := http.Post("https://inf2086.ru/trains_timetable_api/timetable",
        "", bytes.NewBuffer([]byte(key)))
    if err != nil {
        print(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        print(err)
    }
    return string(body)

    }


func main() {
	var srv http.Server
	srv.string = ":https"
	app := http.NewServeMux()
	RequestData()

}
