package cmd

import (
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	repo       string
	version    string
	verbose    bool
	outputJSON bool

	rootCmd = &cobra.Command{
		Use:   "discussions",
		Short: "Manipulate GitHub discussions",
		Long:  `Manipulate GitHub discussions.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.SetVersionTemplate("{{printf \"This is my version %s\" .Version}}")
	rootCmd.Version = version
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().StringVarP(&repo, "repo", "R", "", "Select another repository using the [HOST/]OWNER/REPO format")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Show more output")

	listCmd.Flags().BoolVar(&outputJSON, "json", false, "Show JSON output")
	rootCmd.AddCommand(listCmd)
}

func getRepositoryNameFromLocalDirectory(dir string) (string, error) {
	out, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()
	if err != nil {
		return "", err
	}
	URL := strings.TrimSpace(string(out))
	if strings.HasPrefix(URL, "git@") {
		URL = strings.Replace(URL, "git@", "", 1)
		URL = strings.Replace(URL, ".git", "", 1)
		URL = strings.Replace(URL, ":", "/", 1)
		URL = "https://" + URL
	}

	return URL, nil
}

func initConfig() {
	if verbose {
		log.SetLevel(log.DebugLevel)
	}
	if repo == "" {
		log.Debug("No repository specified, trying to get it from the local directory")
		repo, _ = getRepositoryNameFromLocalDirectory(".")
		if repo == "" {
			log.Fatal("No repository specified and no repository found in the local directory")
		}
	}
	if !strings.HasPrefix(repo, "https://") {
		if len(strings.Split(repo, "/")) != 2 {
			log.Fatalf("Invalid repository format: '%s', must be [HOST]/OWNER/REPO", repo)
		}
		repo = "https://github.com/" + repo
	}
	log.Debug("Using repository: " + repo)
}
