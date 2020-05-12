package go_counter_proc

import (
	"fmt"
	"testing"
)

func TestGoCounter_golang(t *testing.T) {
	url := "https://golang.org"
	siteData, err := countGo(url)
	fmt.Printf("test with url %s returns %d err %v\n", url, siteData.numGo, err)
	if !(siteData.numGo == 20) {
		t.Errorf("test with %s returns %v err %v", url, siteData, err)
	}
}

func TestGoCounter_google(t *testing.T) {
	url := "https://google.org"
	siteData, err := countGo(url)
	fmt.Printf("test with url %s returns %d err %v\n", url, siteData.numGo, err)
	if !(siteData.numGo == 30) {
		t.Errorf("test with %s returns %v err %v", url, siteData, err)
	}
}
