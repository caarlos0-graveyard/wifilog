package main

import (
	"log"
	"os"

	"github.com/caarlos0/wifilog/wifi"
)

func main() {
	f, err := os.OpenFile("connections.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	log.SetOutput(f)
	ssid, err := wifi.New().SSID()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to", ssid)
}
