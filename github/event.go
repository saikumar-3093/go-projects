package githubuserdata

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
)

type event struct {
	Type  string
	Actor struct {
		ID    int
		Login string
	}
	Repo struct {
		ID   int
		Name string
	}
	Payload struct {
		Commits []struct {
			Message string
		}
	}
}

func Client() error {
	token := "" //accesstoken
	repoToken := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	httpClient := oauth2.NewClient(context.Background(), repoToken)
	client := github.NewClient(httpClient)

	repos, _, _ := client.Repositories.List(context.Background(), "", nil)

	fmt.Println(len(repos))

	var rep github.Repository
	for _, v := range repos {
		if *v.Name == "go-project" {
			fmt.Println(*v.Name)
			rep = *v
		}
	}

	resp, err := httpClient.Get(*rep.EventsURL)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(resp.StatusCode)
	// fmt.Println(resp)
	var eventsData []event
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Printing Response")

		json.NewDecoder(resp.Body).Decode(&eventsData)
		// data, _ := io.ReadAll(resp.Body)

		for _, d := range eventsData {
			fmt.Println(d)
		}

	}
	return nil

}

func Event() error {
	fmt.Println("We will be fetching user activity here")

	resp, err := http.Get("https://api.github.com/users/saikumar-3093/events")

	if err != nil {
		return err
	}

	var eventData []event

	json.NewDecoder(resp.Body).Decode(&eventData)

	for _, data := range eventData {

		fmt.Println(data)
	}
	return nil
}
