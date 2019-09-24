package options

type Project struct {
	Name string `json:name`
}

type myFlag struct{}

func (f myFlag) HasChanged() bool { return false }
func (f myFlag) Name() string     { return "my-flag-name" }

// func createRepo(options SyncBuildOptions) *Project {
// 	project := Project{}

// 	client := &http.Client{}

// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s/app/rest/projects/name:%s", options.URL, options.Project), nil)

// 	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", options.Token))
// 	req.Header.Add("Accept", "application/json")

// 	resp, err := client.Do(req)

// 	if err != nil {
// 		panic(err)
// 	}

// 	if resp.StatusCode == 404 {
// 		return nil
// 	}

// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)

// 	err2 := json.Unmarshal([]byte(body), &project)

// 	if err2 != nil {
// 		panic(err2)
// 	}

// 	return &project
// }
