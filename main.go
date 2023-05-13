package main

import (
	"eleceedMonitor/Core"
	"io"
	"log"
	"os"
)

func AssertErrorToNil(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func main() {

	file, err := os.Open("QuerySelector.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	mw := io.MultiWriter(os.Stdout, file)

	var browser = Core.PlaywrightInit()
	var mL, cL = Core.ServerSync()
	Core.Task(browser, mw)

}
