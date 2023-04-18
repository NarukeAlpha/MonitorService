package main

import (
	"eleceedMonitor/Core"
	"fmt"
	"github.com/playwright-community/playwright-go"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
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
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	width := 1104
	height := 724
	viewprt := playwright.BrowserTypeLaunchPersistentContextOptionsViewport{Width: &width, Height: &height}
	var pth = `C:\Program Files (x86)\Microsoft\Edge\Application\msedge.exe`
	extensionPath := "C:\\Users\\bagaa\\AppData\\Local\\Microsoft\\Edge\\User Data\\Default\\Extensions\\odfafepnkmbhccpbejgmiehpchacaeak\\1.48.0_0"
	browser, err := pw.Chromium.LaunchPersistentContext("", playwright.BrowserTypeLaunchPersistentContextOptions{
		Headless:       playwright.Bool(false),
		UserAgent:      &Core.UserAgent[rand.Intn(8)],
		Viewport:       &viewprt,
		ExecutablePath: &pth,
		ColorScheme:    playwright.ColorSchemeDark,
		IgnoreDefaultArgs: []string{
			"--enable-automation",
		},
		Args: []string{
			fmt.Sprintf("--disable-extensions-except=%s", extensionPath),
			fmt.Sprintf("--load-extension=%s", extensionPath),
		},
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	time.Sleep(10 * time.Second)

	if err = browser.Close(); err != nil {
		log.Fatalf("failed to end program?")
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("memory leak wtf?")
	}

	Core.Task(browser, mw)

}
