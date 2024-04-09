package main

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/matheusrosmaninho/github-clean-packages-images/controller"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	repoToken := os.Getenv("INPUT_REPO_TOKEN")
	organization := os.Getenv("INPUT_ORGANIZATION")
	packagesMonitored := os.Getenv("INPUT_LIST_PACKAGE_MONITORED")
	packageType := os.Getenv("INPUT_PACKAGE_TYPE")
	numberVersionsKeepString := os.Getenv("INPUT_NUMBER_VERSIONS")

	numberVersionsKeep, err := strconv.Atoi(numberVersionsKeepString)
	if err != nil {
		panic(err)
	}

	err = controller.Start(repoToken, organization, packagesMonitored, packageType, numberVersionsKeep)
	if err != nil {
		panic(err)
	}
}
