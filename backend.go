package main

import (
    "bytes"
    "io/ioutil"
    "time"
    "net/http"
    "log"
    "github.com/aglyzov/charmap"
    "io"
)

func Decoded(body []byte) string{
	res,_ := charmap.ANY_to_UTF8(body, "KOI8-R")

	return string(res)
}

func RequestTrainData() string {
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
    return Decoded(body)

}
func RequestInAndOutData() string {
	key  := "qdj239j39jt"
    resp, err := http.Post("https://inf2086.ru/trains_timetable_api/updates",
        "", bytes.NewBuffer([]byte(key)))
    if err != nil {
        print(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        print(err)
    }
    return Decoded(body)

}

func TrainsData(resp http.ResponseWriter, req *http.Request) {
	data := RequestTrainData()
	io.WriteString(resp, data)

}

func InAndOutUpdates(resp http.ResponseWriter, req *http.Request) {
	data := RequestInAndOutData()
	io.WriteString(resp, data)
}

func main() {
	srv := &http.Server{
		Addr: ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	http.HandleFunc("/trainstimetable", TrainsData)
	http.HandleFunc("/inandoutupdates", InAndOutUpdates)

	log.Fatal(srv.ListenAndServe())

}
