package Core

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"io"
	"log"
	"time"
)

func Task(browser playwright.BrowserContext, mw io.Writer) {
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
