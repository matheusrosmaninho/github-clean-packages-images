package controller

import (
	"fmt"

	"github.com/matheusrosmaninho/github-clean-packages-images/services"
	"github.com/matheusrosmaninho/github-clean-packages-images/usecase"
)

func Start(repoToken, organization, packagesMonitoredString, packageType string, numberVersionsKeep int) error {
	packagesMonitored, err := usecase.NewPackageMonitored(packagesMonitoredString)
	if err != nil {
		return err
	}

	for _, p := range packagesMonitored.Values {
		packageVersions, err := services.GetPackagesVersions(repoToken, organization, packageType, p)
		if err != nil {
			return err
		}

		if numberVersionsKeep >= len(packageVersions) {
			continue
		}

		message := fmt.Sprintf("Package %s has %d versions", p, len(packageVersions))
		fmt.Println(message)
		fmt.Println("----------------------")
	}
	return nil
}
