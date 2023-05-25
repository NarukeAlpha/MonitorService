package main

import (
	"eleceedMonitor/Core"
	"io"
	"log"
	"os"
	"sync"
)

func AssertErrorToNil(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func main() {
	//making the channels for the go routines to communicate and reduce execution time before monitor starts
	var wg sync.WaitGroup
	mChannel := make(chan []Core.DbMangaEntry)
	pChannel := make(chan []Core.ProxyStruct)
	//opening log file and creating a multiwriter to write to both stdout and file
	file, err := os.Open("QuerySelector.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	mw := io.MultiWriter(os.Stdout, file)

	//waitgroup to launch all 3 go routines and wait until each one is done before attempting to reach from each channel.
	wg.Add(3)
	go Core.MangaSync(mChannel, &wg)
	go Core.ProxyLoad(pChannel, &wg)
	wg.Wait()

	//receiving from each channel and closing them
	mL := <-mChannel
	pL := <-pChannel
	close(mChannel)
	close(pChannel)

	//initializing the monitor
	Core.TaskInit(mw, mL, pL)

}
