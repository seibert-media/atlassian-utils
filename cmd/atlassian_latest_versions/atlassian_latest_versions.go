package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"sort"
	"sync"

	"github.com/seibert-media/atlassian-utils/bamboo"
	"github.com/seibert-media/atlassian-utils/bitbucket"
	"github.com/seibert-media/atlassian-utils/confluence"
	"github.com/seibert-media/atlassian-utils/crowd"
	"github.com/seibert-media/atlassian-utils/jira_core"
	"github.com/seibert-media/atlassian-utils/jira_servicedesk"
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
	err := do(
		writer,
		bambooLatestVersion.LatestVersion,
		confluenceLatestVersion.LatestVersion,
		jiraCoreLatestVersion.LatestVersion,
		jiraServiceDeskLatestVersion.LatestVersion,
		jiraSoftwareLatestVersion.LatestVersion,
		bitbucketLatestVersion.LatestVersion,
		crowdLatestVersion.LatestVersion,
	)
	if err != nil {
		glog.Exit(err)
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
				glog.V(2).Infof("fetch version failed: %v", err)
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
