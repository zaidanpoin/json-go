package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Online    bool   `json:"online"`
}
type barang struct {
	Nama         string `json:"nama"`
	Harga        int    `json:"harga"`
	Jenis_barang string `json:"jenis_barang"`
}

func main() {
	// a := "http://www.belajar-golang.com/basic?nama=zaidaan"
	// e, _ := url.Parse(a)
	// fmt.Println(e.Scheme)
	// fmt.Println(e.Host)
	// fmt.Println(e.Path)
	// t := e.Query()["nama"][0]
	// fmt.Println(t)

	decodeJsonToObject()

}

func decodeJsonToObject() {

	jsonString := `[
    {"nama": "abraham", "harga": 9000, "jenis_barang": "Makanan"},
    {"nama": "qila", "harga": 9000, "jenis_barang": "Makanan"},
    {"nama": "cil", "harga": 9000, "jenis_barang": "Makanan"}
]`
	jsonData := []byte(jsonString)

	var data []barang
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		panic(err)
	}

	for _, a := range data {
		fmt.Println(a.Nama)

	}
}

func encodeObjectTojson() {

}
