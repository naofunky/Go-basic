package main

import (
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Create("logger/logger.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	flags := log.Lshortfile | log.LstdFlags

	warnLogger := log.New(io.MultiWriter(file, os.Stderr), "WARN: ", flags)
	errorLogger := log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", flags)
	warnLogger.Println("warning A")
	errorLogger.Fatalln("enforcement error")
}
