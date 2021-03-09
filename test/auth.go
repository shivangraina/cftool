package main

import (
	"fmt"
	"log"

	oauth2ns "github.com/nmrshll/oauth2-noserver"
	"golang.org/x/oauth2"
)

func main() {
	conf := &oauth2.Config{
		ClientID:     "94593bdeb5f7844ed3b8",                     // also known as client key sometimes
		ClientSecret: "6790b58edb5e6c12e006501345896968b8493298", // also known as secret key
		Scopes:       []string{"account"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}

	client, err := oauth2ns.AuthenticateUser(conf)
	if err != nil {
		log.Fatal(err)
	}

	// use client.Get / client.Post for further requests, the token will automatically be there
	res, _ := client.Get("https://api.github.com/repos/shivangraina/tool/actions/artifacts")
	fmt.Println(res)
}
