package main

import "encoding/json"
import "net/http"
import "fmt"
import "bytes"
import "net/url"

var baseURL = "http://localhost:8080"

//struct sebagai penampung data hasil query
type data_karyawan struct {
	ID      int
	Nama    string
	Umur    int
	Jabatan string
}

//fungsi mengambil data dari api
func ambil_api(karyawan string) (data_karyawan, error) {
	var err error
	var client = &http.Client{}
	var data data_karyawan

	var param = url.Values{}
	param.Set("Nama", karyawan)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/cari_karyawan", payload)

	if err != nil {
		return data, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)

	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

//fungsi main
func main() {
	var karyawan, err = ambil_api("jooprima")
	if err != nil {
		fmt.Println("karyawan tidak tersedia", err.Error())
		return
	}
	fmt.Println("ID : ", karyawan.ID, "Nama : ", karyawan.Nama, "Umur : ", karyawan.Umur, "Jabatan : ", karyawan.Jabatan)
}
