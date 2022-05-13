package main

import (
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
    "time"
    "net/http"
    "log"
)

func Decoded(body string){

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

func TrainsData(resp http.ResponseWriter, req *http.Request) {
	data := RequestData()
	io.WriteString(resp, data)
	fmt.Println(resp)
}
func main() {
	srv := &http.Server{
		Addr: ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	RequestData()

	http.HandleFunc("/trainstimetable", TrainsData)

	log.Fatal(srv.ListenAndServe())

}
