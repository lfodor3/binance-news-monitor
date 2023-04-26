// HeadlessBrowser.go
package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func launchBrowser() *rod.Browser {
	url := launcher.New().Bin("chromium").Headless(true).MustLaunch()
	return rod.New().ControlURL(url).MustConnect()
}

func loadMainPage(browser *rod.Browser) *rod.Page {
	page := browser.MustPage("https://www.binance.com/en/support/announcement/new-cryptocurrency-listing?c=48&navId=48")
	page.MustWaitLoad()
	return page
}

func openNewPage(browser *rod.Browser, url string) *rod.Page {
	page := browser.MustPage(url)
	page.MustWaitLoad()
	return page
}
