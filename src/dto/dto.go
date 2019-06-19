package dto

import (
	"strings"
)

type Repository struct {
	Name string `json:"name"`
	//Url string `json:"url"`
	State   string `json:"state"`
	Project project
	Size    int64
	Id      string
}
type project struct {
	State string
}

type Repositories struct {
	Value []Repository `json:"value"`
	Count int          `json:"count"`
}

type Branches struct {
	Count int `json:"count"`
}

type author struct {
	Name string
	Date string
}
type Commits struct {
	Commits []commit `json:"value"`
}

type commit struct {
	CommitId string
	Author   author
}

type Credentials struct {
	Username string `json:"Username"`
	Password string
	RepoURL  string
}

func GetCommitURL(repoURL string) string {
	return strings.TrimSpace(repoURL) + "/{repositoryId}/commits?api-version=5.0"
}

func GetBranchURL(repoURL string) string {
	return strings.TrimSpace(repoURL) + "/{repositoryId}/refs?api-version=5.0"
}
