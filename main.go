package main

import (
	log "github.com/Sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Println("Hello There.")
}
