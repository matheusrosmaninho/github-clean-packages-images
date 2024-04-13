package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/matheusrosmaninho/github-clean-packages-images/controller"
)

func main() {
	repoToken := os.Getenv("INPUT_REPO_TOKEN")
	organization := os.Getenv("INPUT_ORGANIZATION")
	packagesMonitored := os.Getenv("INPUT_LIST_PACKAGE_MONITORED")
	packageType := os.Getenv("INPUT_PACKAGE_TYPE")
	numberVersionsKeepString := os.Getenv("INPUT_NUMBER_VERSIONS")

	fmt.Println("Starting the process to clean the images ...")

	numberVersionsKeep, err := strconv.Atoi(numberVersionsKeepString)
	if err != nil {
		panic(err)
	}

	if numberVersionsKeep <= 0 {
		panic("The number of versions to keep must be greater than 0 ...")
	}

	if strings.TrimSpace(packagesMonitored) == "" {
		panic("The list of packages monitored cannot be empty ...")
	}

	err = controller.Start(repoToken, organization, packagesMonitored, packageType, numberVersionsKeep)
	if err != nil {
		panic(err)
	}

	fmt.Println("Process finished successfully!")
}
