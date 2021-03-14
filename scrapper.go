package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

//GetUserCode ..
func GetUserCode(SubData FetchCode, CodeReciever chan UserCodeData, wg *sync.WaitGroup) (string, error) {

	// Make HTTP request
	defer wg.Done()

	response, err := http.Get(SubData.SubmissionURL)
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
	fmt.Println("Fetching Code for problem ", SubData.ContestID, SubData.ProblemIndex)
	code := document.Find("#program-source-text").First().Text()

	var DataWithCode UserCodeData
	DataWithCode.ContestID = SubData.ContestID
	DataWithCode.ProblemIndex = SubData.ProblemIndex
	DataWithCode.ProblemCode = code
	DataWithCode.Language = SubData.Language

	CodeReciever <- DataWithCode

	return code, nil
}
