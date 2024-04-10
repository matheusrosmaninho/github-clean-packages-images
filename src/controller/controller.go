package controller

import (
	"fmt"

	"github.com/matheusrosmaninho/github-clean-packages-images/services"
	"github.com/matheusrosmaninho/github-clean-packages-images/usecase"
)

func Start(repoToken, organization, packagesMonitoredString, packageType string, numberVersions int) error {
	packagesMonitored, err := usecase.NewPackageMonitored(packagesMonitoredString)
	if err != nil {
		return err
	}

	for _, p := range packagesMonitored.Values {
		numberVersionsKeep := numberVersions
		packageVersions, err := services.GetPackagesVersions(repoToken, organization, packageType, p)
		if err != nil {
			return err
		}

		fmt.Printf("The package %s has %d versions ...\n", p, len(packageVersions))

		if numberVersionsKeep >= len(packageVersions) {
			fmt.Printf("The package %s has less versions than the number of versions to keep ...\n", p)
			fmt.Println("----------------------------------------")
			continue
		}

		for _, packageVersion := range packageVersions {
			if numberVersionsKeep > 0 {
				fmt.Printf("The version %s of the package %s will not be deleted ...\n", packageVersion.Name, p)
				numberVersionsKeep--
				continue
			}

			if len(packageVersion.Metadata.Container.Tags) > 0 {
				for _, tag := range packageVersion.Metadata.Container.Tags {
					if tag == "buildcache" || tag == "latest" {
						fmt.Printf("The version %s of the package %s has the tag %s, so it will not be deleted ...\n", packageVersion.Name, p, tag)
						numberVersionsKeep--
						continue
					}
				}
			}

			err := services.DeletePackageVersion(repoToken, organization, packageType, p, packageVersion.ID)
			if err != nil {
				return err
			}
			fmt.Printf("The version %s of the package %s was deleted ...\n", packageVersion.Name, p)
			numberVersionsKeep--
		}
		fmt.Println("----------------------------------------")
	}
	return nil
}
