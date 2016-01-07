package main

import (
	"testing"

	"fmt"

	. "github.com/bborbe/assert"
	io_mock "github.com/bborbe/io/mock"
)

func TestDoFail(t *testing.T) {
	var err error
	writer := io_mock.NewWriter()
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
	writer := io_mock.NewWriter()
	err = do(writer, func() (string, error) {
		return "1.2.3", nil
	})
	err = AssertThat(err, NilValue())
	if err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(string(writer.Content()), Is("1.2.3\n")); err != nil {
		t.Fatal(err)
	}
}
