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

//Get Env
var BUCKET_URL, bool = os.LookupEnv("BUCKET_URL")

//create httpClient
var client = &http.Client{}

type Profile struct {
	Project  string `json:"project"`
	Username string `json:"username"`
}

func init() {
	var p Profile
	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a repository",
		Run: func(cmd *cobra.Command, args []string) {
			createRepo()
			fmt.Println("Repository Created")
			fmt.Printf("%+v", p)

		},
	}

	rootCmd.AddCommand(createCmd)
	rootCmd.PersistentFlags().StringVarP(&p.Project, "project", "p", "HAR", "Sets the name of the project")
	rootCmd.PersistentFlags().StringVarP(&p.Username, "username", "n", "Dev-Ops", "Sets the username")
}

func createRepo() *http.Response {
	var p Profile

	//convert json into bytes
	bytesRepresentation, err := json.Marshal(p)
	if err != nil {
		log.Fatalln(err)
	}

	request, err := http.NewRequest("POST", BUCKET_URL, bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-event-key", "repo.create")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	return response
}
