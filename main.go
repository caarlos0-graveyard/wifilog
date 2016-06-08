package main

import (
	"log"
	"os"

	"github.com/caarlos0/wifilog/wifi"
)

func openLogFile() *os.File {
	file, err := os.OpenFile("connections.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	return file
}

func main() {
	file := openLogFile()
	defer file.Close()
	log.SetOutput(file)

	ssid, err := wifi.New().SSID()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to", ssid)
}
