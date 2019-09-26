package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// var help string = "\n\n\n COMMANDS: \n\n    create:     Will create a new repository based on the flags provided\n           Flags: -p:  project  "

//Profile struct for COMMAND create
type Profile struct {
	Project      string `json:"project"`
	Repository   string `json:"reposlug"`
	Organization string `json:"organization,omitempty"`
}

//get Env var
var bucketURL, bool = os.LookupEnv("BUCKET_URL")

//create httpClient
var client = &http.Client{}

func init() {
	var p Profile
	var createCmd = &cobra.Command{
		Use:              "create",
		TraverseChildren: true,
		Short:            "Create a repository",
		Run: func(cmd *cobra.Command, args []string) {
			createRepo(p)
			fmt.Println("Repository Was Sent")
			// fmt.Printf("%+v", p)

		},
	}
	//set COMMAND create
	rootCmd.AddCommand(createCmd)
	//set Flags
	rootCmd.PersistentFlags().StringVarP(&p.Organization, "organization", "o", "", "Sets the organization name (optional)")
	rootCmd.PersistentFlags().StringVarP(&p.Project, "project", "p", "", "Sets the name of the bitbucket project key (required)")
	rootCmd.PersistentFlags().StringVarP(&p.Repository, "repository", "n", "", "Sets the repository name (required)")
	//make required flags
	createCmd.MarkPersistentFlagRequired("repository")
	createCmd.MarkPersistentFlagRequired("project")

	// rootCmd.SetHelpTemplate(help)

}

func createRepo(p Profile) *http.Response {

	//Convert JSON into bytes
	profile := new(bytes.Buffer)
	json.NewEncoder(profile).Encode(p)
	fmt.Printf("%+v", profile)

	if bucketURL != "" {
		fmt.Println("\nBitbucket URL is missing...Check your environment variables")
		os.Exit(1)
	}
	fmt.Println("Repo", p.Repository)

	request, err := http.NewRequest("POST", bucketURL, profile)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-event-key", "repo.create")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	return (response)
}
