package main

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestDo(t *testing.T) {
	var err error
	err = do(nil, nil, "", "", "", "")
	err = AssertThat(err, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
}
