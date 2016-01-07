package main

import (
	"testing"

	"fmt"

	. "github.com/bborbe/assert"
)

func TestDoFail(t *testing.T) {
	_, err := doMap(map[string]LatestVersion{"Test": func() (string, error) {
		return "", fmt.Errorf("fail")
	}})
	err = AssertThat(err, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
}

func TestDoSuccess(t *testing.T) {
	_, err := doMap(map[string]LatestVersion{"Test": func() (string, error) {
		return "1.2.3", nil
	}})
	err = AssertThat(err, NilValue())
	if err != nil {
		t.Fatal(err)
	}
}
