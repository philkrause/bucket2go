package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var BUCKET_URL, bool = os.LookupEnv("BUCKET_URL")
var client = &http.Client{}

type Flags struct {
	Project  string `json:"project"`
	Username string `json:"username"`
}

func init() {

	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "create a repository",
		Run: func(cmd *cobra.Command, args []string) {
			createRepo()
			fmt.Println("Repository Created")
		},

		// createCmd.Flags().StringP(&Repository.Slug, "r", "create.repo", "Set the name of the repo slug")
	}
	flags := Flags{}
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP(flags.Project, "p", "HAR", "Set the name of the project")
	createCmd.Flags().StringP(flags.Username, "n", "Dev-Ops", "Set the name of the flags")
}

// func GetFlags() Flags {
// 	flags := Flags{}

// 	flag.StringVar(&flags.Project, "p", "HAR", "Set the name of the project")
// 	flag.StringVar(&flags.Username, "n", "Dev-Ops", "Set the name of the flags")

// 	return flags
// }

func createRepo() {
	var newRepo Flags

	bytesRepresentation, err := json.Marshal(newRepo)
	if err != nil {
		log.Fatalln(err)
	}

	request, err := http.NewRequest("POST", BUCKET_URL, bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-event-key", "repo.create")

	client.Do(request)

}
