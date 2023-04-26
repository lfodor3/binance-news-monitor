// main.go
package main

import (
	"time"
)

func main() {
	browser := launchBrowser()
	defer browser.Close()

	oldLinks := []string{}
	newLinksChannel := make(chan string)

	rateLimiter := NewRateLimiter(5*time.Second, 1) // Create a rate limiter with 2-second intervals and a burst of 1.

	go processNewLinks(browser, newLinksChannel, rateLimiter)

	for {
		rateLimiter.Wait() // Use the rate limiter to ensure a 5-second interval between page loads.

		page := loadMainPage(browser)
		defer page.Close()

		newLinks := findNewLinks(page, oldLinks)

		// Send new links to the newLinksChannel to process them sequentially.
		for _, newLink := range newLinks {
			newLinksChannel <- newLink
		}

		// Update the oldLinks slice.
		oldLinks = append(oldLinks, newLinks...)

		// Wait for 5 seconds before reloading.
		time.Sleep(5 * time.Second)
	}
}
