package main

import (
	"flag"
	"io"
	"os"
	"runtime"

	"fmt"

	"github.com/seibert-media/atlassian-utils/jira_software"
	atlassian_utils_latest_information "github.com/seibert-media/atlassian-utils/latest_information"
	atlassian_utils_latest_version "github.com/seibert-media/atlassian-utils/latest_version"

	http_client_builder "github.com/bborbe/http/client_builder"
	"github.com/golang/glog"
)

type LatestVersion func() (string, error)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	httpClientBuilder := http_client_builder.New()
	httpClient := httpClientBuilder.Build()
	latestInformations := atlassian_utils_latest_information.New(jira_software.JSON_URL, httpClient.Get)
	latestVersion := atlassian_utils_latest_version.New(latestInformations.VersionInformations)

	writer := os.Stdout
	err := do(
		writer,
		latestVersion.LatestVersion,
	)
	if err != nil {
		glog.Exit(err)
	}
}

func do(
	writer io.Writer,
	latestVersion LatestVersion,
) error {
	version, err := latestVersion()
	if err != nil {
		return err
	}
	fmt.Fprintf(writer, "%s\n", version)
	return nil
}
