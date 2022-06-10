package raffle

import (
	"fmt"
	"sort"
	"time"
)

func List(issueService IssueService, user, repo string, labels []string) error {
	users, err := getUsersFromIssues(issueService, user, repo, labels)
	if err != nil {
		return err
	}

	sort.Strings(users)

	switch len(users) {
	case 0:
		fmt.Printf("There are no participants")
	case 1:
		fmt.Printf("There is only 1 participant: %s\n", users[0])
	default:
		fmt.Printf("There are %d partecipants:\n", len(users))
		for _, user := range users {
			time.Sleep(time.Millisecond * 500)
			fmt.Printf(" - %s\n", user)
		}
	}

	return nil
}
