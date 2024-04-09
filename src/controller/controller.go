package controller

import (
	"fmt"

	"github.com/matheusrosmaninho/github-clean-packages-images/services"
)

func Start(repoToken, organization, packagesMonitoredString, packageType string) error {
	packageVersions, err := services.GetPackagesVersions(repoToken, organization, packageType, packagesMonitoredString)
	if err != nil {
		return err
	}
	fmt.Println(packageVersions)
	return nil
}
