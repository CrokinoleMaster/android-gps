package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"text/template"

	. "github.com/huaruiwu/android-gps/data"
)

type position struct {
	Lat string
	Lng string
}

func (pos position) String() string {
	return fmt.Sprintf("lat: %v, lng: %v", pos.Lat, pos.Lng)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/position", positionHandler)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func positionHandler(w http.ResponseWriter, r *http.Request) {
	var pos position
	err := json.NewDecoder(r.Body).Decode(&pos)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pos)
	go sendPosition(pos)
}

func handler(w http.ResponseWriter, r *http.Request) {
	data, err := Asset("static/index.html")
	if err != nil {
		fmt.Println(err)
	}
	// t, err := template.ParseFiles("static/index.html")
	t, err := template.New("index").Parse(string(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func sendPosition(pos position) {
	conn, err := net.Dial("tcp", ":5554")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Fprintf(conn, "geo fix %v %v\n", pos.Lng, pos.Lat)
}
