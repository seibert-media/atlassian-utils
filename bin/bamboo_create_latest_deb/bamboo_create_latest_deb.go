package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"

	command_list "github.com/bborbe/command/list"
	debian_config "github.com/bborbe/debian_utils/config"
	debian_config_builder "github.com/bborbe/debian_utils/config_builder"
	debian_config_parser "github.com/bborbe/debian_utils/config_parser"
	debian_copier "github.com/bborbe/debian_utils/copier"
	debian_latest_package_creator "github.com/bborbe/debian_utils/latest_package_creator"
	debian_package_creator "github.com/bborbe/debian_utils/package_creator"
	debian_package_creator_by_reader "github.com/bborbe/debian_utils/package_creator_by_reader"

	"github.com/bborbe/atlassian_utils/bamboo"
	atlassian_utils_latest_information "github.com/bborbe/atlassian_utils/latest_information"
	atlassian_utils_latest_tar_gz_url "github.com/bborbe/atlassian_utils/latest_tar_gz_url"
	atlassian_utils_latest_version "github.com/bborbe/atlassian_utils/latest_version"
	"github.com/bborbe/debian_utils/tar_gz_extractor"
	http_client_builder "github.com/bborbe/http/client_builder"
	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

const (
	PARAMETER_LOGLEVEL = "loglevel"
	PARAMETER_CONFIG   = "config"
)

type CreatePackage func(config *debian_config.Config, sourceDir string, targetDir string) error
type LatestVersion func() (string, error)

func main() {
	defer logger.Close()
	logLevelPtr := flag.String(PARAMETER_LOGLEVEL, log.INFO_STRING, log.FLAG_USAGE)
	configPtr := flag.String(PARAMETER_CONFIG, "", "path to config")
	flag.Parse()
	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Debugf("set log level to %s", *logLevelPtr)

	runtime.GOMAXPROCS(runtime.NumCPU())

	httpClientBuilder := http_client_builder.New()
	httpClient := httpClientBuilder.Build()
	latestInformations := atlassian_utils_latest_information.New(bamboo.JSON_URL, httpClient.Get)
	latestUrl := atlassian_utils_latest_tar_gz_url.New(latestInformations.VersionInformations)
	latestVersion := atlassian_utils_latest_version.New(latestInformations.VersionInformations)

	commandListProvider := func() command_list.CommandList {
		return command_list.New()
	}
	copier := debian_copier.New()
	debianPackageCreator := debian_package_creator.New(commandListProvider, copier)
	extractor := tar_gz_extractor.New()
	creatorByReader := debian_package_creator_by_reader.New(commandListProvider, debianPackageCreator, extractor.ExtractTarGz)
	latestDebianPackageCreator := debian_latest_package_creator.New(httpClient.Get, latestUrl.LatestConfluenceTarGzUrl, latestVersion.LatestVersion, creatorByReader.CreatePackage)
	config_parser := debian_config_parser.New()

	writer := os.Stdout
	err := do(writer, latestDebianPackageCreator.CreateLatestDebianPackage, config_parser, *configPtr, latestVersion.LatestVersion)
	if err != nil {
		logger.Fatal(err)
		logger.Close()
		os.Exit(1)
	}
}

func do(writer io.Writer, createPackage CreatePackage, config_parser debian_config_parser.ConfigParser, configpath string, latestVersion LatestVersion) error {
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
	sourceDir := fmt.Sprintf("atlassian-bamboo-%s", config.Version)
	targetDir := bamboo.TARGET
	return createPackage(config, sourceDir, targetDir)
}

func createDefaultConfig() *debian_config.Config {
	config := debian_config.DefaultConfig()
	config.Name = bamboo.PACKAGE_NAME
	config.Architecture = bamboo.ARCH
	return config
}
