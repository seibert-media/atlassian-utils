package main

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestDo(t *testing.T) {
	var err error
	err = do(nil, nil, "", "", "", "")
	if err = AssertThat(err, NotNilValue()); err != nil {
		t.Fatal(err)
	}
}

func TestExtractAtlassianVersion(t *testing.T) {
	if err := AssertThat(extractAtlassianVersion("1.2.3"), Is("1.2.3")); err != nil {
		t.Fatal(err)
	}
}

func TestExtractAtlassianVersionWithTag(t *testing.T) {
	if err := AssertThat(extractAtlassianVersion("1.2.3-1.0.0"), Is("1.2.3")); err != nil {
		t.Fatal(err)
	}
}
