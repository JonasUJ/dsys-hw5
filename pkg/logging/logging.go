package logging

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	Logger = log.New(writer(), fmt.Sprintf("%s | ", hostname()), log.Ltime)
)

func writer() io.Writer {
	logfile := fmt.Sprintf("/app/logs/%s.log", hostname())
	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("fail to open log file: %v", err)
	}

	writer := io.MultiWriter(os.Stdout, file)

	return writer
}

func hostname() string {
	name, err := os.Hostname()
	if err != nil {
		log.Fatalf("could not get hostname: %v", err)
	}

	return name
}
