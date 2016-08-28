package latest_tar_gz_url

import (
	"fmt"

	"strings"

	atlassian_information "github.com/bborbe/atlassian_utils/information"
	"github.com/golang/glog"
)

type VersionInformations func() ([]atlassian_information.VersionInformation, error)

type LatestConfluenceTarGzUrl interface {
	LatestConfluenceTarGzUrl() (string, error)
}

type latestConfluenceTarGzUrl struct {
	versionInformations VersionInformations
}

func New(versionInformations VersionInformations) *latestConfluenceTarGzUrl {
	l := new(latestConfluenceTarGzUrl)
	l.versionInformations = versionInformations
	return l
}

func (l *latestConfluenceTarGzUrl) LatestConfluenceTarGzUrl() (string, error) {
	glog.V(2).Infof("LatestConfluenceTarGzUrl")
	infos, err := l.versionInformations()
	if err != nil {
		return "", err
	}
	for _, info := range infos {
		if len(info["tarUrl"]) > 0 && strings.Contains(info["tarUrl"], "tar.gz") {
			return info["tarUrl"], nil
		}
		if len(info["zipUrl"]) > 0 && strings.Contains(info["zipUrl"], "tar.gz") {
			return info["zipUrl"], nil
		}
	}
	return "", fmt.Errorf("not found")
}
