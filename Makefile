install:
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/atlassian_latest_versions/atlassian_latest_versions.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/bamboo_create_deb/bamboo_create_deb.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/bamboo_create_latest_deb/bamboo_create_latest_deb.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/bamboo_latest_tar_gz_url/bamboo_latest_tar_gz_url.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/bamboo_latest_version/bamboo_latest_version.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/bitbucket_create_deb/bitbucket_create_deb.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/bitbucket_create_latest_deb/bitbucket_create_deb.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/bitbucket_latest_tar_gz_url/bitbucket_latest_tar_gz_url.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/bitbucket_latest_version/bitbucket_latest_version.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/confluence_create_deb/confluence_create_deb.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/confluence_create_latest_deb/confluence_create_latest_deb.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/confluence_latest_tar_gz_url/confluence_latest_tar_gz_url.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/confluence_latest_version/confluence_latest_version.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/crowd_create_deb/crowd_create_deb.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/crowd_create_latest_deb/crowd_create_latest_deb.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/crowd_latest_tar_gz_url/crowd_latest_tar_gz_url.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/crowd_latest_version/crowd_latest_version.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/jira_core_create_deb/jira_core_create_deb.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/jira_core_create_latest_deb/jira_core_create_latest_deb.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/jira_core_latest_tar_gz_url/jira_core_latest_tar_gz_url.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/jira_core_latest_version/jira_core_latest_version.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/jira_servicedesk_create_deb/jira_servicedesk_create_deb.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/jira_servicedesk_create_latest_deb/jira_servicedesk_create_latest_deb.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/jira_servicedesk_latest_tar_gz_url/jira_servicedesk_latest_tar_gz_url.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/jira_servicedesk_latest_version/jira_servicedesk_latest_version.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/jira_software_create_deb/jira_software_create_deb.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/jira_software_create_latest_deb/jira_software_create_latest_deb.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/jira_software_latest_tar_gz_url/jira_software_latest_tar_gz_url.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/jira_software_latest_version/jira_software_latest_version.go
test:
	GO15VENDOREXPERIMENT=1 go test -cover `glide novendor`
vet:
	go tool vet .
	go tool vet --shadow .
lint:
	golint -min_confidence 1 ./...
errcheck:
	errcheck -ignore '(Close|Write)' ./...
check: lint vet errcheck
cov:
	mkdir -p target
	go test -coverprofile=target/coverage.out ./...
	go tool cover -func=target/coverage.out
	go tool cover -html=target/coverage.out
format:
	find . -name "*.go" -exec gofmt -w "{}" \;
	goimports -w=true .
prepare:
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/Masterminds/glide
	go get -u github.com/golang/lint/golint
	go get -u github.com/kisielk/errcheck
	glide install
update:
	glide up
clean:
	rm -rf vendor target
