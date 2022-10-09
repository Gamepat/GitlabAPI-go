package package_api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type GitlabApiConnection struct {
	Client       *http.Client
	GitlabApiUrl string
	Token        string
}

// GetPackages Get packages from project
func (a *GitlabApiConnection) GetPackages(id string, oBy orderBy, srt sort,
	pkgType packageType, pkgName string, incVl bool, status packageStatus) (*PackageApiResponse, error) {

	if a.Client == nil || a.Token == "" {
		return nil, errors.New("members Client or Token not set")
	}

	req, err := http.NewRequest(http.MethodGet, a.GitlabApiUrl, nil)
	if err != nil {
		return nil, err
	}

	// Set auth
	req.Header.Set("PRIVATE-TOKEN", a.Token)

	// Set attributes
	q := req.URL.Query()
	q.Add("order_by", oBy.String())
	q.Add("sort", srt.String())
	if pkgType != PackageTypeAll {
		q.Add("package_type", pkgType.String())
	}
	if pkgName != "" {
		q.Add("package_name", pkgName)
	}
	if incVl {
		q.Add("include_versionless", "true")
	}
	q.Add("status", status.String())
	req.URL.RawQuery = q.Encode()

	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("request returned with status: " + resp.Status)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data PackageApiResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	log.Printf("Results: %v\n", data)

	return &data, nil
}

func GetPackage(pkgId string) {

}

func DeletePackage(pkgId string) {

}

func GetPackageFiles(pkgId string) {

}
