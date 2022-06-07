package raffle

import (
	"context"

	"github.com/google/go-github/v45/github"
)

func GetIssues(client *github.Client, user, repo string, labels []string) ([]*github.Issue, error) {
	opt := &github.IssueListByRepoOptions{
		ListOptions: github.ListOptions{PerPage: 50},
		Labels:      labels,
	}

	var allIssues []*github.Issue
	for {
		issues, resp, err := client.Issues.ListByRepo(context.Background(), user, repo, opt)
		if err != nil {
			return nil, err
		}
		allIssues = append(allIssues, issues...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return allIssues, nil
}

func GetUsersFromIssues(issues []*github.Issue) []string {
	users := []string{}

	usersMap := make(map[string]struct{})
	for _, issue := range issues {
		if issue.GetUser().Login != nil {
			username := *issue.GetUser().Login

			// if we haven't seen this user before, add it to the list
			if _, seen := usersMap[username]; !seen {
				users = append(users, username)
			}
			usersMap[username] = struct{}{}
		}
	}

	return users
}
