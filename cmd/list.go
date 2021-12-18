package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/shurcooL/graphql"
	log "github.com/sirupsen/logrus"

	"github.com/cli/go-gh"
	"github.com/spf13/cobra"
)

type DiscussionsSummary struct {
	Title  string
	Number int
}

func getGitHubDiscussions(name string, owner string) ([]DiscussionsSummary, error) {
	log.Debugf("Getting GitHub discussions for name: %s owner: %s", name, owner)

	client, err := gh.GQLClient(nil)
	if err != nil {
		return nil, err
	}

	var query struct {
		Repository struct {
			Discussions struct {
				Nodes []struct {
					Title  string
					Number int
				}
			} `graphql:"discussions(last: $last)"`
		} `graphql:"repository(name: $name, owner: $owner)"`
	}

	variables := map[string]interface{}{
		"name":  graphql.String(name),
		"owner": graphql.String(owner),
		"last":  graphql.Int(100),
	}
	err = client.Query("DiscussionsTitles", &query, variables)
	if err != nil {
		return nil, err
	}

	var titles []DiscussionsSummary

	for _, node := range query.Repository.Discussions.Nodes {
		titles = append(titles, DiscussionsSummary{Title: node.Title, Number: node.Number})
	}
	return titles, nil

}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List discussions",
	Long:  `List discussions.`,
	Run: func(cmd *cobra.Command, args []string) {
		name := strings.Replace(repo, "https://github.com/", "", 1)
		owner := strings.Split(name, "/")[0]
		name = strings.Split(name, "/")[1]
		discussions, err := getGitHubDiscussions(name, owner)
		if err != nil {
			log.Fatal(err)
		}

		if outputJSON {
			str, err := json.MarshalIndent(discussions, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(str))
		} else {
			for _, discussion := range discussions {
				fmt.Printf("[%d] %s\n", discussion.Number, discussion.Title)
			}
		}
	},
}
