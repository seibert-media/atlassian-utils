package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/bborbe/atlassian_utils/jira_software"
	atlassian_utils_latest_information "github.com/bborbe/atlassian_utils/latest_information"
	atlassian_utils_latest_tar_gz_url "github.com/bborbe/atlassian_utils/latest_tar_gz_url"
	atlassian_utils_latest_version "github.com/bborbe/atlassian_utils/latest_version"
	command_list "github.com/bborbe/command/list"
	debian_config "github.com/bborbe/debian_utils/config"
	debian_config_builder "github.com/bborbe/debian_utils/config_builder"
	debian_config_parser "github.com/bborbe/debian_utils/config_parser"
	debian_copier "github.com/bborbe/debian_utils/copier"
	debian_latest_package_creator "github.com/bborbe/debian_utils/latest_package_creator"
	debian_package_creator "github.com/bborbe/debian_utils/package_creator"
	debian_package_creator_by_reader "github.com/bborbe/debian_utils/package_creator_by_reader"
	debian_tar_gz_extractor "github.com/bborbe/debian_utils/tar_gz_extractor"
	debian_zip_extractor "github.com/bborbe/debian_utils/zip_extractor"
	http_client_builder "github.com/bborbe/http/client_builder"
	http_requestbuilder "github.com/bborbe/http/requestbuilder"
	"github.com/golang/glog"
)

const (
	PARAMETER_CONFIG = "config"
	PARAMETER_TARGET = "target"
)

type CreatePackage func(config *debian_config.Config, sourceDir string, targetDir string) error
type LatestVersion func() (string, error)

var (
	configPtr    = flag.String(PARAMETER_CONFIG, "", "path to config")
	targetDirPtr = flag.String(PARAMETER_TARGET, jira_software.TARGET, "target")
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	httpClientBuilder := http_client_builder.New()
	httpClient := httpClientBuilder.Build()
	latestInformations := atlassian_utils_latest_information.New(jira_software.JSON_URL, httpClient.Get)
	latestUrl := atlassian_utils_latest_tar_gz_url.New(latestInformations.VersionInformations)
	latestVersion := atlassian_utils_latest_version.New(latestInformations.VersionInformations)

	commandListProvider := func() command_list.CommandList {
		return command_list.New()
	}
	config_parser := debian_config_parser.New()
	copier := debian_copier.New()
	zipExtractor := debian_zip_extractor.New()
	tarGzExtractor := debian_tar_gz_extractor.New()
	requestbuilderProvider := http_requestbuilder.NewHTTPRequestBuilderProvider()
	debianPackageCreator := debian_package_creator.New(commandListProvider, copier, tarGzExtractor.ExtractTarGz, zipExtractor.ExtractZip, httpClient.Do, requestbuilderProvider.NewHTTPRequestBuilder)
	creatorByReader := debian_package_creator_by_reader.New(commandListProvider, debianPackageCreator, tarGzExtractor.ExtractTarGz)
	latestDebianPackageCreator := debian_latest_package_creator.New(httpClient.Get, latestUrl.LatestTarGzUrl, latestVersion.LatestVersion, creatorByReader.CreatePackage)

	err := do(
		latestDebianPackageCreator.CreateLatestDebianPackage,
		config_parser,
		*configPtr,
		latestVersion.LatestVersion,
		*targetDirPtr,
	)
	if err != nil {
		glog.Exit(err)
	}
}

func do(createPackage CreatePackage,
	config_parser debian_config_parser.ConfigParser,
	configpath string,
	latestVersion LatestVersion,
	targetDir string,
) error {
	var err error
	config := createDefaultConfig()
	if len(configpath) > 0 {
		if config, err = config_parser.ParseFileToConfig(config, configpath); err != nil {
			return err
		}
	}
	config_builder := debian_config_builder.NewWithConfig(config)
	config = config_builder.Build()
	config.Version, err = latestVersion()
	if err != nil {
		return err
	}
	sourceDir := fmt.Sprintf("atlassian-jira-software-%s-standalone", config.Version)
	return createPackage(config, sourceDir, targetDir)
}

func createDefaultConfig() *debian_config.Config {
	config := debian_config.DefaultConfig()
	config.Name = jira_software.PACKAGE_NAME
	config.Architecture = jira_software.ARCH
	return config
}
