package main

import (
	"backend/api/src/config"
	"backend/api/src/routes"
	"fmt"
)

func main() {
	config.ConnectDB()

	fmt.Println("Koneksi database berhasil diinisialisasi.")
	fmt.Println("Koneksi database berhasil diinisialisasi 2.")
	
	router := routes.RouterSetup()

	router.Run()
}