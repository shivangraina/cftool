package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	oauth2ns "github.com/nmrshll/oauth2-noserver"
	"golang.org/x/oauth2"
)

//GetClientWithToken ..
func GetClientWithToken() *oauth2ns.AuthorizedClient {
	conf := &oauth2.Config{
		ClientID:     "94593bdeb5f7844ed3b8",                     // also known as client key sometimes
		ClientSecret: "6790b58edb5e6c12e006501345896968b8493298", // also known as secret key
		Scopes:       []string{"repo"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}

	client, err := oauth2ns.AuthenticateUser(conf)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

//CreateEmptyRepositry ..
func CreateEmptyRepositry(client *oauth2ns.AuthorizedClient, RepoName string) (string, error) {

	values := map[string]string{"name": RepoName, "auto_init": "true", "private": "false"}

	val, _ := json.Marshal(values)
	req, err := http.NewRequest(http.MethodPost, "https://api.github.com/user/repos", bytes.NewBuffer(val))
	if err != nil {
		return "", fmt.Errorf("Error in Requesting")
	}

	req.Header.Add("Accept", "application/vnd.github.v3+json")

	res, err := client.Do(req)
	if res.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("Error in creating repo ,status code%d", res.StatusCode)
	}
	body, err := ioutil.ReadAll(res.Body)
	bodystring := string(body)
	return bodystring, nil

}

//CreateContestFiles ..
func CreateContestFiles(client *oauth2ns.AuthorizedClient,usercode UserCodeData, RepoName string, Username string) (string, error) {

	d := usercode.ProblemCode
	base64data := base64.StdEncoding.EncodeToString([]byte(d))
	Path := fmt.Sprintf("%s/%s.%s", fmt.Sprint(usercode.ContestID),usercode.ProblemIndex, GetLanguageExtension(usercode.Language))
    CommitMessage:=fmt.Sprintf("Added Code for %s%s",fmt.Sprint(usercode.ContestID),usercode.ProblemIndex)
	values := map[string]string{"message": CommitMessage, "content": base64data}
	val, _ := json.Marshal(values)
	URL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", Username, RepoName, Path)
	req, err := http.NewRequest(http.MethodPut, URL, bytes.NewBuffer(val))

	if err != nil {
		return "", fmt.Errorf("Error in Requesting")
	}

	req.Header.Add("Accept", "application/vnd.github.v3+json")
	res, err := client.Do(req)
	if res.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("Error in creating repo ,status code%d", res.StatusCode)
	}
	body, err := ioutil.ReadAll(res.Body)
	bodystring := string(body)
	return bodystring, nil

}

//GetLanguageExtension ..
func GetLanguageExtension(language string) string {
	var Languages map[string]string
	jsonFile, err := os.Open("language.json")
	defer jsonFile.Close()

	if err != nil {
		
		return "txt"
	}
	//check here
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Languages)
	return Languages[language]

}
