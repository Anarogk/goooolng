package concurrency

import (
	"reflect"
	"testing"
	"time" //using for benchmarking
)

func slowStubWebCheck(_ string) bool {
	time.Sleep(100 * time.Microsecond)
	return true
}

func mockChecker(Url string) bool {
	if Url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func CheckWebsiteTest(t *testing.T) {

	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	t.Run("check_site", func(t *testing.T) {
		got := CheckWebsite(mockChecker, websites)

		want := map[string]bool{
			"http://google.com":          true,
			"http://blog.gypsydave5.com": true,
			"waat://furhurterwe.geds":    true,
		}

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}

	})
}

// run benchmark for slow-checker func that uses time.sleep() to deliberately be slow
func BenchmarkSlowStubWebCheck(b *testing.B) {
	// make a slice of strings that are initialized empty
	url := make([]string, 100)

	// run for loop for length of url
	for i := range url {
		// initialize string to random string
		url[i] = "a string"
	}
	// reset timer
	b.ResetTimer()
	// run benchmark for b.N times
	for i := 0; i < b.N; i++ {
		CheckWebsite(slowStubWebCheck, url)
	}

}
