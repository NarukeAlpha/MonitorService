package Core

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"io"
	"log"
	"math"
	"math/rand"
	"time"
)

func TaskInit(mw io.Writer, mL []DbMangaEntry, pL []ProxyStruct) {
	//var browser = PlaywrightInit()
	var pLng = len(pL)
	var mLng = len(mL)

	if mLng <= pLng {
		//launch one task per proxy up to the amount of task available

		for i := 0; i < mLng; i++ {
			go Task(mw, pL[i], mL, i)
		}

	} else if mLng > pLng {
		var stPointFl float64
		stPointFl = float64(mLng) / float64(pLng)
		stPoint := int(math.Round(stPointFl))

		for i := 0; i < mLng; i++ {
			stP := stPoint * i
			go Task(mw, pL[i], mL, stP)
		}
	}

	//var startingPoint = mLng / pLng
	//adding a an algorithim to split the starting point of each gotask evenly across the proxies

}

func PlaywrightInit(proxy ProxyStruct) playwright.BrowserContext {
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
		UserAgent:      &UserAgent[rand.Intn(8)],
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
	return browser
}

func Task(mw io.Writer, proxy ProxyStruct, manga []DbMangaEntry, stPoint int) {
	browser := PlaywrightInit(proxy)
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	defer page.Close()
	for i := stPoint; i < len(manga); i++ {
		if manga[i].Didentifier == "" {
			log.Panicf("null value on identifier of manga entry :%v", manga[i])
		} else if manga[i].Didentifier == "Release" {
			//algorithm will check the next paged supposed to go live when the chapter is releaqsed
		} else {
			//algorithm will check the page in the ChapterLink page for an identifier to be gone.  usually it will be some kind of countdown clock
		}
	}

	if _, err = page.Goto("https://readeleceed.com/manga/eleceed-chapter-239/"); err != nil {
		log.Printf("coudln't hit webpage chapter specific")
	}

	time.Sleep(5 * time.Second)
	iframes, _ := page.QuerySelectorAll("iframe")
	for _, iframe := range iframes {
		iframe.Evaluate("this.remove()")
	}

	countdownEls, _ := page.QuerySelectorAll("[data-type='countdown']")
	if countdownEls == nil {
		fmt.Fprintln(mw, "failed to create element counter")
	} else if len(countdownEls) == 0 {
		fmt.Fprintln(mw, "PAGE IS LIVE")
		WebhookSend()
	} else {
		fmt.Fprintln(mw, "Page not live, will keep monitoring")

	}
	time.Sleep(10 * time.Second)
}
