package main

import (
	"flag"
	"io"
	"os"
	"runtime"

	"fmt"

	"github.com/bborbe/atlassian_utils/bamboo"
	atlassian_utils_latest_information "github.com/bborbe/atlassian_utils/latest_information"
	atlassian_utils_latest_version "github.com/bborbe/atlassian_utils/latest_version"
	http_client "github.com/bborbe/http/client"
	"github.com/bborbe/log"
	"github.com/bborbe/atlassian_utils/confluence"
	"github.com/bborbe/atlassian_utils/jira_core"
	"github.com/bborbe/atlassian_utils/jira_servicedesk"
	"github.com/bborbe/atlassian_utils/jira_software"
	"sort"
)

var logger = log.DefaultLogger

const (
	PARAMETER_LOGLEVEL = "loglevel"
)

type LatestVersion func() (string, error)

func main() {
	defer logger.Close()
	logLevelPtr := flag.String(PARAMETER_LOGLEVEL, log.INFO_STRING, log.FLAG_USAGE)
	flag.Parse()
	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Debugf("set log level to %s", *logLevelPtr)

	runtime.GOMAXPROCS(runtime.NumCPU())

	httpClient := http_client.New()

	bambooLatestInformations := atlassian_utils_latest_information.New(bamboo.JSON_URL, httpClient.Get)
	bambooLatestVersion := atlassian_utils_latest_version.New(bambooLatestInformations.VersionInformations)

	confluenceLatestInformations := atlassian_utils_latest_information.New(confluence.JSON_URL, httpClient.Get)
	confluenceLatestVersion := atlassian_utils_latest_version.New(confluenceLatestInformations.VersionInformations)

	jiraCorelatestInformations := atlassian_utils_latest_information.New(jira_core.JSON_URL, httpClient.Get)
	jiraCoreLatestVersion := atlassian_utils_latest_version.New(jiraCorelatestInformations.VersionInformations)

	jiraServiceDeskLatestInformations := atlassian_utils_latest_information.New(jira_servicedesk.JSON_URL, httpClient.Get)
	jiraServiceDeskLatestVersion := atlassian_utils_latest_version.New(jiraServiceDeskLatestInformations.VersionInformations)

	jiraSoftwareLatestInformations := atlassian_utils_latest_information.New(jira_software.JSON_URL, httpClient.Get)
	jiraSoftwareLatestVersion := atlassian_utils_latest_version.New(jiraSoftwareLatestInformations.VersionInformations)

	writer := os.Stdout
	err := do(writer, bambooLatestVersion.LatestVersion, confluenceLatestVersion.LatestVersion, jiraCoreLatestVersion.LatestVersion, jiraServiceDeskLatestVersion.LatestVersion, jiraSoftwareLatestVersion.LatestVersion)
	if err != nil {
		logger.Fatal(err)
		logger.Close()
		os.Exit(1)
	}
}

func do(writer io.Writer, bambooLatestVersion LatestVersion, confluenceLatestVersion LatestVersion, jiraCoreLatestVersion LatestVersion, jiraServiceDeskLatestVersion LatestVersion, jiraSoftwareLatestVersion LatestVersion) error {
	latestVersions := map[string]LatestVersion{
		"Bamoo":bambooLatestVersion,
		"Confluence":confluenceLatestVersion,
		"Jira-Core":jiraCoreLatestVersion,
		"Jira-ServiceDesk":jiraServiceDeskLatestVersion,
		"Jira-Software":jiraSoftwareLatestVersion,
	}
	list, err := doMap(latestVersions)
	if err != nil {
		return err
	}
	sort.Strings(list)
	for _, result := range list {
		fmt.Fprintf(writer, result)
	}
	return nil
}

func doMap(latestVersions map[string]LatestVersion) ([]string, error) {
	var list []string
	for name, version := range latestVersions {
		version, err := version()
		if err != nil {
			return nil, err
		}
		list = append(list, fmt.Sprintf("%s: %s\n", name, version))
	}
	return list, nil
}
