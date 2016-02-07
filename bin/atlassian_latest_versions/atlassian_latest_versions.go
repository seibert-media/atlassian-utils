package main

import (
	"flag"
	"io"
	"os"
	"runtime"

	"fmt"

	"sort"

	"sync"

	"github.com/bborbe/atlassian_utils/bamboo"
	"github.com/bborbe/atlassian_utils/bitbucket"
	"github.com/bborbe/atlassian_utils/confluence"
	"github.com/bborbe/atlassian_utils/crowd"
	"github.com/bborbe/atlassian_utils/jira_core"
	"github.com/bborbe/atlassian_utils/jira_servicedesk"
	"github.com/bborbe/atlassian_utils/jira_software"
	atlassian_utils_latest_information "github.com/bborbe/atlassian_utils/latest_information"
	atlassian_utils_latest_version "github.com/bborbe/atlassian_utils/latest_version"
	http_client "github.com/bborbe/http/client" 	http_client_builder "github.com/bborbe/http/client/builder" 	"github.com/bborbe/log"
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

	httpClientBuilder := http_client_builder.New()
	httpClient := http_client.New(httpClientBuilder.Build())

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

	bitbucketLatestInformations := atlassian_utils_latest_information.New(bitbucket.JSON_URL, httpClient.Get)
	bitbucketLatestVersion := atlassian_utils_latest_version.New(bitbucketLatestInformations.VersionInformations)

	crowdLatestInformations := atlassian_utils_latest_information.New(crowd.JSON_URL, httpClient.Get)
	crowdLatestVersion := atlassian_utils_latest_version.New(crowdLatestInformations.VersionInformations)

	writer := os.Stdout
	err := do(writer, bambooLatestVersion.LatestVersion, confluenceLatestVersion.LatestVersion, jiraCoreLatestVersion.LatestVersion, jiraServiceDeskLatestVersion.LatestVersion, jiraSoftwareLatestVersion.LatestVersion, bitbucketLatestVersion.LatestVersion, crowdLatestVersion.LatestVersion)
	if err != nil {
		logger.Fatal(err)
		logger.Close()
		os.Exit(1)
	}
}

func do(writer io.Writer, bambooLatestVersion LatestVersion, confluenceLatestVersion LatestVersion, jiraCoreLatestVersion LatestVersion, jiraServiceDeskLatestVersion LatestVersion, jiraSoftwareLatestVersion LatestVersion, bitbucketLatestVersion LatestVersion, crowdLatestVersion LatestVersion) error {
	latestVersions := map[string]LatestVersion{
		"Bamoo":            bambooLatestVersion,
		"Confluence":       confluenceLatestVersion,
		"Jira-Core":        jiraCoreLatestVersion,
		"Jira-ServiceDesk": jiraServiceDeskLatestVersion,
		"Jira-Software":    jiraSoftwareLatestVersion,
		"Bitbucket":        bitbucketLatestVersion,
		"Crowd":            crowdLatestVersion,
	}
	list := doMap(latestVersions)
	sort.Strings(list)
	for _, result := range list {
		fmt.Fprintf(writer, "%s\n", result)
	}
	return nil
}

func doMap(latestVersions map[string]LatestVersion) []string {
	var wg sync.WaitGroup
	var list []string
	results := make(chan string)
	done := make(chan bool)
	go func() {
		for result := range results {
			list = append(list, result)
		}
		done <- true
	}()
	for n, v := range latestVersions {
		wg.Add(1)
		version := v
		name := n
		go func() {
			version, err := version()
			if err != nil {
				logger.Debugf("fetch version failed: %v", err)
				version = "failed"
			}
			results <- fmt.Sprintf("%s: %s", name, version)
			wg.Done()
		}()
	}
	wg.Wait()
	close(results)

	<-done

	return list
}
