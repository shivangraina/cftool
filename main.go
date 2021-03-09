package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

//BaseURL ..
var BaseURL string = "https://codeforces.com"

// create an Http Client
var client http.Client

//GetUserSubmissions .. Fetches Submission in batches of count
func GetUserSubmissions(userhandle string, from, count int) ([]Submission, error) {
	var submissions []Submission
	url := fmt.Sprintf(BaseURL+"/api/user.status?handle=%s&from=%s&count=%s", userhandle, fmt.Sprint(from), fmt.Sprint(count))
	fmt.Println(url)
	var wrapper Results
	FetchJSON(url, &wrapper)
	if len(wrapper.Result) == 0 {
		return submissions, errors.New("Empty List")

	}
	for _, submission := range wrapper.Result {
		subURL := fmt.Sprintf(BaseURL+"/contest/%s/submission/%s", fmt.Sprint(submission.ContestID), fmt.Sprint(submission.ID))
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

func main() {
	start := 1
	for {

		submission, err := GetUserSubmissions("Siddhant1", start*50+1, 50)
		fmt.Println(submission)
		if err != nil {
			break
		}
		start++
	}
}

//FetchJSON ..
func FetchJSON(url string, wrapper interface{}) error {
	res, err := client.Get(url)
	if err != nil {
		log.Println(err)
		return err
	}
	if res.StatusCode != http.StatusOK {
		log.Println(res)
		return errors.New(fmt.Sprint(res))

	}
	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)

	dec.Decode(&wrapper)
	//fmt.Print(wrapper)
	return err

}
