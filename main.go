package main

import (
	"context"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
)

func main() {
	simple_first()
	simple_second()
}

func simple_first() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://golang.org/pkg/time/`),
		chromedp.Text(`#pkg-overview`, &res, chromedp.NodeVisible, chromedp.ByID),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(strings.TrimSpace(res))
}

func simple_second() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx, submit(`https://github.com/search`, `//input[@name="q"]`, `chromedp`, &res))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("got: `%s`", strings.TrimSpace(res))
}

func submit(urlstr, sel, q string, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(sel),
		chromedp.SendKeys(sel, q),
		chromedp.Submit(sel),
		chromedp.WaitNotPresent(`//*[@id="js-pjax-container"]//h2[contains(., 'Search more than')]`),
		chromedp.Text(`(//*[@id="js-pjax-container"]//ul[contains(@class, "repo-list")]/li[1]//p)[1]`, res),
	}
}