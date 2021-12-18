package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/pete0emerson/gh-discussions/cmd"
	// "github.com/cli/go-gh"
)

func main() {

	debug := os.Getenv("DEBUG")
	if debug == "true" {
		log.SetLevel(log.DebugLevel)
	}

	cmd.Execute()
}

// 	fmt.Println(discussions.Hello())

// 	fmt.Println("hi world, this is the gh-discussions extension!")
// 	client, err := gh.RESTClient(nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	response := struct{ Login string }{}
// 	err = client.Get("user", &response)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("running as %s\n", response.Login)
// }

/*
	discussionsRepo := viper.GetString("discussions_repo")
	if discussionsRepo == "" {
		log.Debug("discussions_repo not set in config file")
		discussionsRepo, err = getRepositoryNameFromLocalDirectory(os.Getenv("PWD"))
		if err != nil {
			log.Debug

	foo := viper.Get("foo")
	fmt.Printf("foo: %v\n", foo)
	fmt.Println(viper.AllKeys())
*/
