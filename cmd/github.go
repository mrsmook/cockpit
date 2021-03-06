package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// Repositories is a representation of a github repos list
type Repositories []struct {
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    struct {
		Login string `json:"login"`
		ID    int    `json:"id"`
	} `json:"owner"`
	HTMLURL     string      `json:"html_url"`
	Description interface{} `json:"description"`
	Fork        bool        `json:"fork"`
	URL         string      `json:"url"`
	CreatedAt   time.Time   `json:"created_at"`
	GitURL      string      `json:"git_url"`
	Homepage    interface{} `json:"homepage"`
	Language    string      `json:"language"`
}


// sayhelloCmd represents the sayhello command
var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "get mrsmook github repos list",
	Long: `Get the repositories list of MrSmooK from github and parse them
developed to be run under cron task`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := http.Get("https://api.github.com/users/mrsmook/repos")
		if err != nil {
			log.Fatal(err)
		}
		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		if res.StatusCode != http.StatusOK {
			log.Fatal("Unexpected status code", res.StatusCode)
		}
		if err != nil {
			log.Fatal(err)
		}
		var data Repositories
		err = json.Unmarshal(body, &data)
		repo, _ := json.Marshal(data)
		err = ioutil.WriteFile("output.json", repo, 0644)
	},
}

func init() {
	rootCmd.AddCommand(githubCmd)
}
