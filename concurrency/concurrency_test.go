package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url == "http://google.com"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://yahoo.com",
	}

	want := map[string]bool{
		"http://google.com":   true,
		"http://facebook.com": false,
		"http://yahoo.com":    false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
