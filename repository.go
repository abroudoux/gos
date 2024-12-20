package main

import (
	"fmt"
	"net/http"
)

func getLatestRealeaseVersion(userName string, repoName string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", userName, repoName)
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error while fetching latest release: %v", err)
	}

	latestVersion := res.Header.Get("tag_name")
	return latestVersion, nil
}