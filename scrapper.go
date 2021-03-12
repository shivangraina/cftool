package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

//GetUserCode ..
func GetUserCode(SubData FetchCode, CodeReciever chan UserCodeData, wg *sync.WaitGroup, m *sync.Mutex, client1 *http.Client) (string, error) {

	// Make HTTP request
	defer wg.Done()
	m.Lock()
	response, err := client1.Get(SubData.SubmissionURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK {
		log.Println(response)
		log.Fatal()

	}

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		fmt.Errorf("Error loading HTTP response body. ", err)

	}

	code := document.Find("#program-source-text").First().Text()
	if code == "" {
		time.Sleep(5 * time.Second)
		fmt.Println("SLEEPING")
	}

	fmt.Println(code, SubData.SubmissionURL)

	//bodystring := string(body)
	//fmt.Println(bodystring,"BOOOOOOOTY",response.StatusCode)
	var DataWithCode UserCodeData
	DataWithCode.ContestID = SubData.ContestID
	DataWithCode.ProblemIndex = SubData.ProblemIndex
	DataWithCode.ProblemCode = code
	DataWithCode.Language = SubData.Language
	m.Unlock()
	CodeReciever <- DataWithCode

	return code, nil
}
