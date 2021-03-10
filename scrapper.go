package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)


func GetUserCode(Url string)(string,error) {
	// Make HTTP request
	response, err := http.Get(Url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
     	var code string
	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return code,fmt.Errorf("Error loading HTTP response body. ", err)
		
	}

	// Find all pre tag and process them with the function
	// defined earlier
	document.Find("pre").Each(func(index int, element *goquery.Selection) {
		id, _ := element.Attr("id")
		if id=="program-source-text"{
           code= element.Text()
		}
	})
	return code,nil
}
