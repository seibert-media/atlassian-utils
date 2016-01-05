package latest_information

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	confluence_information "github.com/bborbe/atlassian_utils/information"
	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

const JSON_URL = "https://my.atlassian.com/download/feeds/current/confluence.json"

type Download func(url string) (resp *http.Response, err error)

type VersionInfo interface {
	VersionInformations() ([]confluence_information.VersionInformation, error)
}

type versionInfo struct {
	download Download
}

func New(download Download) *versionInfo {
	v := new(versionInfo)
	v.download = download
	return v
}

func (v *versionInfo) VersionInformations() ([]confluence_information.VersionInformation, error) {
	logger.Debugf("VersionInformations")
	content, err := getContent(v.download)
	if err != nil {
		return nil, err
	}
	return parseInfos(content)
}

func getContent(download Download) ([]byte, error) {
	var resp *http.Response
	var err error
	if resp, err = download(JSON_URL); err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func parseInfos(content []byte) ([]confluence_information.VersionInformation, error) {
	var list []confluence_information.VersionInformation
	c := content[10 : len(content) - 1]
	if err := json.Unmarshal(c, &list); err != nil {
		return nil, err
	}
	return list, nil
}
