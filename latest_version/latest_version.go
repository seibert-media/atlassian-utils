package latest_version

import (
	"fmt"

	atlassian_information "github.com/seibert-media/atlassian-utils/information"
)

type VersionInformations func() ([]atlassian_information.VersionInformation, error)

type LatestVersion interface {
	LatestVersion() (string, error)
}

type latestVersion struct {
	versionInformations VersionInformations
}

func New(versionInformations VersionInformations) *latestVersion {
	l := new(latestVersion)
	l.versionInformations = versionInformations
	return l
}

func (l *latestVersion) LatestVersion() (string, error) {
	infos, err := l.versionInformations()
	if err != nil {
		return "", err
	}
	for _, info := range infos {
		if info["platform"] == "Unix" || info["platform"] == "Mac OS X, Unix" || info["platform"] == "Unix, Mac OS X" {
			return info["version"], nil
		}
	}
	return "", fmt.Errorf("not found")
}
