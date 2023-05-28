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
		//need to add back go for concurrency when debugging is finished

		for i := 0; i < mLng; i++ {
			Task(mw, pL[i], mL, i)
		}

	} else if mLng > pLng {
		var stPointFl float64
		stPointFl = float64(mLng) / float64(pLng)
		stPoint := int(math.Round(stPointFl))

		for i := 0; i < mLng; i++ {
			stP := stPoint * i
			Task(mw, pL[i], mL, stP)
		}
	}

}

func PlaywrightInit(proxy ProxyStruct) playwright.BrowserContext {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	width := 1104
	height := 724
	viewprt := playwright.BrowserTypeLaunchPersistentContextOptionsViewport{Width: &width, Height: &height}
	//	var pth = `C:\Program Files (x86)\Microsoft\Edge\Application\msedge.exe`
	//extensionPath := "C:\\Users\\bagaa\\AppData\\Local\\Microsoft\\Edge\\User Data\\Default\\Extensions\\odfafepnkmbhccpbejgmiehpchacaeak\\1.48.0_0"
	var pwProxyStrct = playwright.BrowserTypeLaunchPersistentContextOptionsProxy{
		Server:   &proxy.ip,
		Username: &proxy.usr,
		Password: &proxy.pw,
	}
	browser, err := pw.Chromium.LaunchPersistentContext("", playwright.BrowserTypeLaunchPersistentContextOptions{
		Headless:  playwright.Bool(false),
		UserAgent: &UserAgent[rand.Intn(8)],
		Proxy:     &pwProxyStrct,
		Viewport:  &viewprt,
		//		ExecutablePath: &pth,
		ColorScheme: playwright.ColorSchemeDark,
		IgnoreDefaultArgs: []string{
			"--enable-automation",
		},
		//Args: []string{
		//	fmt.Sprintf("--disable-extensions-except=%s", extensionPath),
		//	fmt.Sprintf("--load-extension=%s", extensionPath),
		//},
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
	for {
		for i := stPoint; i < len(manga); i++ {
			cLink := ChapterLinkIncrementer(manga[i].DchapterLink, manga[i].DlastChapter)
			identifier := IdentifierDeRegex(manga[i].Didentifier)

			if _, err = page.Goto(cLink); err != nil {
				log.Printf("coudln't hit webpage chapter specific")

			}

			if manga[i].Didentifier == "" {
				log.Panicf("null value on identifier of manga entry :%v", manga[i])
			} else if manga[i].Didentifier == "Release" {
				//algorithm will check the next paged supposed to go live when the chapter is releaqsed

				page.WaitForLoadState("load")
				checkURL := page.URL()
				if checkURL == cLink {
					fmt.Fprintln(mw, "PAGE IS LIVE")
					manga[i].DlastChapter = manga[i].DlastChapter + 1
					manga[i].DchapterLink = cLink
					MangaUpdate(manga[i])
					WebhookSend(manga[i])

				} else {
					fmt.Fprintln(mw, "Page not live, will keep monitoring")
				}
			} else {
				//algorithm will check the page in the ChapterLink page for an identifier to be gone.  usually it will be some kind of countdown clock
				//This works by checking to see how many selectors of the identifier it can find.  If it can't find any, then the page is live

				countdownIdentifier, _ := page.QuerySelectorAll(identifier)
				if countdownIdentifier == nil {
					fmt.Fprintln(mw, "Failed to create slice of countdown elements")
				} else if len(countdownIdentifier) == 0 {
					fmt.Fprintln(mw, "PAGE IS LIVE")
					manga[i].DlastChapter = manga[i].DlastChapter + 1
					manga[i].DchapterLink = cLink
					MangaUpdate(manga[i])
					WebhookSend(manga[i])

				} else {
					fmt.Fprintln(mw, "Page not live, will keep monitoring")
				}

			}
		}

		time.Sleep(30 * time.Second)
	}
}
