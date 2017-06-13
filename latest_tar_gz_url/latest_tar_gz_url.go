package latest_tar_gz_url

import (
	"fmt"

	"strings"

	atlassian_information "github.com/bborbe/atlassian_utils/information"
	"github.com/golang/glog"
)

type VersionInformations func() ([]atlassian_information.VersionInformation, error)

type LatestTarGzUrl interface {
	LatestConfluenceTarGzUrl() (string, error)
}

type latestTarGzUrl struct {
	versionInformations VersionInformations
}

func New(versionInformations VersionInformations) *latestTarGzUrl {
	l := new(latestTarGzUrl)
	l.versionInformations = versionInformations
	return l
}

func (l *latestTarGzUrl) LatestTarGzUrl() (string, error) {
	glog.V(2).Infof("LatestTarGzUrl")
	infos, err := l.versionInformations()
	if err != nil {
		return "", err
	}
	glog.V(4).Infof("found %d infos", len(infos))
	for _, info := range infos {
		if len(info["tarUrl"]) > 0 && strings.Contains(info["tarUrl"], "tar.gz") {
			return info["tarUrl"], nil
		}
		if len(info["zipUrl"]) > 0 && strings.Contains(info["zipUrl"], "tar.gz") {
			return info["zipUrl"], nil
		}
	}
	return "", fmt.Errorf("can't find tar.gz in tarUrl or zipUrl")
}
