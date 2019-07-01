package main

import (
	"cgt.name/pkg/go-mwclient"
)

type MediaWiki struct {
	Client *mwclient.Client
}

func (mw *MediaWiki) Init(URL string, userAgent string, username string, password string) {
	var err error
	// Initialize a *Client with New(), specifying the wiki's API URL
	// and your HTTP User-Agent. Try to use a meaningful User-Agent.
	//mw.Client, err = mwclient.New("https://rt.4amlunch.net/api.php", "myWikibot")
	mw.Client, err = mwclient.New(URL, userAgent)
	if err != nil {
		panic(err)
	}

	// Log in.
	//err = mw.Client.Login("Wonko@SkyNet", "kr35kvn8r7rra19r3c8td37c0ipm8pde")
	err = mw.Client.Login(username, password)
	if err != nil {
		panic(err)
	}
}

func (mw *MediaWiki) WritePage(title string, text string) {
	token, err := mw.Client.GetToken(mwclient.CSRFToken)
	if err != nil {
		panic(err)
	}

	// Specify parameters to send.
	parameters := map[string]string{
		"action": "edit",
		"format": "json",
		"title":  title,
		"text":   text,
		"token":  token,
		"bot":    "true",
	}

	// Make the request.
	_, err = mw.Client.Post(parameters)
	if err != nil {
		panic(err)
	}

	// Print the *jason.Object
	//fmt.Println(resp)
}
