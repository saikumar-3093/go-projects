package githubuserdata

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmbracelet/lipgloss"
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
		Ref      string
		Ref_type string
		Commits  []struct {
			Message string
		}
	}
}

func Event(user string) error {

	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events", user))

	if err != nil {
		return err
	}

	if resp.StatusCode == 404 {
		return fmt.Errorf("\033[31m\033[1muser not found. please check the username\033[0m")
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("\033[31m\033[1merror fetching data: %d\033[0m", resp.StatusCode)
	}

	var eventData []event

	err = json.NewDecoder(resp.Body).Decode(&eventData)

	if err != nil {
		return err
	}

	DisplayUserActivity(user, eventData)
	return nil
}

func DisplayUserActivity(user string, events []event) {

	if len(events) == 0 {
		fmt.Println("\033[33mNo events for the user:\033[0m", user)
	}

	commitCount := 0

	fmt.Printf("user \033[33m%s\033[0m recent activities\n", user)
	fmt.Println("Output:")
	var action string
	for _, event := range events {

		switch event.Type {

		case "PushEvent":
			commitCount = len(event.Payload.Commits)
			action = fmt.Sprintf("Pushed \033[33m\033[1m%v commit(s)\033[0m to \033[33m%s\033[0m Repository", commitCount, event.Repo.Name)

		case "ForkEvent":
			action = fmt.Sprintf("Forked \033[33m%s\033[0m Repository", event.Repo.Name)

		case "CreateEvent":
			if event.Payload.Ref_type == "branch" {
				action = fmt.Sprintf("Created \033[33m'%s'\033[0m Branch in \033[33m%s\033[0m Repository", event.Payload.Ref, event.Repo.Name)
			}
		case "WatchEvent":
			action = fmt.Sprintf("Added \033[33m%s\033[0m repository to favourites", event.Repo.Name)

		case "IssuesEvent":
			action = fmt.Sprintf("Created an Issue in \033[33m%s\033[0m repo", event.Repo.Name)
		default:
			action = fmt.Sprintf("\033[33m%s\033[0m in \033[33m%s\033[0m", event.Type, event.Repo.Name)
		}
		fmtOutput := lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, true, false).
			BorderForeground(lipgloss.Color("#3C3C3C")).
			Render(fmt.Sprintf("- %s", action))

		fmt.Println(fmtOutput)

	}

}
