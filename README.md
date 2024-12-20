# Concurrent Web Scraper in Go

A concurrent web scraper that extracts links from a list of URLs.

**Features:**

* **Concurrency:** Utilizes goroutines and channels to fetch and process web pages concurrently, improving efficiency.
* **Worker Pool:** Employs a worker pool pattern to control the number of concurrent requests and avoid overloading the target websites.
* **Link Extraction:** Extracts all links (`<a>` tags) from the HTML source code of the web pages.
* **Absolute URLs:** Resolves relative links to absolute URLs using the base URL of each page.
* **Synchronization:** Uses `sync.WaitGroup` to synchronize the worker goroutines and ensure proper execution order.

**How it works:**

1. **Worker Pool:** A fixed number of worker goroutines are launched.
2. **URL Distribution:**  The URLs to be scraped are sent to the workers through a channel (`urlChan`).
3. **Fetching and Parsing:** Each worker fetches a URL, parses the HTML content using `golang.org/x/net/html`, and extracts the links.
4. **Absolute URL Conversion:** Relative links are converted to absolute URLs using the base URL of the page.
5. **Result Collection:** The extracted links are sent back to the main goroutine through a channel (`results`).
6. **Output:** The main goroutine collects and prints the extracted links.

**Concurrency Mechanisms:**

* **Goroutines:**  Enable concurrent execution of the worker functions.
* **Channels:** Facilitate communication and data transfer between goroutines.
* **`sync.WaitGroup`:**  Used to wait for all worker goroutines to finish before closing the `results` channel.

**Usage:**

1. Make sure you have Go installed on your system.
2. Save the code as a `.go` file (e.g., `scraper.go`).
3. Run the program using `go run scraper.go --urls="https://example.com,https://example2.com, ...etc"`.

**Dependencies:**

* `golang.org/x/net/html`:  For HTML parsing.

**Example Usage:**

`go run fue.go --urls="https://timevko.website,https://old.reddit.com"`

**Potential Improvements:**

* **Error Handling:**  More robust error handling for network requests and HTML parsing.
* **Politeness:** Implement delays between requests to avoid overloading the target servers.
* **Data Storage:**  Store the extracted links in a file or database.
* **Deduplication:**  Remove duplicate links from the output.
* **Advanced Extraction:** Use CSS selectors (e.g., with the `goquery` library) for more specific link extraction.
