package main

import (
	"flag"
	"log"
	
	"go-redis/config"
    "go-redis/server"
)

func setupFlags() {
	flag.StringVar(&config.Host, "Host", "0.0.0.0", "host for the go redis server")
	flag.IntVar(&config.Port, "Port", 7379, "Port for the go redis server")
	flag.Parse()
}

func main() {
	setupFlags()
	log.Println("go redis!!")
    server.RunSyncTCPServer() 
}