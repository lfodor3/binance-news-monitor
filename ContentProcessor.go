// ContentProcessor.go
package main

import (
	"fmt"
	"github.com/go-rod/rod"
	"log"
	"strings"
	"time"
)

func findNewLinks(page *rod.Page, oldLinks []string) []string {
	links := page.MustElements("a")
	newLinks := []string{}

	for _, link := range links {
		href := link.MustAttribute("href")
		if contains(*href, "/en/support/announcement/") {
			fullLink := "https://www.binance.com" + *href
			if !containsSlice(oldLinks, fullLink) {
				newLinks = append(newLinks, fullLink)
			}
		}
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	if len(newLinks) > 0 {
		log.Printf("%v: New announcements have been found. %v new pages to be processed.", currentTime, len(newLinks))
	} else {
		log.Printf("%v: No new announcements have been found.", currentTime)
	}
	return newLinks
}

func processNewLinks(browser *rod.Browser, newLinksChannel chan string, rateLimiter *RateLimiter) {
	for newLink := range newLinksChannel {
		fmt.Println(brightYellow("New link found:"), newLink)

		rateLimiter.Wait() // Use the rate limiter to ensure a 2-second interval between opening new pages.

		// Open the new link in a new page.
		newPage := openNewPage(browser, newLink)

		// Find the title and date elements.
		title := strings.TrimSuffix(newPage.MustElement("title").MustText(), " | Binance Support")
		date := newPage.MustElement("div.css-1ebhcfx").MustText()

		// Print the title and date in the same line.
		fmt.Printf("Date: %s | Title: %s\n", date, title)

		// Get the full body of the page.
		body := newPage.MustElement("body").MustText()

		// Find the start and end indices of the relevant text.
		startIndex := strings.Index(body, title)
		endIndex := strings.Index(body, "Thanks for your support!")

		if startIndex != -1 && endIndex != -1 {
			// Slice the body text to only include the relevant content.
			announcementContent := body[startIndex+len(title) : endIndex]
			fmt.Println(announcementContent)
		} else {
			fmt.Println("Debug:")
			fmt.Printf("Title: %s\nStart index: %d\nEnd index: %d\nBody:\n%s\n", title, startIndex, endIndex, body)
			fmt.Println("=============================================================================================")
		}
		// Close the new page.
		newPage.MustClose()
	}
}

// Check if a slice contains a string.
func containsSlice(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

// Check if a string contains a substring.
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

func brightYellow(text string) string {
	return fmt.Sprintf("\033[93m%s\033[0m", text)
}
