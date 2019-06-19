package azure

import (
	"azure-git-utils/src/dto"
	"azure-git-utils/src/utils/http"
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func replaceRepositoryID(url string, repositoryID string) string {
	return strings.Replace(url, "{repositoryId}", repositoryID, -1)
}

func getBranches(creds dto.Credentials, repositoryID string) dto.Branches {
	replaceString := replaceRepositoryID(dto.GetBranchURL(creds.RepoURL), repositoryID)
	//fmt.Println(replaceString)
	httpresponse := http.Get(creds, replaceString)
	var branches dto.Branches
	//log.Println(httpresponse)
	if err := json.Unmarshal([]byte(httpresponse), &branches); err != nil {
		//log.Fatal(httpresponse)
		log.Fatal(err)
	}
	//fmt.Println(branches)
	return branches
}

func GetCommits(creds dto.Credentials, repositoryID string) dto.Commits {
	replaceString := replaceRepositoryID(dto.GetCommitURL(creds.RepoURL), repositoryID)
	//	fmt.Println(replaceString)
	httpresponse := http.Get(creds, replaceString)
	log.Println(httpresponse)
	var commits dto.Commits
	//log.Println(httpresponse)
	if err := json.Unmarshal([]byte(httpresponse), &commits); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(commits)
	return commits
}

func findRepoByName(creds dto.Credentials, name string) string {
	repos := getRepositories(creds, creds.RepoURL).Value
	idString := ""
	for _, s := range repos {
		if name == s.Name {
			idString = s.Id
			return idString
		}
	}
	return idString
}

func getRepositories(creds dto.Credentials, repoURL string) dto.Repositories {
	var repositories dto.Repositories
	httpresponse := http.Get(creds, repoURL)
	//log.Println(httpresponse)
	if err := json.Unmarshal([]byte(httpresponse), &repositories); err != nil {
		log.Fatal(httpresponse)
		log.Fatal(err)
	}
	return repositories
}

func HandleListRepo(creds dto.Credentials, showRepo bool) {
	if showRepo == true {
		repos := getRepositories(creds, creds.RepoURL).Value
		for _, s := range repos {
			//	fmt.Println(i,s)
			fmt.Println(s.Name)
		}
	}
}

func HandleCountRepo(creds dto.Credentials, optCountRepo bool) {
	if optCountRepo == true {
		fmt.Println("Total repositories = ", getRepositories(creds, creds.RepoURL).Count)

	}
}

func HandleoptRepoName(optNumBranches bool, optRepoName string) {

	if optRepoName != "" {
		if optNumBranches == false {
			fmt.Println("Please provide branch with -b")
		}
	}

}

func HandleNumBranches(creds dto.Credentials, optNumBranches bool, optRepoName string) {

	if optNumBranches == true {
		if optRepoName == "" {
			fmt.Println("Please specify repository name using -n or --repo")
			os.Exit(0)

		} else {
			repoID := findRepoByName(creds, optRepoName)
			//fmt.Println(repoID)
			fmt.Println("Number is branches for ", optRepoName+" is", getBranches(creds, repoID))
			//fmt.Println("Last commit ", getCommits(creds,repoID).Commits[0].Author)

		}

	}
}

func GetRepoAndBranchesTable(creds dto.Credentials, optListRepoWithBranches bool) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Repository", "Branch count"})
	repos := getRepositories(creds, creds.RepoURL).Value
	for _, s := range repos {
		//fmt.Println(s.Id)
		rec := []string{s.Name, strconv.Itoa(getBranches(creds, s.Id).Count)}
		table.Append(rec)
	}
	table.Render()

}

func HandleMaxBranches(creds dto.Credentials, optMaxBranches bool) {
	if optMaxBranches == true {
		repo, num := getRepoWithHighestBranches(creds)
		fmt.Println("Repository with max branches = "+repo.Name, " Number of Branches =", num)
	}
}

func getRepoWithHighestBranches(creds dto.Credentials) (dto.Repository, int) {
	var keys []int
	m := make(map[int]dto.Repository)
	repos := getRepositories(creds, creds.RepoURL).Value
	for _, s := range repos {
		m[getBranches(creds, s.Id).Count] = s
	}

	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return m[keys[len(keys)-1]], keys[len(keys)-1]
}
