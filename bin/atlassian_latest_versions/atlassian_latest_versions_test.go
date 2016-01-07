package main

import (
	"testing"

	"fmt"

	. "github.com/bborbe/assert"
)

func TestDoFail(t *testing.T) {
	var err error
	list := doMap(map[string]LatestVersion{"Test": func() (string, error) {
		return "", fmt.Errorf("foo")
	}})
	if err = AssertThat(len(list), Is(1)); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(list[0], Is("Test: failed")); err != nil {
		t.Fatal(err)
	}
}

func TestDoSuccess(t *testing.T) {
	var err error
	list := doMap(map[string]LatestVersion{"Test": func() (string, error) {
		return "1.2.3", nil
	}})
	if err = AssertThat(len(list), Is(1)); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(list[0], Is("Test: 1.2.3")); err != nil {
		t.Fatal(err)
	}
}
