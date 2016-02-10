package main

import (
	"testing"

	"fmt"

	"bytes"

	. "github.com/bborbe/assert"
)

func TestDoFail(t *testing.T) {
	var err error
	writer := bytes.NewBufferString("")
	err = do(writer, func() (string, error) {
		return "", fmt.Errorf("fail")
	})
	err = AssertThat(err, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
}

func TestDoSuccess(t *testing.T) {
	var err error
	writer := bytes.NewBufferString("")
	err = do(writer, func() (string, error) {
		return "1.2.3", nil
	})
	err = AssertThat(err, NilValue())
	if err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(writer.String(), Is("1.2.3\n")); err != nil {
		t.Fatal(err)
	}
}
