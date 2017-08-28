package main

import (
	"testing"

	"fmt"

	. "github.com/bborbe/assert"
	debian_config "github.com/bborbe/debian_utils/config"
)

func TestDo(t *testing.T) {
	var err error
	err = do(func(config *debian_config.Config, sourceDir string, targetDir string) error {
		return fmt.Errorf("foo")
	}, nil, "", func() (string, error) {
		return "1.2.3", nil
	}, "")
	err = AssertThat(err, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
}
