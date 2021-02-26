package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)
//BaseURL ..
var BaseURL string = "https://codeforces.com/api" 

//Results ..
type Results struct {
	Result []Result `json:"result"`
}
//Result ..
type Result struct {
	ID              int64   `json:"id"`
	ContestID       int64   `json:"contestId"`
	Problem         Problem `json:"problem"`
	ProgrammingLang string  `json:"programmingLanguage"`
	Verdict         string  `json:"verdict"`
	Author          Author  `json:"author"`
}
// Problem ..
type Problem struct {
	Index  string   `json:"index"`
	Name   string   `json:"name"`
	Tags   []string `json:"tags"`
	Rating int64    `json:"rating"`
}
//Author ..
type Author struct {
	ParticipantType string `json:"participantType"`
}

func main() {

	fmt.Println(PlayedContests())

}
//FetchContestData ..
func FetchContestData(jobs <-chan int64, results chan<- Results) {
	var Handle string = "RemoteCodeExecution" //take from cmd
	var wrapper Results

	for contestid := range jobs {

		url := fmt.Sprintf(BaseURL+"/contest.status?contestId=%d&handle=%s&from=1&count=50", contestid, Handle)
		fmt.Println(url)

		FetchJSON(url, &wrapper)
		results <- wrapper

	}

}
//GetAllContests ..
func GetAllContests() []int64 {
	var wrapper Results
	var contestids []int64
	url := fmt.Sprintf(BaseURL + "/contest.list")
	

	FetchJSON(url, &wrapper)
	for _, result := range wrapper.Result {
		contestids = append(contestids, result.ID)

	}
	return contestids

}
//PlayedContests ..
func PlayedContests() []int64 {

	contestids := GetAllContests()

	var played []int64
	jobs := make(chan int64, 100)
	results := make(chan Results, 100)
	for i := 0; i <= 1000; i++ {
		go FetchContestData(jobs, results)
	}
	for _, contestid := range contestids {
		jobs <- contestid

	}
	close(jobs)
	for _ = range jobs {
		wrapper := <-results
		for _, submission := range wrapper.Result {
			fmt.Println(submission)
			if submission.Author.ParticipantType == "CONTESTANT" {
				played = append(played, submission.ContestID)
				break

			}
		}

	}
	return played
}
//FetchJSON ..
func FetchJSON(url string, wrapper interface{}) error {
	res, err := http.Get(url)
	if err != nil {
        fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)
	return dec.Decode(&wrapper)

}
