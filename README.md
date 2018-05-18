# Atlassian Utils

Package provide some atlassian utils

## Install

`go get github.com/seibert-media/atlassian-utils/cmd/atlassian_latest_versions`

`go get github.com/seibert-media/atlassian-utils/cmd/bamboo_create_deb`

`go get github.com/seibert-media/atlassian-utils/cmd/bamboo_create_latest_deb`

`go get github.com/seibert-media/atlassian-utils/cmd/bamboo_latest_tar_gz_url`

`go get github.com/seibert-media/atlassian-utils/cmd/bamboo_latest_version`

`go get github.com/seibert-media/atlassian-utils/cmd/bitbucket_create_deb`

`go get github.com/seibert-media/atlassian-utils/cmd/bitbucket_create_latest_deb`

`go get github.com/seibert-media/atlassian-utils/cmd/bitbucket_latest_tar_gz_url`

`go get github.com/seibert-media/atlassian-utils/cmd/bitbucket_latest_version`

`go get github.com/seibert-media/atlassian-utils/cmd/confluence_create_deb`

`go get github.com/seibert-media/atlassian-utils/cmd/confluence_create_latest_deb`

`go get github.com/seibert-media/atlassian-utils/cmd/confluence_latest_tar_gz_url`

`go get github.com/seibert-media/atlassian-utils/cmd/confluence_latest_version`

`go get github.com/seibert-media/atlassian-utils/cmd/crowd_create_deb`

`go get github.com/seibert-media/atlassian-utils/cmd/crowd_create_latest_deb`

`go get github.com/seibert-media/atlassian-utils/cmd/crowd_latest_tar_gz_url`

`go get github.com/seibert-media/atlassian-utils/cmd/crowd_latest_version`

`go get github.com/seibert-media/atlassian-utils/cmd/jira_core_create_deb`

`go get github.com/seibert-media/atlassian-utils/cmd/jira_core_create_latest_deb`

`go get github.com/seibert-media/atlassian-utils/cmd/jira_core_latest_tar_gz_url`

`go get github.com/seibert-media/atlassian-utils/cmd/jira_core_latest_version`

`go get github.com/seibert-media/atlassian-utils/cmd/jira_servicedesk_create_deb`

`go get github.com/seibert-media/atlassian-utils/cmd/jira_servicedesk_create_latest_deb`

`go get github.com/seibert-media/atlassian-utils/cmd/jira_servicedesk_latest_tar_gz_url`

`go get github.com/seibert-media/atlassian-utils/cmd/jira_servicedesk_latest_version`

`go get github.com/seibert-media/atlassian-utils/cmd/jira_software_create_deb`

`go get github.com/seibert-media/atlassian-utils/cmd/jira_software_create_latest_deb`

`go get github.com/seibert-media/atlassian-utils/cmd/jira_software_latest_tar_gz_url`

`go get github.com/seibert-media/atlassian-utils/cmd/jira_software_latest_version`

## Create Confluence Debian Package

```
confluence_create_deb \
-logtostderr \
-v=2 \
-config confluence-config.json \
-path atlassian-confluence-5.9.5.tar.gz \
-version 5.9.5
```
Sample confluence.json

```
{
  "name": "confluence",
  "section": "utils",
  "priority": "optional",
  "architecture": "all",
  "maintainer": "Benjamin Borbe <bborbe@rocketnews.de>",
  "description": "Confluence",
  "postinst": "src/github.com/bborbe/atlassian-confluence/postinst",
  "postrm": "src/github.com/bborbe/atlassian-confluence/postrm",
  "prerm": "src/github.com/bborbe/atlassian-confluence/prerm",
  "depends": [
    "oracle-java8-installer"
  ],
  "files": [
    {
      "source": "src/github.com/bborbe/atlassian-confluence/etc/init.d/confluence",
      "target": "/etc/init.d/confluence"
    }
  ]
}
```
