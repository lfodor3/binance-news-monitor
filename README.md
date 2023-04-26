# Binance News Monitor

This project is a simple web scraper that periodically loads the Binance
announcements main page and processes new links it finds on the page. The
scraper is implemented in Go and uses the go-rod package to interact with the
headless browser.

## Folder Structure

your_project/
│
├── main.go
├── ContentProcessor.go
├── HeadlessBrowser.go
├── RateLimiter.go
└── README.md

## Architecture Diagram

+-------------+     +------------------+     +----------------+     +-------------+
|   main.go   | --> | ContentProcessor | <-- | HeadlessBrowser | <-- | RateLimiter |
+-------------+     +------------------+     +----------------+     +-------------+

### main.go
This file is the entry point of the program. It initializes the browser, rate
limiter, and the new links channel, then enters an infinite loop to load the
main page and find new links.

### ContentProcessor.go
This file contains functions for processing the content found on the main page
of the website. It includes `findNewLinks`, which finds new announcement links,
and `processNewLinks`, which processes the new links as they are received on the
new links channel.

### HeadlessBrowser.go
This file contains functions related to managing the headless browser using the
go-rod package. It includes `launchBrowser`, which launches a new headless
Chromium browser, `loadMainPage`, which loads the main Binance announcement
page, and `openNewPage`, which opens a new page with the provided URL.

### RateLimiter.go
This file contains the implementation of a custom rate limiter. The rate limiter
is used to control the rate at which certain operations are performed, such as
requests to a server or web scraping. It includes the `RateLimiter` struct and
methods like `NewRateLimiter` and `Wait`.
