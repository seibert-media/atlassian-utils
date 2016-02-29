package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/bborbe/atlassian_utils/jira_software"
	command_list "github.com/bborbe/command/list"
	debian_config "github.com/bborbe/debian_utils/config"
	debian_config_builder "github.com/bborbe/debian_utils/config_builder"
	debian_config_parser "github.com/bborbe/debian_utils/config_parser"
	debian_copier "github.com/bborbe/debian_utils/copier"
	debian_package_creator "github.com/bborbe/debian_utils/package_creator"
	debian_package_creator_archive "github.com/bborbe/debian_utils/package_creator_archive"
	debian_package_creator_by_reader "github.com/bborbe/debian_utils/package_creator_by_reader"
	debian_tar_gz_extractor "github.com/bborbe/debian_utils/tar_gz_extractor"
	debian_zip_extractor "github.com/bborbe/debian_utils/zip_extractor"
	http_client_builder "github.com/bborbe/http/client_builder"
	http_requestbuilder "github.com/bborbe/http/requestbuilder"
	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

const (
	PARAMETER_LOGLEVEL               = "loglevel"
	PARAMETER_CONFIG                 = "config"
	PARAMETER_CONFLUENCE_TAR_GZ_PATH = "path"
	PARAMETER_CONFLUENCE_VERSION     = "version"
)

type ConfigBuilderWithConfig func(config *debian_config.Config) debian_config_builder.ConfigBuilder

func main() {
	defer logger.Close()
	logLevelPtr := flag.String(PARAMETER_LOGLEVEL, log.INFO_STRING, log.FLAG_USAGE)
	tarGzPathPtr := flag.String(PARAMETER_CONFLUENCE_TAR_GZ_PATH, "", "path to  tar gz")
	versionPtr := flag.String(PARAMETER_CONFLUENCE_VERSION, "", "version")
	configPtr := flag.String(PARAMETER_CONFIG, "", "path to config")
	flag.Parse()
	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Debugf("set log level to %s", *logLevelPtr)

	runtime.GOMAXPROCS(runtime.NumCPU())

	commandListProvider := func() command_list.CommandList {
		return command_list.New()
	}
	config_parser := debian_config_parser.New()
	copier := debian_copier.New()
	zipExtractor := debian_zip_extractor.New()
	tarGzExtractor := debian_tar_gz_extractor.New()
	httpClientBuilder := http_client_builder.New().WithoutProxy()
	httpClient := httpClientBuilder.Build()
	requestbuilderProvider := http_requestbuilder.NewHttpRequestBuilderProvider()
	debianPackageCreator := debian_package_creator.New(commandListProvider, copier, tarGzExtractor.ExtractTarGz, zipExtractor.ExtractZip, httpClient.Do, requestbuilderProvider.NewHttpRequestBuilder)
	creatorByReader := debian_package_creator_by_reader.New(commandListProvider, debianPackageCreator, tarGzExtractor.ExtractTarGz)
	debianPackageCreatorArchive := debian_package_creator_archive.New(creatorByReader.CreatePackage)

	writer := os.Stdout
	err := do(writer, debianPackageCreatorArchive, config_parser, *tarGzPathPtr, *configPtr, *versionPtr)
	if err != nil {
		logger.Fatal(err)
		logger.Close()
		os.Exit(1)
	}
}

func do(writer io.Writer, debianPackageCreatorArchive debian_package_creator_archive.DebianPackageCreator, config_parser debian_config_parser.ConfigParser, tarGzPath string, configpath string, version string) error {
	if len(tarGzPath) == 0 {
		return fmt.Errorf("parameter %s missing", PARAMETER_CONFLUENCE_TAR_GZ_PATH)
	}
	var err error
	config := createDefaultConfig()
	if len(configpath) > 0 {
		if config, err = config_parser.ParseFileToConfig(config, configpath); err != nil {
			return err
		}
	}
	config_builder := debian_config_builder.NewWithConfig(config)
	config_builder.Version(version)
	config = config_builder.Build()
	if len(config.Version) == 0 {
		return fmt.Errorf("paramter %s missing", PARAMETER_CONFLUENCE_VERSION)
	}
	sourceDir := fmt.Sprintf("atlassian-jira-software-%s-standalone", config.Version)
	targetDir := jira_software.TARGET
	return debianPackageCreatorArchive.CreatePackage(tarGzPath, config, sourceDir, targetDir)
}

func createDefaultConfig() *debian_config.Config {
	config := debian_config.DefaultConfig()
	config.Name = jira_software.PACKAGE_NAME
	config.Architecture = jira_software.ARCH
	return config
}
