package go_counter_proc

import (
	"fmt"
	"io"
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

// testReader for emulating substitution instead of Stdin
type testReader []string

// implementing Read method for reading from the list of strings
// into buffer with terminating every string with \n
func (obj *testReader) Read(p []byte) (n int, err error) {
	if len(*obj) > 0 {
		// get first string and add \n for splitting by Scanner
		retStr := (*obj)[0] + "\n"
		// copy string into buffer
		retStrLen := copy(p[:], retStr)
		// remove first string from the list
		*obj = (*obj)[1:]
		return retStrLen, nil
	} else {
		// if the list of strings is empty then return end of file
		return 0, io.EOF
	}
}

// test InputCountGo function
// with list of urls from object with io.Reader interface
func TestGoCounter_list(t *testing.T) {
	var inp_list testReader
	inp_list = append(inp_list, "https://golang.org")
	inp_list = append(inp_list, "https://x.x")
	inp_list = append(inp_list, "https://google.com")

	countGo := InputCountGo(&inp_list, 2)
	if countGo != 26 {
		t.Errorf("in a list of sites counted %d", countGo)
	}
}
