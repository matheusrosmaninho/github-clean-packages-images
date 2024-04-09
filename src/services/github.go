package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	GITHUB_API_URL = "https://api.github.com"
)

type PackageVersionDetail struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	URL            string `json:"url"`
	PackageHTMLURL string `json:"package_html_url"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	HTMLURL        string `json:"html_url"`
	Metadata       struct {
		PackageType string `json:"package_type"`
		Container   struct {
			Tags []string `json:"tags"`
		} `json:"container"`
	} `json:"metadata"`
}

var headersBase = map[string]string{
	"Accept":               "application/vnd.github+json",
	"X-GitHub-Api-Version": "2022-11-28",
}

func GetPackagesVersions(repoToken, organization, packageType, packageName string) ([]PackageVersionDetail, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s/orgs/%s/packages/%s/%s/versions", GITHUB_API_URL, organization, packageType, packageName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("error creating request: " + err.Error())
	}

	for key, value := range headersBase {
		req.Header.Set(key, value)
	}
	req.Header.Add("Authorization", "Bearer "+repoToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("error making request: " + err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("error reading response: " + err.Error())
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.New("error response status: " + resp.Status + string(body))
	}

	var packages []PackageVersionDetail
	err = json.Unmarshal(body, &packages)

	if err != nil {
		return nil, errors.New("error decoding response: " + err.Error())
	}

	return packages, nil
}

func DeletePackageVersion(repoToken, organization, packageType, packageName string, packageVersionId int) error {
	client := &http.Client{}
	url := fmt.Sprintf("%s/orgs/%s/packages/%s/%s/versions/%d", GITHUB_API_URL, organization, packageType, packageName, packageVersionId)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return errors.New("error creating request: " + err.Error())
	}

	for key, value := range headersBase {
		req.Header.Set(key, value)
	}
	req.Header.Add("Authorization", "Bearer "+repoToken)

	resp, err := client.Do(req)
	if err != nil {
		return errors.New("error making request: " + err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error reading response: " + err.Error())
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Println(url)
		return errors.New("error response status: " + resp.Status + string(body))
	}

	return nil
}
