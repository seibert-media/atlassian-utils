package latest_version

import (
	"fmt"

	confluence_information "github.com/bborbe/atlassian_utils/information"
)

type VersionInformations func() ([]confluence_information.VersionInformation, error)

type LatestVersion interface {
	LatestConfluenceVersion() (string, error)
}

type latestVersion struct {
	versionInformations VersionInformations
}

func New(versionInformations VersionInformations) *latestVersion {
	l := new(latestVersion)
	l.versionInformations = versionInformations
	return l
}

func (l *latestVersion) LatestConfluenceVersion() (string, error) {
	infos, err := l.versionInformations()
	if err != nil {
		return "", err
	}
	for _, info := range infos {
		if info["platform"] == "Unix" || info["platform"] == "Mac OS X, Unix" {
			return info["version"], nil
		}
	}
	return "", fmt.Errorf("not found")
}
