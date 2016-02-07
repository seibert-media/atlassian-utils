package main

import (
	"flag"
	"io"
	"os"
	"runtime"

	"fmt"

	"github.com/bborbe/atlassian_utils/confluence"
	atlassian_utils_latest_information "github.com/bborbe/atlassian_utils/latest_information"
	atlassian_utils_latest_version "github.com/bborbe/atlassian_utils/latest_version"
	http_client "github.com/bborbe/http/client"
 	http_client_builder "github.com/bborbe/http/client/builder"
 	"github.com/bborbe/log"
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
	latestInformations := atlassian_utils_latest_information.New(confluence.JSON_URL, httpClient.Get)
	latestVersion := atlassian_utils_latest_version.New(latestInformations.VersionInformations)

	writer := os.Stdout
	err := do(writer, latestVersion.LatestVersion)
	if err != nil {
		logger.Fatal(err)
		logger.Close()
		os.Exit(1)
	}
}

func do(writer io.Writer, latestVersion LatestVersion) error {
	version, err := latestVersion()
	if err != nil {
		return err
	}
	fmt.Fprintf(writer, "%s\n", version)
	return nil
}
