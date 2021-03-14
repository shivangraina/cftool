package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

//baseURL ..
var baseURL string = "https://codeforces.com"

//GetUserSubmissions .. Fetches Submission in batches of count
func GetUserSubmissions(userHandle string, from, count int) ([]Submission, error) {
	var submissions []Submission
	url := fmt.Sprintf(baseURL+"/api/user.status?handle=%s&from=%s&count=%s", userHandle, fmt.Sprint(from), fmt.Sprint(count))
	fmt.Println(url)
	var wrapper Results
	FetchJSON(url, &wrapper)
	if len(wrapper.Result) == 0 {
		return submissions, fmt.Errorf("Empty List")

	}
	for _, submission := range wrapper.Result {
		subURL := fmt.Sprintf(baseURL+"/contest/%s/submission/%s", fmt.Sprint(submission.ContestID), fmt.Sprint(submission.ID))
		sub := Submission{
			ContestID:     submission.ContestID,
			ProblemIndex:  submission.Problem.Index,
			ProblemName:   submission.Problem.Name,
			Verdict:       submission.Verdict,
			Language:      submission.ProgrammingLang,
			ProblemRating: submission.Problem.Rating,
			Tags:          submission.Problem.Tags,
			SubmissionID:  submission.ID,
			SubmissionURL: subURL,
		}
		submissions = append(submissions, sub)
	}
	return submissions, nil

}

//UserCodeData ..recieving channel from cf channel type to send to github channel
type UserCodeData struct {
	ContestID    int64
	ProblemIndex string
	ProblemCode  string
	Language     string
}

//FetchCode .. for sending to scrapper channel
type FetchCode struct {
	ContestID     int64
	ProblemIndex  string
	SubmissionURL string
	Language      string
}

func main() {
	PageIndex := 1
	length := 100
	var CfUsername, reponame, owner string

	flag.StringVar(&CfUsername, "h", "", "provide your cf handle")
	flag.StringVar(&owner, "g", "", "provide your github username")
	flag.StringVar(&reponame, "n", "", "provide the name of repo to be created")

	flag.Parse()
	

	CodeReciever := make(chan UserCodeData, length)

	var wg sync.WaitGroup
	for {

		submission, err := GetUserSubmissions(CfUsername, PageIndex*length+1, length)

		if err != nil {
			break
		}
		for _, sub := range submission {
			if sub.Verdict != "OK" {
				continue
			}
			var data FetchCode
			data.ContestID = sub.ContestID
			data.ProblemIndex = sub.ProblemIndex
			data.SubmissionURL = sub.SubmissionURL
			data.Language = sub.Language

			wg.Add(1)
			go GetUserCode(data, CodeReciever, &wg)
		}
		PageIndex++
	}
	// close the channel in the background
	go func() {
		wg.Wait()
		close(CodeReciever)
	}()
	client := GetClientWithToken()

	CreateEmptyRepositry(client, reponame)
	for code := range CodeReciever {
	

			CreateContestFiles(client, code, reponame, owner)

		
	}

}
func httpGet(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {

		return []byte{}, fmt.Errorf("could not do a get,statuscode %d", res.StatusCode)

	}

	raw, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return []byte{}, err
	}

	return raw, nil
}

//FetchJSON ..
func FetchJSON(url string, wrapper interface{}) error {
	raw, err := httpGet(url)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(bytes.NewReader(raw))

	dec.Decode(&wrapper)
	return nil

}
