package concurrency

// placeholder for func
type WebsiteChecker func(string) bool

// check website takes func and url slice of sting
func CheckWebsite(wc WebsiteChecker, Url []string) map[string]bool {
	// created new map for checked results
	res := make(map[string]bool)

	// run func on url slice
	for _, url := range Url {
		// concurrency go-routine func
		// usually used with anonymous function
		go func(u string) {
			res[u] = wc(u)
		}(url)
	}

	return res
}
