package main

import (
	"GitlabApi/api/package_api"
	"GitlabApi/utils"
	"log"
	"net/http"
)

const gitlabApiUrl string = "https://gitlab.com/api/v4/"

func main() {

	// Load api-token from secrets file
	secrets, err := utils.LoadSecrets("blub.json")
	if err != nil {
		log.Fatalf("error: %s\n", err)

	}

	con := package_api.GitlabApiConnection{
		Client:       &http.Client{},
		GitlabApiUrl: gitlabApiUrl,
		Token:        secrets.ApiToken,
	}

	_, err = con.GetPackages("abc", package_api.OrderByCreatedAt, package_api.SortAscending, package_api.PackageTypeAll, "", false, package_api.PackageStatusDefault)
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

}
