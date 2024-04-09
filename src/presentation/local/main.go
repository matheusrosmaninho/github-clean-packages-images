package main

import (
	"os"

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

	err = controller.Start(repoToken, organization, packagesMonitored, packageType)
	if err != nil {
		panic(err)
	}
}
