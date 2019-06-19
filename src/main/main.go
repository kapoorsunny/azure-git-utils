package main

import (
	"azure-git-utils/src/dto"
	"azure-git-utils/src/utils/azure"
	"azure-git-utils/src/utils/file"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/pborman/getopt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	CONFIGURATION_HOME = "/gotgit"
	CONFIGURATION_FILE = "/cred.db"
	VERSION            = "0.1"
	UTIL_NAME          = "goagit"
)

func isOneTimeSetup() bool {
	if !file.IsExist(os.Getenv("HOME") + CONFIGURATION_HOME) {
		os.Mkdir(os.Getenv("HOME")+CONFIGURATION_HOME, 0700)
		os.Create(os.Getenv("HOME") + CONFIGURATION_HOME + CONFIGURATION_FILE)
		//	fmt.Println(err)
		return true
	} else if !file.IsExist(os.Getenv("HOME") + CONFIGURATION_HOME + CONFIGURATION_FILE) {
		os.Create(os.Getenv("HOME") + CONFIGURATION_HOME + CONFIGURATION_FILE)
		return true
	} else {
		return false
	}
}

func oneTimeSetup() dto.Credentials {
	fmt.Print("Enter user name ")
	var userName, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	userName = strings.TrimSuffix(userName, "\n")
	fmt.Print("Enter PAN ")
	var password, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	password = strings.TrimSuffix(password, "\n")
	fmt.Print("Enter repository url ")
	var repoUrl, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	repoUrl = strings.TrimSuffix(repoUrl, "\n")
	return dto.Credentials{userName, password, repoUrl}

}

func main() {
	fmt.Println("Tool: ", UTIL_NAME)
	fmt.Println("Latest Version:" + VERSION)
	fmt.Println("-------------------")

	if isOneTimeSetup() {
		creds := oneTimeSetup()
		data, _ := json.Marshal(creds)
		ioutil.WriteFile(os.Getenv("HOME")+CONFIGURATION_HOME+CONFIGURATION_FILE, []byte(data), 0700)
	}
	returnData, _ := ioutil.ReadFile(os.Getenv("HOME") + CONFIGURATION_HOME + CONFIGURATION_FILE)
	//fmt.Println("Return data ",string(returnData))
	var credential dto.Credentials
	json.Unmarshal(returnData, &credential)
	// TODO: need to sort out command line options. still learning about it:)
	optListRepos := getopt.BoolLong("listRepos", 't', "true", "List all repositories ")
	optCountRepo := getopt.BoolLong("countRepo", 'c', "", "Count of repository ")
	optNumBranches := getopt.BoolLong("numBranches", 'b', "true", "List number of branches per repository ")
	optRepoName := getopt.StringLong("repo", 'n', "", "Name of repository ")
	optMaxBranches := getopt.BoolLong("maxBranch", 'm', "", "Repository with max branches ")
	//optLastCommit := getopt.StringLong("lastCommit", 'l', "", "Last commit on repository ")
	optHelp := getopt.BoolLong("help", 0, "Help")
	optListRepoWithBranches := getopt.BoolLong("RepoBranches", 'a', "", "Name of repository and branches table  ")

	getopt.Parse()

	if *optHelp {
		getopt.Usage()
		os.Exit(0)
	}
	azure.HandleListRepo(credential, *optListRepos)
	azure.HandleCountRepo(credential, *optCountRepo)
	azure.HandleNumBranches(credential, *optNumBranches, *optRepoName)
	azure.HandleMaxBranches(credential, *optMaxBranches)
	azure.HandleoptRepoName(*optNumBranches, *optRepoName)
	azure.GetRepoAndBranchesTable(credential, *optListRepoWithBranches)

	/*commits := azure.GetCommits(credential, "70626f4d-e6bc-484c-9315-d89814579bee")
	fmt.Println(commits.Commits[0])
	fmt.Println(commits.Commits[1])*/

	os.Exit(-1)

	/* if *optLastCommit == "" {

	}*/

}
