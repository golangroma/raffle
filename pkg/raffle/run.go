package raffle

import (
	"fmt"
	"math/rand"
	"time"
)

func Run(issueService IssueService, user, repo string, labels []string) error {
	users, err := getUsersFromIssues(issueService, user, repo, labels)
	if err != nil {
		return err
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(users), func(i, j int) {
		users[i], users[j] = users[j], users[i]
	})

	fmt.Printf("Winner is: %s\n\n", users[0])

	fmt.Println("Full ranking:")
	for i, user := range users {
		fmt.Printf(" %3d) %s\n", i+1, user)
	}

	return nil
}
