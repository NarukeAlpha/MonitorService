package main

import (
	"eleceedMonitor/Core"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/playwright-community/playwright-go"
	"io"
	"log"
	"os"
	"os/signal"
	"sync"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load .env file")
	}
	wbKey := fmt.Sprintf(os.Getenv("webKey"))
	err = playwright.Install()

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

	//waitgroup to launch all 2 go routines and wait until each one is done before attempting to reach from each channel.
	wg.Add(2)
	go Core.MangaSync(mChannel, &wg)
	go Core.ProxyLoad(pChannel, &wg)
	//	wg.Wait()

	//receiving from each channel and closing them
	mL := <-mChannel
	pL := <-pChannel
	close(mChannel)
	close(pChannel)

	log.Printf("Starting monitor")
	//initializing the monitor
	Core.TaskInit(mw, mL, pL, wbKey)

	//for {
	//	//infinite loop to keep the program running
	//	//might add an open server to interaqct directly with the program?
	//	//This for loop is never hit at the moment since there is no concurrency running
	//	//but it is here for future use when go Core.TaskInit is used, then we can do something, probably open a server with gorilla mux for health/monitor/resync
	//

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

}
