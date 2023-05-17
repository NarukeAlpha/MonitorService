package Core

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"io"
	"log"
	"math/rand"
	"time"
)

func Task(browser playwright.BrowserContext, mw io.Writer, MangaList []DbMangaEntry, ChapterList []DbChapterEntry) {
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("https://www.readeleceed.com"); err != nil {
		log.Println("Coudln't load source website")

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

func TaskInit(mw io.Writer, mL []DbMangaEntry, cL []DbChapterEntry, pL []ProxyStruct) {
	var browser = PlaywrightInit()
	var pLng = len(pL)
	var clng = len(cL)
	if clng < pLng {

	}

	//var startingPoint = clng / pLng
	//adding a an algorithim to split the starting point of each gotask evenly across the proxies

}

func PlaywrightInit() playwright.BrowserContext {
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
