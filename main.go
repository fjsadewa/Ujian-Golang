package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type DataUser struct{
	Nama, Nim, Alamat string
}

var data []DataUser

func selamatDatang(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "Selamat Datang di Server Lokal")
}
func inputData(w http.ResponseWriter, r *http.Request)  {
	err := r.ParseMultipartForm(0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	nama := r.FormValue("nama")
	nim := r.FormValue("nim")
	alamat := r.FormValue("alamat")

	if r.FormValue("nama") == "" || r.FormValue("nim") == ""||r.FormValue("alamat") == ""{
		http.Error(w,"Data yang dikirim kosong",http.StatusBadRequest)
		return
	}

	newDataUser := DataUser{
		Nama : nama, 
		Nim : nim, 
		Alamat : alamat,
	}

	data = append(data, newDataUser)
	
	str,err := json.Marshal(data)
	if err != nil{
		io.WriteString(w, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(str))
	
}
func hasilData(w http.ResponseWriter, r *http.Request)  {
	str,err := json.Marshal(data)
	if err != nil{
		io.WriteString(w, err.Error())
	}
	
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(str))
	
}

func main()  {
	multiplexer := http.NewServeMux()
	multiplexer.HandleFunc("/home",selamatDatang)
	multiplexer.HandleFunc("/nama",inputData)
	multiplexer.HandleFunc("/semuadata",hasilData)

	err := http.ListenAndServe(":3030",multiplexer)
	if err != nil{
		panic(err)
	}
}