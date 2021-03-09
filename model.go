 package main
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

// Submission ..
type Submission struct{
	ContestID     int64  `json:"contestId"`
	ProblemIndex string   `json:"index"`
	ProblemName   string   `json:"name"`
	Verdict        string  `json:"verdict"`
	ProblemURL     string   `json:"problemurl"`
	ProblemRating   int64  `json:"rating"`
	Language 		string   `json:"language"`
	TimeStamp       string    `json:"timestamp"`	
	Tags            []string `json:"tags"`
	SubmissionID   int64      `json:"id"`
	SubmissionURL  string    `json:"submsissionurl"`
}
