package main

import (
	"flag"
	"io"
	"os"
	"runtime"

	"fmt"

	"github.com/seibert-media/atlassian-utils/jira_servicedesk"
	atlassian_utils_latest_information "github.com/seibert-media/atlassian-utils/latest_information"
	atlassian_utils_latest_tar_gz_url "github.com/seibert-media/atlassian-utils/latest_tar_gz_url"

	http_client_builder "github.com/bborbe/http/client_builder"
	"github.com/golang/glog"
)

type LatestUrl func() (string, error)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	httpClientBuilder := http_client_builder.New()
	httpClient := httpClientBuilder.Build()
	latestInformations := atlassian_utils_latest_information.New(jira_servicedesk.JSON_URL, httpClient.Get)
	latestUrl := atlassian_utils_latest_tar_gz_url.New(latestInformations.VersionInformations)

	writer := os.Stdout
	err := do(
		writer,
		latestUrl.LatestTarGzUrl,
	)
	if err != nil {
		glog.Exit(err)
	}
}

func do(
	writer io.Writer,
	latestUrl LatestUrl,
) error {
	version, err := latestUrl()
	if err != nil {
		return err
	}
	fmt.Fprintf(writer, "%s\n", version)
	return nil
}
