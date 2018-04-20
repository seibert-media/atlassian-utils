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

## Continuous integration

[Jenkins](https://jenkins.benjamin-borbe.de/job/Go-Atlassian-Utils/)

## Copyright and license

    Copyright (c) 2016, Benjamin Borbe <bborbe@rocketnews.de>
    All rights reserved.
    
    Redistribution and use in source and binary forms, with or without
    modification, are permitted provided that the following conditions are
    met:
    
       * Redistributions of source code must retain the above copyright
         notice, this list of conditions and the following disclaimer.
       * Redistributions in binary form must reproduce the above
         copyright notice, this list of conditions and the following
         disclaimer in the documentation and/or other materials provided
         with the distribution.

    THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
    "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
    LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
    A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
    OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
    SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
    LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
    DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
    THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
    (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
    OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
