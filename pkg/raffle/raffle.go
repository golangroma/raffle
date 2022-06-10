package raffle

import (
	"context"

	"github.com/google/go-github/v45/github"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . IssueService
type IssueService interface {
	ListByRepo(ctx context.Context, owner, repo string, opts *github.IssueListByRepoOptions) ([]*github.Issue, *github.Response, error)
}

func getUsersFromIssues(issueService IssueService, user, repo string, labels []string) ([]string, error) {
	issues, err := getIssues(issueService, user, repo, labels)
	if err != nil {
		return nil, err
	}

	return getUniqueUsersFromIssues(issues), nil
}

func getIssues(issueService IssueService, user, repo string, labels []string) ([]*github.Issue, error) {
	opt := &github.IssueListByRepoOptions{
		ListOptions: github.ListOptions{PerPage: 50},
		Labels:      labels,
	}

	var allIssues []*github.Issue
	for {
		issues, resp, err := issueService.ListByRepo(context.Background(), user, repo, opt)
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

func getUniqueUsersFromIssues(issues []*github.Issue) []string {
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
