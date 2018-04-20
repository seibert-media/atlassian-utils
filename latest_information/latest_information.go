package latest_information

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	atlassian_information "github.com/seibert-media/atlassian-utils/information"
	"github.com/golang/glog"
)

type Download func(url string) (resp *http.Response, err error)

type VersionInfo interface {
	VersionInformations() ([]atlassian_information.VersionInformation, error)
}

type versionInfo struct {
	download Download
	jsonUrl  string
}

func New(jsonUrl string, download Download) *versionInfo {
	v := new(versionInfo)
	v.download = download
	v.jsonUrl = jsonUrl
	return v
}

func (v *versionInfo) VersionInformations() ([]atlassian_information.VersionInformation, error) {
	glog.V(2).Infof("VersionInformations")
	content, err := getContent(v.download, v.jsonUrl)
	if err != nil {
		return nil, err
	}
	glog.V(4).Infof("json content: %s", string(content))
	return parseInfos(content)
}

func getContent(download Download, jsonUrl string) ([]byte, error) {
	var resp *http.Response
	var err error
	if resp, err = download(jsonUrl); err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func parseInfos(content []byte) ([]atlassian_information.VersionInformation, error) {
	var list []atlassian_information.VersionInformation
	c := content[10 : len(content)-1]
	if err := json.Unmarshal(c, &list); err != nil {
		return nil, err
	}
	return list, nil
}
