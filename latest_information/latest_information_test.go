package latest_information

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsVersionInfo(t *testing.T) {
	b := New("", nil)
	var i *VersionInfo
	err := AssertThat(b, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseInfosSimple(t *testing.T) {
	infos, err := parseInfos([]byte(`downloads([{"version":"5.8.4","platform":"Unix"}])`))
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(len(infos), Is(1)); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(infos[0]["version"], Is("5.8.4")); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(infos[0]["platform"], Is("Unix")); err != nil {
		t.Fatal(err)
	}
}

func TestParseInfos(t *testing.T) {
	infos, err := parseInfos([]byte(`downloads([{"description":"5.9.2 - Linux Installer (64 bit)","zipUrl":"https://www.atlassian.com/software/confluence/downloads/binary/atlassian-confluence-5.9.2-x64.bin","tarUrl":null,"md5":"","size":"443.2 MB","released":"07-Dec-2015","type":"Binary","platform":"Unix","version":"5.9.2","releaseNotes":"http://confluence.atlassian.com/display/DOC/Confluence+5.9.2+Release+Notes","upgradeNotes":"http://confluence.atlassian.com/display/DOC/Confluence+5.9.2+Upgrade+Notes"},{"description":"5.9.2 - Linux Installer (32 bit)","zipUrl":"https://www.atlassian.com/software/confluence/downloads/binary/atlassian-confluence-5.9.2-x32.bin","tarUrl":null,"md5":"","size":"445.2 MB","released":"07-Dec-2015","type":"Binary","platform":"Unix","version":"5.9.2","releaseNotes":"http://confluence.atlassian.com/display/DOC/Confluence+5.9.2+Release+Notes","upgradeNotes":"http://confluence.atlassian.com/display/DOC/Confluence+5.9.2+Upgrade+Notes"},{"description":"5.9.2 - Windows Installer (64 bit)","zipUrl":"https://www.atlassian.com/software/confluence/downloads/binary/atlassian-confluence-5.9.2-x64.exe","tarUrl":null,"md5":"","size":"439.5 MB","released":"07-Dec-2015","type":"Binary","platform":"Windows","version":"5.9.2","releaseNotes":"http://confluence.atlassian.com/display/DOC/Confluence+5.9.2+Release+Notes","upgradeNotes":"http://confluence.atlassian.com/display/DOC/Confluence+5.9.2+Upgrade+Notes"},{"description":"5.9.2 - Windows Installer (32 bit)","zipUrl":"https://www.atlassian.com/software/confluence/downloads/binary/atlassian-confluence-5.9.2-x32.exe","tarUrl":null,"md5":"","size":"438.8 MB","released":"07-Dec-2015","type":"Binary","platform":"Windows","version":"5.9.2","releaseNotes":"http://confluence.atlassian.com/display/DOC/Confluence+5.9.2+Release+Notes","upgradeNotes":"http://confluence.atlassian.com/display/DOC/Confluence+5.9.2+Upgrade+Notes"},{"description":"5.9.2 - Standalone (TAR.GZ Archive)","zipUrl":"https://www.atlassian.com/software/confluence/downloads/binary/atlassian-confluence-5.9.2.tar.gz","tarUrl":null,"md5":"","size":"371.9 MB","released":"07-Dec-2015","type":"Binary","platform":"Mac OS X, Unix","version":"5.9.2","releaseNotes":"http://confluence.atlassian.com/display/DOC/Confluence+5.9.2+Release+Notes","upgradeNotes":"http://confluence.atlassian.com/display/DOC/Confluence+5.9.2+Upgrade+Notes"},{"description":"5.9.2 - Standalone (ZIP Archive)","zipUrl":"https://www.atlassian.com/software/confluence/downloads/binary/atlassian-confluence-5.9.2.zip","tarUrl":null,"md5":"","size":"373.7 MB","released":"07-Dec-2015","type":"Binary","platform":"Windows","version":"5.9.2","releaseNotes":"http://confluence.atlassian.com/display/DOC/Confluence+5.9.2+Release+Notes","upgradeNotes":"http://confluence.atlassian.com/display/DOC/Confluence+5.9.2+Upgrade+Notes"}])`))
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(len(infos), Is(6)); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(infos[0]["version"], Is("5.9.2")); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(infos[0]["platform"], Is("Unix")); err != nil {
		t.Fatal(err)
	}
}

func TestParseInfosZipUrl(t *testing.T) {
	infos, err := parseInfos([]byte(`downloads([{"description":"2.12.0 - Standalone (ZIP Archive)","zipUrl":"https://www.atlassian.com/software/crowd/downloads/binary/atlassian-crowd-2.12.0.zip","tarUrl":null,"md5":null,"size":"172.4 MB","released":"26-Apr-2017","type":"Binary","platform":"Windows","version":"2.12.0","releaseNotes":"http://confluence.atlassian.com/display/CROWD/Crowd+2.12+Release+Notes","upgradeNotes":"http://confluence.atlassian.com/display/CROWD/Crowd+2.12+Upgrade+Notes"},{"description":"2.12.0 - Standalone (TAR.GZ Archive)","zipUrl":"https://www.atlassian.com/software/crowd/downloads/binary/atlassian-crowd-2.12.0.tar.gz","tarUrl":null,"md5":null,"size":"171.2 MB","released":"26-Apr-2017","type":"Binary","platform":"Unix, Mac OS X","version":"2.12.0","releaseNotes":"http://confluence.atlassian.com/display/CROWD/Crowd+2.12+Release+Notes","upgradeNotes":"http://confluence.atlassian.com/display/CROWD/Crowd+2.12+Upgrade+Notes"},{"description":"2.12.0 - CrowdID Server - WAR","zipUrl":"https://www.atlassian.com/software/crowd/downloads/binary/atlassian-crowd-openid-2.12.0-war.zip","tarUrl":null,"md5":null,"size":"40.2 MB","released":"26-Apr-2017","type":"Binary","platform":"Windows, Unix, Mac OS X","version":"2.12.0","releaseNotes":"http://confluence.atlassian.com/display/CROWD/Crowd+2.12+Release+Notes","upgradeNotes":"http://confluence.atlassian.com/display/CROWD/Crowd+2.12+Upgrade+Notes"},{"description":"2.12.0 - Crowd Server - WAR","zipUrl":"https://www.atlassian.com/software/crowd/downloads/binary/atlassian-crowd-2.12.0-war.zip","tarUrl":null,"md5":null,"size":"81.9 MB","released":"26-Apr-2017","type":"Binary","platform":"Windows, Unix, Mac OS X","version":"2.12.0","releaseNotes":"http://confluence.atlassian.com/display/CROWD/Crowd+2.12+Release+Notes","upgradeNotes":"http://confluence.atlassian.com/display/CROWD/Crowd+2.12+Upgrade+Notes"}])`))
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(len(infos), Is(4)); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(infos[0]["zipUrl"], Is("https://www.atlassian.com/software/crowd/downloads/binary/atlassian-crowd-2.12.0.zip")); err != nil {
		t.Fatal(err)
	}
}
