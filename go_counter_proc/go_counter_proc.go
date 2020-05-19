package go_counter_proc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

// type for storing the url with related data
type siteStat struct {
	url   string
	numGo int // the number of "Go" in the body of response from url
}

// countGo counts the number of "Go" in the body of response from url
func countGo(url string) (siteStat, error) {
	// send GET request
	resp, err := http.Get(url)
	if err != nil {
		return siteStat{url, 0}, fmt.Errorf("request to %s failed, %s", url, err)
	}
	defer resp.Body.Close()
	// check the status of the response
	if resp.StatusCode != http.StatusOK {
		return siteStat{url, 0}, fmt.Errorf("request to %s failed, %s", url, resp.Status)
	}
	// check content type
	contentType := resp.Header.Get("Content-Type")
	if contentType != "text/html" && !strings.HasPrefix(contentType, "text/html;") {
		return siteStat{url, 0}, fmt.Errorf("the type of doc at %s is not text/html but %s", url, contentType)
	}
	// put the body of response into buffer
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("read from body of response from %s error %s", url, err)
	}
	// count the number of "Go" instanses in the body of response
	numGo := bytes.Count(resp_body, []byte("Go"))
	return siteStat{url, numGo}, nil
}

// reads from stdin urls,
// counts and prints the number of "Go" in the body of every entered url,
// returns the total number of "Go" in all urls
// input parameter k is the number of simultaneousely fetched urls
//func InputCountGo(inpSrc *os.File, k int) int {
func InputCountGo(inpSrc io.Reader, k int) int {
	if inpSrc == nil {
		inpSrc = io.Reader(os.Stdin)
	}
	if k == 0 {
		k = 5
	}
	urlsCh := make(chan string)      // channel with urls to handle
	numGoList := make(chan int)      // channel to pass number of Go from fetcher to total counter
	inpListRdy := make(chan bool)    // signal of input list is sent to handling
	tokens := make(chan struct{}, k) // tokens for limiting simultaneously running goroutines
	var wg sync.WaitGroup            // waitgroup for waiting for running gouroutines are finished

	// start fetching goroutines with limiting of number of goroutines
	go func() {
		for link := range urlsCh {
			go func(link string) {
				tokens <- struct{}{} // wait for finishing of started more than limit goroutines
				siteData, err := countGo(link)
				if err != nil {
					fmt.Errorf("with url %s error %v", link, err)
				}
				fmt.Printf("Count for %s: %d\n", siteData.url, siteData.numGo)
				numGoList <- siteData.numGo
				<-tokens
			}(link)
		}
	}()

	// read sites statistics and accumulate total nomber fo "Go"
	tot_num := 0
	go func() {
		for num := range numGoList {
			tot_num += num
			wg.Done()
		}
	}()

	// fill input channel with input values
	go func() {
		scanner := bufio.NewScanner(inpSrc)
		for scanner.Scan() {
			wg.Add(1)
			addText := scanner.Text()
			urlsCh <- addText
		}
		inpListRdy <- true
	}()
	// wait for filling input channel before starting of waiting for output channel is empty
	<-inpListRdy
	// wait for all sites was counted
	wg.Wait()
	return tot_num
}
